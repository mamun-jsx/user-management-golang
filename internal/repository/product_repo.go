package repository

import (
	"github.com/mamun-jsx/user-management-golang/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

// make an instance of product Repository

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// create a new Product to database

func (r *ProductRepository) Create(product *models.ProductModel) error {
	return r.db.Create(product).Error
}
