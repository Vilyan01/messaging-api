package data

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	ConversationID int    `json:"conversation_id" gorm:"index"`
	SenderID       int    `json:"sender_id"`
	Body           string `json:"body" gorm:type:varchar(255)"`
}
