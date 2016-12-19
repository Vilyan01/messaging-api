package handlers

import (
	//"errors"
	"encoding/json"
	"github.com/Vilyan01/messaging-api/data"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// find the user
	w.Header().Set("Content-Type", "application/json")
	id, _ := IdFromRequest(r)
	user, err := data.FindUserByID(id)
	// check for error
	if err != nil {
		ej, _ := json.Marshal(NewHandlerError(err))
		w.WriteHeader(http.StatusNotFound)
		w.Write(ej)
		return
	}
	// render the user in JSON form
	user_json, err := json.Marshal(user)
	if err != nil {
		ej, _ := json.Marshal(NewHandlerError(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ej)
		return
	}
	w.Write(user_json)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get the request bodya.
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var u data.User
	err := decoder.Decode(&u)
	if err != nil {
		ej, _ := json.Marshal(NewHandlerError(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ej)
		return
	}
	err = u.SaveUser()
	if err != nil {
		ej, _ := json.Marshal(NewHandlerError(err))
		w.WriteHeader(http.StatusConflict)
		w.Write(ej)
		return
	}
	user_json, err := json.Marshal(u)
	w.Write(user_json)
}
