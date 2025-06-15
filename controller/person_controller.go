package controller

import (
	"go_api/model"
	"go_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	personUsecase usecase.PersonUseCase
}

func NewPersonController(personUseCase usecase.PersonUseCase) PersonController {
	return PersonController{
		personUsecase: personUseCase,
	}
}

func (p *PersonController) GetPerson(ctx *gin.Context) {
	person, err := p.personUsecase.GetPerson()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, person)
}

func (p *PersonController) CreatePerson (ctx *gin.Context) {
	var person model.Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if person.Name == "" || person.Age < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Person name cannot be empty and age cannot be negative"})
		return
	}

	if err := p.personUsecase.CreatePerson(person); err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, person)
}	