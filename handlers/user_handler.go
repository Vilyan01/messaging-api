package handlers

import (
	"errors"
	"github.com/Vilyan01/messaging-api/data"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetUser(r render.Render, p martini.Params) {
	id, err := IdFromParams(p)
	if err != nil {
		r.JSON(500, errors.New("Error pulling ID from params"))
		return
	}
	user, err := data.FindUserByID(id)
	if err != nil {
		r.JSON(404, NotFoundError("user"))
		return
	}
	r.JSON(200, user)
}
