package main

import (
    "gorm.io/gorm"
)


type Post struct {
    gorm.Model
    Title string `json:"title"`
    Content string `json:"content"`
    UserID uint `json:"userId"`
}

type Posts []Post
