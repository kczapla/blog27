package main

import (
    "gorm.io/gorm"
)

type UserRepository interface {
    Get(id string) (User, error)
    Create(user User) error
    Delete(id string) error
    Update(user User) error
    Query() (Users, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepository {
    return userRepository{db}
}

func (r userRepository) Get(id string) (User, error) {
    var user User
    result := r.db.First(&user, id)
    return user, result.Error
}

func (r userRepository) Create(user User) error {
    return r.db.Create(&user).Error
}

func (r userRepository) Delete(id string) error {
    return r.db.Delete(&User{}, id).Error
}

func (r userRepository) Update(user User) error {
    return r.db.Save(&user).Error
}

func (r userRepository) Query() (Users, error) {
    var users Users
    result := r.db.Find(&users)
    return users, result.Error
}

