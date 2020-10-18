package main

import (
    "gorm.io/gorm"
)


type UserPost struct {
    UserId uint `json:"UserId"`
    PostId uint `json:"PostId"`
}
