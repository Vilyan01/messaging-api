package data

import (
	"github.com/jinzhu/gorm"
)

type Conversation struct {
	gorm.Model
	Title    string    `json:"title" gorm:"type:varchar(100);unique_index"`
	Users    []User    `json:"participants" gorm:"many2many:user_conversations;"`
	Messages []Message `json:"messages"`
}
