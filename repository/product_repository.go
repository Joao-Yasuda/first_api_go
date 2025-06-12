package repository

import (
	"go_api/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	result := pr.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
