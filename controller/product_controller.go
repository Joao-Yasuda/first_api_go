package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: productUseCase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductById(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }
    product, err := p.productUseCase.GetProductById(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
    var product model.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    if product.Name == "" || product.Price < 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product name cannot be empty and price cannot be negative"})
        return
    }

    if err := p.productUseCase.CreateProduct(product); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product: " + err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, product)
}

func (p *ProductController) UpdateProduct(ctx *gin.Context){
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if product.Name == "" || product.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product name cannot be empty and price cannot be negative"})
		return
	}

	if err := p.productUseCase.UpdateProduct(id, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product: " + err.Error()})
		return
	}

}


func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := p.productUseCase.DeleteProduct(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
