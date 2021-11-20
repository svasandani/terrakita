package db

import (
	"fmt"
	"log"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/svasandani/terrakita/internal/benchmark"
)

/******** Generic Database Functions ********/

var db *sql.DB

// ConnectToDatabase - Connects to the database.
func ConnectToDatabase(dbc DatabaseConnection) error {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Database)

	dbl, err := sql.Open("mysql", conn)
	if err != nil {
		log.Print("Error creating database connection!")
		return err
	}

	benchmark.StartTare()
	err = dbl.Ping()
	benchmark.StopTare()
	if err != nil {
		log.Print("Error connecting to database!")
		return err
	}

	db = dbl

	return nil
}

// Filter - Retrieves all ling and linglet property-value pairs that fit the criteria.
func Filter(frq FilterRequest) (FilterResponse, error) {
	err := validateFilterRequest(frq)
	if err != nil {
		log.Print("Error in filter request!")
		return FilterResponse{}, err
	}

	hasLings := len(frq.Lings) != 0
	hasLingProperties := len(frq.LingProperties) != 0
	hasLinglets := len(frq.Linglets) != 0
	hasLingletProperties := len(frq.LingletProperties) != 0

	if hasLings {
		benchmark.Start("Filter lings only")
		fr, err := filterLings(frq)
		benchmark.Stop("Filter lings only")
		if err != nil {
			log.Print("Error filtering lings!")
			return FilterResponse{}, err
		}

		return fr, nil
	} else if hasLinglets {
		benchmark.Start("Filter linglets only")
		fr, err := filterLinglets(frq)
		benchmark.Stop("Filter linglets only")
		if err != nil {
			log.Print("Error filtering linglets!")
			return FilterResponse{}, err
		}

		return fr, nil
	} else if hasLingProperties {

	} else if hasLingletProperties {

	}

	return FilterResponse{}, nil
}

func FilterLings(flr FilterLingsRequest) (FilterLingsResponse, error) {
	lings := make([]FilterLingsResponseLing, len(flr.Lings))

	// pass group then lings into query args
	qargs := make([]interface{}, len(flr.Lings)+1)
	qargs[0] = flr.Group
	for i, id := range flr.Lings {
		qargs[i+1] = id
	}

	// SELECT lings
	stmt := `SELECT id, name FROM lings WHERE group_id = ? AND depth = 0 AND id IN (?` + strings.Repeat(",?", len(flr.Lings)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l FilterLingsResponseLing

		err = ls.Scan(&l.Id, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingsResponse{}, err
		}

		qargs = make([]interface{}, len(flr.LingProperties)+2)
		qargs[0] = flr.Group
		qargs[1] = l.Id
		for i, id := range flr.LingProperties {
			qargs[i+2] = id
		}

		// if we need all properties, perform an extra query in case of early return
		if len(flr.LingProperties) > 0 && flr.LingPropertiesInclusive {
			stmt = `SELECT COUNT(properties.id) FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ? AND lings_properties.property_id IN (?` + strings.Repeat(",?", len(flr.LingProperties)-1) + `)`
			c := db.QueryRow(stmt, qargs...)

			var count int
			c.Scan(&count)

			if count != len(flr.LingProperties) {
				continue
			}
		}

		// build statement dynamically
		stmt = "SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ?"
		if len(flr.LingProperties) != 0 {
			stmt += ` AND lings_properties.property_id IN (?` + strings.Repeat(",?", len(flr.LingProperties)-1) + `)`
		}

		// SELECT properties
		ps, err := db.Query(stmt, qargs...)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterLingsResponse{}, err
		}
		defer ps.Close()

		pvs := make([]NameValuePair, 0)

		for ps.Next() {
			var nv NameValuePair

			err = ps.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterLingsResponse{}, err
			}

			pvs = append(pvs, nv)
		}

		l.PropertyValuePairs = pvs
		lings[i] = l
		i++
	}

	return FilterLingsResponse{
		Type:  "filter",
		On:    "lings",
		Lings: lings[:i],
	}, nil
}

func FilterLinglets(fllr FilterLingletsRequest) (FilterLingletsResponse, error) {
	lmap := make(map[Ling][]FilterLingletsResponseLinglet)

	// pass group then lings into query args
	qargs := make([]interface{}, len(fllr.Linglets)+1)
	qargs[0] = fllr.Group
	for i, id := range fllr.Linglets {
		qargs[i+1] = id
	}

	stmt := `SELECT ling.id, ling.name, linglet.id, linglet.name FROM lings AS ling INNER JOIN lings AS linglet ON ling.id = linglet.parent_id WHERE linglet.group_id = ? AND linglet.depth = 1 AND linglet.id IN (?` + strings.Repeat(",?", len(fllr.Linglets)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingletsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l Ling
		var ll FilterLingletsResponseLinglet

		err = ls.Scan(&l.Id, &l.Name, &ll.Id, &ll.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingletsResponse{}, err
		}

		qargs = make([]interface{}, len(fllr.LingletProperties)+2)
		qargs[0] = fllr.Group
		qargs[1] = ll.Id
		for i, id := range fllr.LingletProperties {
			qargs[i+2] = id
		}

		// if we need all properties, perform an extra query for early return
		if len(fllr.LingletProperties) > 0 && fllr.LingletPropertiesInclusive {
			stmt = `SELECT COUNT(properties.id) FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ? AND lings_properties.property_id IN (?` + strings.Repeat(",?", len(fllr.LingletProperties)-1) + `)`
			c := db.QueryRow(stmt, qargs...)

			var count int
			c.Scan(&count)

			if count != len(fllr.LingletProperties) {
				continue
			}
		}

		// build statement dynamically
		stmt = "SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ?"
		if len(fllr.LingletProperties) != 0 {
			stmt += ` AND lings_properties.property_id IN (?` + strings.Repeat(",?", len(fllr.LingletProperties)-1) + `)`
		}

		// Select properties
		ps, err := db.Query(stmt, qargs...)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterLingletsResponse{}, err
		}
		defer ps.Close()

		pvs := make([]NameValuePair, 0)

		for ps.Next() {
			var nv NameValuePair

			err = ps.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterLingletsResponse{}, err
			}

			pvs = append(pvs, nv)
		}

		ll.PropertyValuePairs = pvs

		if a, ok := lmap[l]; ok {
			lmap[l] = append(a, ll)
		} else {
			lmap[l] = []FilterLingletsResponseLinglet{ll}
		}

		i++
	}

	lings := make([]FilterLingletsResponseLing, 0)

	for k, e := range lmap {
		lings = append(lings, FilterLingletsResponseLing{
			Id:       k.Id,
			Name:     k.Name,
			Linglets: e,
		})
	}

	return FilterLingletsResponse{
		Type:  "filter",
		On:    "linglets",
		Lings: lings,
	}, nil
}
