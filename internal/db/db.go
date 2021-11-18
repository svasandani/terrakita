package db

import (
	// "errors"
	"fmt"
	"log"
	"strconv"

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

	sel, err := db.Prepare("select id, name from lings where id = ?")
	defer sel.Close()
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterResponse{}, err
	}

	var l Ling

	err = sel.QueryRow(971).Scan(&l.Id, &l.Name)
	if err != nil {
		log.Print("Error executing database request!")
		return FilterResponse{}, err
	}

	lings := make([]FilterResponseLing, 1)
	pv := make([]PropertyValuePair, 1)

	pv[0] = PropertyValuePair{
		Property: strconv.Itoa(l.Id),
		Value:    l.Name,
	}

	lings[0] = FilterResponseLing{
		PropertyValuePairs: pv,
	}

	return FilterResponse{
		Lings: lings,
	}, nil
}
