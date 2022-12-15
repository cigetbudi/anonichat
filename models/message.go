package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text     string `json:"text"`
	Like     int64  `json:"like"`
	Location string `json:"location"`
	UserID   int64  `json:"user_id"`
	User     User
}

func GetAllMessages() (*[]Message, error) {
	var mes []Message
	err := DB.Debug().Model(&Message{}).Limit(50).Find(&mes).Error
	if err != nil {
		return &[]Message{}, err
	}
	return &mes, nil
}

func GetMessagesByUserID(id uint) (*[]Message, error) {
	var mes []Message
	err := DB.Where("user_id = ?", id).Find(&mes).Error
	if err != nil {
		return &[]Message{}, err
	}
	return &mes, nil
}

func CreateMessage(m *Message) error {

	var err error
	err = DB.Debug().Model(Message{}).Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteMessage(id, user_id uint) error {
	m := Message{}
	var err error
	err = DB.Where("id = ? AND user_id = ?", id, user_id).Find(&m).Error
	if err != nil {
		return err
	}
	fmt.Println(m)
	err = DB.Where("id = ? AND user_id = ?", id, user_id).Delete(&Message{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddLike(id uint) error {
	err := DB.Exec(`update messages set "like" = "like" + 1 where id = ? `, id).Error
	if err != nil {
		return err
	}
	return nil
}

func UnLike(id uint) error {
	err := DB.Exec(`update messages set "like" = "like" - 1 where id = ? `, id).Error
	if err != nil {
		return err
	}
	return nil
}
