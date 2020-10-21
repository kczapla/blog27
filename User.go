package main

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name" gorm:"unique"`
    Email string `json:"email" gorm:"unique"`
    Posts []Post `json:"posts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

