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

var (
	// The data models that are used throughout this app will use this to
	// communicate.
	db Database
)

func init() {
	var err error
	// Create the database for the subsequent models to use
	db, err = NewDatabase("messaging_test")
	if err != nil {
		initErr := fmt.Sprintf("Unable to initialize database connection: %v", err)
		fmt.Println(initErr)
		os.Exit(1)
	}
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
