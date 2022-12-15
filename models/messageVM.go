package models

type MessageVM struct {
	Text     string `json:"text"`
	Like     int64  `json:"like"`
	Location string `json:"location"`
	UserID   int64  `json:"user_id"`
}
