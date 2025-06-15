package usecase

import (
	"go_api/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProductById(id int) (model.Product, error)
	CreateProduct(product model.Product) error
	UpdateProduct(id int, product model.Product) error
	DeleteProduct(id int) error
}

type ProductUseCase struct {
	productRepository ProductRepository
}

func NewProductUseCase(productRepository ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: productRepository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUseCase) GetProductById(id int) (model.Product, error){
	return pu.productRepository.GetProductById(id)
}

func (pu *ProductUseCase) CreateProduct(product model.Product) error {
	return pu.productRepository.CreateProduct(product) 
}

func (pu *ProductUseCase) UpdateProduct(id int, product model.Product) error {
	return pu.productRepository.UpdateProduct(id, product)
}

func (pu *ProductUseCase) DeleteProduct(id int) error{
	return pu.productRepository.DeleteProduct(id)
}
