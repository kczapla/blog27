package main

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string `json:"name"`
    Email string `json:"email"`
    Posts []Post `json:"posts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

