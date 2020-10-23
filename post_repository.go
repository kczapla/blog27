package main

import (
    "gorm.io/gorm"
)

type Repository interface {
    Get(id string) (Post, error)
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
    return repository{db}
}

func (r repository) Get(id string) (Post, error) {
    var post Post
    result := r.db.First(&post, id)
    return post, result.Error
}
