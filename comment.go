package main

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
	Content string `json:"content"`
}

type Comments []Comment
