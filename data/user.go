package data

import (
	"github.com/jinzhu/gorm"
)

/*
  Struct to hold user information and relay it to the database.
*/
type User struct {
	gorm.Model
	Username       string         `json:"username" gorm:"type:varchar(25);unique_index"`
	Email          string         `json:"email" gorm:"type:varchar(100);unique_index"`
	Conversations  []Conversation `json:"conversations" gorm:"many2many:user_conversations;"`
	HashedPassword string         `json:"-"`
	Session        Session        `json:"-"`
	SessionID      int            `json:"-"`
}

func FindUserByID(id int) (User, error) {
	var user User
	err := db.DB.First(&user, 1).Error
	return user, err
}
