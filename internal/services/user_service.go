package services

import (
	"CleanArchitecture/internal/models"
	"CleanArchitecture/internal/repositories"
)

type UserService interface {
	AddUser(user models.User) error
	FindUser(id int) (models.User, error)
	RemoveUser(id int) error
}

type DefaultUserService struct {
	repo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *DefaultUserService {
	return &DefaultUserService{repo: repo}
}

func (s *DefaultUserService) AddUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *DefaultUserService) FindUser(id int) (models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *DefaultUserService) RemoveUser(id int) error {
	return s.repo.DeleteUserByID(id)
}
