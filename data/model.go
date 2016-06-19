package data

import (
	"log"
)

/*
	All models will inherit from this struct and it's methods for database
	entry and retrieval
*/
type Model struct {
	TableName string `json:"-",gorethink:"-"` // The name of the table that each model will be saved in.
}

/*
	Saves a model to the database.  Since there is currently no database, it will
	just print the TableName for now.
*/
func (m Model) Save() {
	conn, err := NewDBConnection()
	if err != nil {
		log.Fatalln(err.Error())
	}
	conn.InsertObject(m)
}
