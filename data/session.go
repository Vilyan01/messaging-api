package data

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID     int       `json:"user_id" gorm:"index"`
	Expires    time.Time `json:"expires"`
	SessionKey string    `json:"key" gorm:"type:varchar(255)"`
}
