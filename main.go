package main

import (
	"github.com/Vilyan01/messaging-api/handlers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Test struct {
	Name string `json:"name"`
}

func main() {
	// Use a classic martini server to host the API
	m := martini.Classic()
	// Use the renderer to easily render JSON/HTML
	m.Use(render.Renderer())
	// Test function for now...
	m.Get("/", RenderTest)
	m.Get("/users/:id", handlers.GetUser)
	// Start the API
	m.Run()
}

func RenderTest(r render.Render) {
	n := Test{Name: "Heya"}
	r.JSON(200, n)
}
