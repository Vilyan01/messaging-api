package main

import (
	"github.com/Vilyan01/messaging-api/data"
)

func main() {
	usr := data.NewUser("hello@go.com", "Shankmeister")
	usr.Save()
	usr.Type()
}
