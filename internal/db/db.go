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
    log.Print("Error creating database connection!");
		return err
	}

	err = dbl.Ping()
	if err != nil {
    log.Print("Error connecting to database!");
		return err
	}

	db = dbl

  return nil
}

// Filter - Retrieves all ling and linglet property-value pairs that fit the criteria.
func Filter(frq FilterRequest) (FilterResponse, error) {
  // TODO

  return FilterResponse{}, nil
}
