package data

import (
	"github.com/jinzhu/gorm"
)

/*
  Struct to hold user information and relay it to the database.
*/
type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"type:varchar(25);unique_index"`
	Email          string `json:"email" gorm:"type:varchar(100);unique_index"`
	HashedPassword string `json:"-"`
}

func init() {
	// When the application starts, we need to automigrate.
	db.DB.AutoMigrate(&User{})
}

func FindUserByID(id int) (User, error) {
	var user User
	err := db.DB.First(&user, 1).Error
	return user, err
}
