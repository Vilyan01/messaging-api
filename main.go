package main

import (
	"github.com/Vilyan01/messaging-api/data"
)

func main() {
	usr := data.NewClient("hello@go.com", "rherherhaerterter")
	usr.Save()
}
