package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"password"`
	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}

func (u *User) SaveUser() (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err

	}
	return u, nil
}

// func (u *User) BeforeSave() error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)

// 	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

// 	return nil
// }
