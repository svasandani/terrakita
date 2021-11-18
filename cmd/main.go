package main

import (
  "fmt"
	"log"
	"net/http"

  "github.com/svasandani/terrakita/internal/api"
)

func main() {
	fmt.Println("terrakita")
  s := createBackendServer()
	log.Fatal(s.ListenAndServe())
}


func createBackendServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/search/filter", api.PostSearchFilterHandler)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	return &server
}
