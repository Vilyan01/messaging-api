package handlers

import (
	"errors"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
