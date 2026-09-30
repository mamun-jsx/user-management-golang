package repository

import (
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
