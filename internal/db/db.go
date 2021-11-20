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

func FilterLings(flr FilterLingsRequest) (FilterLingsResponse, error) {
	err := validateFilterLingsRequest(flr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingsResponse{}, err
	}

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

func FilterLingProperties(flpr FilterLingPropertiesRequest) (FilterLingPropertiesResponse, error) {
	err := validateFilterLingPropertiesRequest(flpr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingPropertiesResponse{}, err
	}

	properties := make([]FilterLingPropertiesResponseProperty, len(flpr.LingProperties))

	// pass group then ling properties into query args
	qargs := make([]interface{}, len(flpr.LingProperties)+1)
	qargs[0] = flpr.Group
	for i, id := range flpr.LingProperties {
		qargs[i+1] = id
	}

	// SELECT properties
	stmt := `SELECT id, name FROM properties WHERE group_id = ? AND depth = 0 AND id IN (?` + strings.Repeat(",?", len(flpr.LingProperties)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingPropertiesResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var p FilterLingPropertiesResponseProperty

		err = ls.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingPropertiesResponse{}, err
		}

		qargs = make([]interface{}, len(flpr.Lings)+2)
		qargs[0] = flpr.Group
		qargs[1] = p.Id
		for i, id := range flpr.Lings {
			qargs[i+2] = id
		}

		// if we need all lings, perform an extra query in case of early return
		if len(flpr.Lings) > 0 && flpr.LingsInclusive {
			stmt = `SELECT COUNT(lings.id) FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings_properties.group_id = ? AND ling_id = ? AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(flpr.Lings)-1) + `)`
			c := db.QueryRow(stmt, qargs...)

			var count int
			c.Scan(&count)

			if count != len(flpr.Lings) {
				continue
			}
		}

		// build statement dynamically
		stmt = "SELECT lings.id, lings.name, lings_properties.value FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings_properties.group_id = ? AND property_id = ?"
		if len(flpr.Lings) != 0 {
			stmt += ` AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(flpr.Lings)-1) + `)`
		}

		// SELECT lings
		ls, err := db.Query(stmt, qargs...)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterLingPropertiesResponse{}, err
		}
		defer ls.Close()

		lvs := make([]NameValuePair, 0)

		for ls.Next() {
			var nv NameValuePair

			err = ls.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterLingPropertiesResponse{}, err
			}

			lvs = append(lvs, nv)
		}

		p.LingValuePairs = lvs
		properties[i] = p
		i++
	}

	return FilterLingPropertiesResponse{
		Type:       "filter",
		On:         "ling_properties",
		Properties: properties[:i],
	}, nil
}

func FilterLinglets(fllr FilterLingletsRequest) (FilterLingletsResponse, error) {
	err := validateFilterLingletsRequest(fllr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingletsResponse{}, err
	}

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
