package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Set up routes
	initRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
