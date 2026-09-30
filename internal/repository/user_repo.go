package repository

import (
	"errors"

	"github.com/mamun-jsx/user-management-golang/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// make an instance of UserRepository

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// ! Create a new User to database

func (r *UserRepository) Create(user *models.UserModel) error {
	return r.db.Create(user).Error
}

// find All user from database

func (r *UserRepository) FirstAll() ([]models.UserModel, error) {
	var allUsers []models.UserModel
	err := r.db.Find(&allUsers).Error
	return allUsers, err
}

// find a single user by email

func (r *UserRepository) FindUserByEmail(email string) (*models.UserModel, error) {
	var userByEmail models.UserModel
	err := r.db.First(&userByEmail, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userByEmail, nil
}

// ? delete a single user from database

func (r *UserRepository) delete(email string) (*models.UserModel, error) {
	var user models.UserModel
	// search the user an return the memory location

	err := r.db.Where(&user, "email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// now delete the record from database
	if err := r.db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
