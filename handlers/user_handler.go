package handlers

import (
	//"errors"
	"github.com/Vilyan01/messaging-api/data"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// find the user
	user, err := data.FindUserById(1)
	// check for error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// render the user in JSON form
	user_json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(user_json)
}
