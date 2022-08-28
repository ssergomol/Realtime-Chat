package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password []byte    `json:"-"`
	Messages []Message `json:"messages"`
}

type Message struct {
	gorm.Model
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"userid"`
}
