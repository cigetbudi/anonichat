package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Text     string `json:"text"`
	Like     int64  `json:"like"`
	Location string `json:"location"`
	UserID   int64  `json:"user_id"`
	User     User
}
