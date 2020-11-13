package main

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string  `gorm:"unique"`
	Posts []*Post `gorm:"many2many:post_tags;"`
}

type Tags []Tag
