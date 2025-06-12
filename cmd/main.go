package main

import (
	"go_api/controller"
	"go_api/db"
	"go_api/repository"
	"go_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	ProductController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8080")
}
