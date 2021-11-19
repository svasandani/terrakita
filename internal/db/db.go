package db

import (
	// "errors"
	"fmt"
	"log"

	// "regexp"
	// "time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

	err = dbl.Ping()
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
		fr, err := filterLings(frq)
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

	for i, lid := range frq.Lings {
		// Select lings
		sel, err := db.Prepare("SELECT id, name FROM lings WHERE id = ?")
		defer sel.Close()
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterResponse{}, err
		}

		var l Ling

		err = sel.QueryRow(lid).Scan(&l.Id, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterResponse{}, err
		}

		// Select properties
		rows, err := db.Query("SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE ling_id = ? and lings_properties.group_id = ?", lid, frq.Group)
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

		lings[i] = FilterResponseLing{
			Id:                 l.Id,
			Name:               l.Name,
			PropertyValuePairs: pvs,
		}
	}

	return FilterResponse{
		Lings: lings,
	}, nil
}
