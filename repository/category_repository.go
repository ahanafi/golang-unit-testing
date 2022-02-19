package repository

import "github.com/ahanafi/golang-unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}