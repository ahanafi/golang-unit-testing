package service

import (
	"errors"
	"github.com/ahanafi/golang-unit-testing/entity"
	"github.com/ahanafi/golang-unit-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found!")
	} else {
		return category, nil
	}
}
