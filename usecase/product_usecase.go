package usecase

import (
	"go_api/model"
)

type ProductRepository interface {
	GetProducts() ([]model.Product, error)
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
