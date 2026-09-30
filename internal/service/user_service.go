package service

import "github.com/mamun-jsx/user-management-golang/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

// create an instance of User Service

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// create a user 
