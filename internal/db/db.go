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

	hasLing := len(frq.Lings) != 0
	hasLingProperties := len(frq.LingProperties) != 0
	hasLinglet := len(frq.Linglets) != 0
	// hasLingletProperties := len(frq.LingletProperties) != 0

	if hasLing && !hasLingProperties && !hasLinglet {
		benchmark.Start("Filter lings only")
		fr, err := filterLings(frq)
		benchmark.Stop("Filter lings only")
		if err != nil {
			log.Print("Error filtering lings!")
			return FilterResponse{}, err
		}

		return fr, nil
	}

	return FilterResponse{}, nil
}

func filterLings(frq FilterRequest) (FilterResponse, error) {
	lings := make([]FilterResponseLing, len(frq.Lings))

	// pass group then lings into query args
	qargs := make([]interface{}, len(frq.Lings) + 1)
	qargs[0] = frq.Group
	for i, id := range frq.Lings {
		qargs[i + 1] = id
	}

	stmt := `SELECT id, name FROM lings WHERE group_id = ? AND id IN (?` + strings.Repeat(",?", len(qargs) - 2) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l FilterResponseLing

		err = ls.Scan(&l.Id, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterResponse{}, err
		}

		// Select properties
		rows, err := db.Query("SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ?", frq.Group, l.Id)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterResponse{}, err
		}
		defer rows.Close()

		var pvs []PropertyValuePair

		for rows.Next() {
			var pv PropertyValuePair

			err = rows.Scan(&pv.Id, &pv.Property, &pv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterResponse{}, err
			}

			pvs = append(pvs, pv)
		}

		l.PropertyValuePairs = pvs
		lings[i] = l
		i++
	}

	return FilterResponse{
		Lings: lings,
	}, nil
}
