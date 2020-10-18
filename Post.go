package main

import (
    "gorm.io/gorm"
)


type Post struct {
    gorm.Model
    Title string `json:"title"`
    Content string `json:"content"`
    UserID uint `json:"userId"`
    User User //`json:"user"`
}

type Posts []Post
