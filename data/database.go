package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dbName string) (Database, error) {
	// variables to hold data
	var (
		err      error
		db       *gorm.DB
		database Database
	)

	// get password for database from environment variable
	pass := os.Getenv("MESSAGING_PASS")
	user := os.Getenv("MESSAGING_USER")
	// create settings string for database connection
	settings := fmt.Sprintf("host=localhost user=%v dbname=%v sslmode=disable password=%v", user, dbName, pass)
	// Open the database connection
	db, err = gorm.Open("postgres", settings)
	// Check for errors opening database
	if err != nil {
		return database, err
	}
	// create the new database object to hold the connection
	database = Database{DB: db}
	return database, err
}
