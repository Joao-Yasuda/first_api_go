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

func (pr *ProductRepository) GetProductById(id int) (model.Product, error) {
	var products model.Product
	result := pr.db.Where("id = ?", id).Find(&products)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) error {
	result := pr.db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) UpdateProduct(id int, product model.Product) error {
	result := pr.db.Model(&model.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) DeleteProduct(id int) error{
	result := pr.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
		
	}
	return nil
}
