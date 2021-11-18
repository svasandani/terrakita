package main

import (
  "fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("terrakita")
  s := createBackendServer()
	log.Fatal(s.ListenAndServe())
}


func createBackendServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
    fmt.Println(path)
		w.Write([]byte(fmt.Sprintf("hello")))
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	return &server
}
