package main

import (
	"gorm.io/gorm"
)

type Repository interface {
	Get(id string) (Post, error)
	Create(post Post) error
	Delete(id string) error
	Update(post Post) error
	Query(post Post) (Posts, error)
	QueryTags(post Post) (Tags, error)
	AddTag(post Post, tag Tag) error
	RemoveTag(post Post, tag Tag) error
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

func (r repository) Create(post Post) error {
	return r.db.Create(&post).Error
}

func (r repository) Delete(id string) error {
	return r.db.Delete(&Post{}, id).Error
}

func (r repository) Update(post Post) error {
	return r.db.Save(&post).Error
}

func (r repository) Query(post Post) (Posts, error) {
	var posts Posts
	result := r.db.Where(&post).Find(&posts)
	return posts, result.Error
}

func (r repository) QueryTags(post Post) (Tags, error) {
	var tags Tags
	result := r.db.Model(&post).Association("Tags").Find(&tags)
	return tags, result
}

func (r repository) AddTag(post Post, tag Tag) error {
	return r.db.Model(&post).Association("Tags").Append(&tag)
}

func (r repository) RemoveTag(post Post, tag Tag) error {
	return r.db.Model(&post).Association("Tags").Delete(&tag)
}
