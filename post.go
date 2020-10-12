package main

import (
    "gorm.io/gorm"
)


type Post struct {
    gorm.Model
    Title string `json:"Title"`
    Content string `json:"Content"`
}
