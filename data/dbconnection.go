package data

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

const DBNAME = "messaging"

/*
	Struct to hold the database connection.
*/
type DBConnection struct {
	Session *r.Session
}

/*
	Creates a new DB connection session
*/
func NewDBConnection() (DBConnection, error) {
	var err error
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		fmt.Println(err)
		return DBConnection{}, err
	}
	return DBConnection{Session: session}, err
}

/*
	Insert an object into a table
*/

func (d DBConnection) InsertObject(m Model) error {
	res, err := r.DB(DBNAME).Table(m.TableName).Insert(m).RunWrite(d.Session)
	fmt.Println(res)
	return err
}
