package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"user_name"`
	Password string `json:"password" gorm:"password"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"email"`
}
