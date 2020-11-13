package main

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	UserID   uint     `json:"userId"`
	Comments Comments `json:"posts" gorm:"constraint:OnDelete:SET NULL;"`
	Tags     []*Tag   `json:"tags" gorm:"many2many:post_tags"`
}

type Posts []Post
