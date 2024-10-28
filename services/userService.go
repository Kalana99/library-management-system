package services

import (
    "library-management-system/models"
    "library-management-system/repositories"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    return s.repo.GetUserByID(id)
}

func (s *UserService) DeleteUser(id uint) error {
    return s.repo.DeleteUser(id)
}
