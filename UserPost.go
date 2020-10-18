package main

import (
    "gorm.io/gorm"
)


type UserPost struct {
    gorm.Model
    UserId uint `json:"UserId"`
    PostId uint `json:"PostId"`
}
