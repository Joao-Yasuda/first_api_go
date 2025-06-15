package main

import (
	"go_api/controller"
	"go_api/db"
	"go_api/model"
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

	err = dbConnection.AutoMigrate(&model.Person{}, &model.Product{})
    if err != nil {
        panic("Failed to migrate database: " + err.Error())
    }

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	ProductController := controller.NewProductController(productUseCase)

	personRepository := repository.NewPersonRepository(dbConnection)
	personUseCase := usecase.NewPersonUseCase(personRepository)
	PersonController := controller.NewPersonController(personUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:id", ProductController.GetProductById)
	server.DELETE("/products/:id", ProductController.DeleteProduct)
	server.PUT("/products/:id", ProductController.UpdateProduct)

	server.GET("/person", PersonController.GetPerson)
	server.POST("/person", PersonController.CreatePerson)
	server.Run(":8080")
}
