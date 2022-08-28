package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password []byte    `json:"-"`
	Messages []Message `json:"messages"`
}

type Message struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"userid"`
}
