package repositories

import (
    "library-management-system/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    err := r.DB.First(&user, id).Error
    return &user, err
}

func (r *UserRepository) DeleteUser(id uint) error {
    return r.DB.Delete(&models.User{}, id).Error
}