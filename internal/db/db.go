package db

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/******** Generic Database Functions ********/

var db *sql.DB

// ConnectToDatabase - Connects to the database.
func ConnectToDatabase(dbc DatabaseConnection) *sql.DB {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Database)

	dbl, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Errorf("Error opening connection to database: %+v", err)
	}

	err = dbl.Ping()
	if err != nil {
		fmt.Errorf("Error establishing connection to database: %+v", err)
	}

	db = dbl

	return dbl
}

// Filter - Retrieves all ling and linglet property-value pairs that fit the criteria.
func Filter(frq FilterRequest) (FilterResponse, error) {
  // TODO
}
