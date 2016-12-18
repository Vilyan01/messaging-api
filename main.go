package main

import (
	"github.com/Vilyan01/messaging-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// create new router
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handlers.GetUser)
}
