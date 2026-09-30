package service

import (
	"errors"
	"github.com/mamun-jsx/user-management-golang/internal/models"
	"github.com/mamun-jsx/user-management-golang/internal/repository"
	"strings"
)

type UserService struct {
	repo *repository.UserRepository
}

// create an instance of User Service

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// create a user

func (s *UserService) CreateUser(user *models.UserModel) error {
	// 1. Business Validation: Validate Name
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("user name is required") // Fixed "book name" typo
	}
	// 2. Business Validation: Validate Email
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}
	// 3. Business Logic: Check if user already exists by email
	existingUser, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		// If it's a real DB error (not "record not found"), return it
		// Assuming FindUserByEmail handles gorm.ErrRecordNotFound by returning nil, nil
		return err
	}
	if existingUser != nil {
		return errors.New("email is already registered")
	}

	// 4. Set Default Business Values if needed
	if user.Role == "" {
		user.Role = "customer" // Assign a default role if empty
	}
	// 5. Call the Repository to save the record to the database
	// Assuming your repo has a Create method similar to the ProductRepository
	err = s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
