package main

import (
	"gorm.io/gorm"
)

type TagRepository interface {
	Get(id string) (Tag, error)
	Create(tag Tag) error
	Delete(id string) error
	Update(tag Tag) error
	Query(tag Tag) (Tags, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) tagRepository {
	return tagRepository{db}
}

func (r tagRepository) Get(id string) (Tag, error) {
	var tag Tag
	result := r.db.First(&tag, id)
	return tag, result.Error
}

func (r tagRepository) Create(tag Tag) error {
	return r.db.Create(&tag).Error
}

func (r tagRepository) Delete(id string) error {
	return r.db.Delete(&Tag{}, id).Error
}

func (r tagRepository) Update(tag Tag) error {
	return r.db.Save(&tag).Error
}

func (r tagRepository) Query(tag Tag) (Tags, error) {
	var tags Tags
	result := r.db.Where(&tag).Find(&tags)
	return tags, result.Error
}
