package main

import (
    "gorm.io/gorm"
)


type User struct {
    gorm.Model
    Nickname string `json:"Nickname"`
    Email string `json:"Email"`
}

type UserPost struct {
    gorm.Model
    UserId uint `json:"UserId"`
    PostId uint `json:"PostId"`
}
