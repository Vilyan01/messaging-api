package data

import (
	"github.com/jinzhu/gorm"
)

/*
  Struct to hold user information and relay it to the database.
*/
type User struct {
	gorm.Model
	Username       string         `json:"username" gorm:"type:varchar(25);unique_index" binding:"required"`
	Email          string         `json:"email" gorm:"type:varchar(100);unique_index" binding:"required"`
	Conversations  []Conversation `json:"conversations" gorm:"many2many:user_conversations;"`
	HashedPassword string         `json:"-"`
	Session        Session        `json:"-"`
	SessionID      int            `json:"-"`
}

func FindUserByID(id int) (User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	return user, err
}

func (u User) SaveUser() error {
	return db.DB.Create(&u).Error
}
