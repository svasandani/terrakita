package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/svasandani/terrakita/internal/db"
  "github.com/svasandani/terrakita/internal/api"
)

type EVDB struct {
	Username string
	Password string
	Host string
	Port string
	Database string
}

type EVS struct {
	Host string
	Port string
}

type EnvVars struct {
	DatabaseConnection EVDB
	Server EVS
}

func main() {
	ev := parseEnvVariables()

	evdb := ev.DatabaseConnection
	err := db.ConnectToDatabase(db.DatabaseConnection{
		Username: evdb.Username,
		Password: evdb.Password,
		Host: evdb.Host,
		Port: evdb.Port,
		Database: evdb.Database,
	})
	if err != nil {
		log.Fatalf("Couldn't connect to database! %+v", err)
	}
	
	fmt.Println("terrakita")
  s := createBackendServer(ev)
	log.Fatal(s.ListenAndServe())
}

func parseEnvVariables() EnvVars {
	dbu := flag.String("databaseUsername", "terraling", "Username for MySQL")
	dbp := flag.String("databasePassword", "terraling", "Password for MySQL")
	dbh := flag.String("databaseHost", "localhost", "Host for MySQL")
	dbpt := flag.String("databasePort", "3306", "Port for MySQL")
	dbd := flag.String("databaseDatabase", "terraling", "Database for MySQL")
	
	sh := flag.String("serverHost", "localhost", "Host to launch the server on")
	sp := flag.String("serverPort", "7107", "Port to launch the server on")

	return EnvVars{
		DatabaseConnection: EVDB{
			Username: *dbu,
			Password: *dbp,
			Host: *dbh,
			Port: *dbpt,
			Database: *dbd,
		},
		Server: EVS{
			Host: *sh,
			Port: *sp,
		},
	}
}

func createBackendServer(ev EnvVars) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/search/filter", api.PostSearchFilterHandler)

	s := ev.Server
	addr := fmt.Sprintf("%v:%v", s.Host, s.Port)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &server
}
