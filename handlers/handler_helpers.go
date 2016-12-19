package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func IdFromRequest(r *http.Request) (int, error) {
	v := mux.Vars(r)
	return strconv.Atoi(v["id"])
}
