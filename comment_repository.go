package main

import (
    "gorm.io/gorm"
)

type CommentRepository interface {
    Get(id string) (Comment, error)
    Create(comment Comment) error
    Delete(id string) error
    Update(comment Comment) error
    Query() (Comments, error)
}

type commentRepository struct {
    db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) commentRepository {
    return commentRepository{db}
}

func (r commentRepository) Get(id string) (Comment, error) {
    var comment Comment
    result := r.db.First(&comment, id)
    return comment, result.Error
}

func (r commentRepository) Create(comment Comment) error {
    return r.db.Create(&comment).Error
}

func (r commentRepository) Delete(id string) error {
    return r.db.Delete(&Comment{}, id).Error
}

func (r commentRepository) Update(comment Comment) error {
    return r.db.Save(&comment).Error
}

func (r commentRepository) Query() (Comments, error) {
    var comments Comments
    result := r.db.Find(&comments)
    return comments, result.Error
}

