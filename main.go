package main

import (
	"github.com/Vilyan01/messaging-api/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// create new router
	r := mux.NewRouter()

	// define routes
	r.HandleFunc("/users/{id}", handlers.GetUser)

	// start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
