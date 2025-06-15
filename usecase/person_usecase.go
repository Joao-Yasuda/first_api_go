package usecase

import (
	"go_api/model"
)

type PersonRepository interface {
	GetPerson() ([]model.Person, error)
	CreatePerson(person model.Person) error
}

func NewPersonUseCase (PersonRepository PersonRepository) PersonUseCase{
	return PersonUseCase{
		personRepository: PersonRepository,
	}
}

type PersonUseCase struct{
	personRepository PersonRepository
}

func (p *PersonUseCase) GetPerson() ([]model.Person, error){
	return p.personRepository.GetPerson()
}

func (p *PersonUseCase) CreatePerson(person model.Person) error{
	return p.personRepository.CreatePerson(person)
}

