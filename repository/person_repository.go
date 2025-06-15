package repository

import (
	"go_api/model"

	"gorm.io/gorm"
)

	type PersonRepository struct{
		db *gorm.DB
	}

	func NewPersonRepository(db *gorm.DB) *PersonRepository{
		return &PersonRepository{
			db:db,
		}
	}

	func (pr *PersonRepository) GetPerson() ([]model.Person, error){
		var person []model.Person
		result := pr.db.Find(&person)
		if result.Error != nil{
			return nil, result.Error
		}
		return person, nil
	}

	func (pr *PersonRepository) CreatePerson(person model.Person) error{
		result := pr.db.Create(&person)
		if result.Error != nil{
			return result.Error
		}
		return nil
	}