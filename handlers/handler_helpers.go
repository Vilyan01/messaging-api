package handlers

import (
	"github.com/go-martini/martini"
	"strconv"
)

func IdFromParams(p martini.Params) (int, error) {
	return strconv.Atoi(p["id"])
}
