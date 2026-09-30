package repository

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

// make an instance of product Repository

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// create a new Product to database

