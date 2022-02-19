package service

import (
	"github.com/ahanafi/golang-unit-testing/entity"
	"github.com/ahanafi/golang-unit-testing/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetFail(t *testing.T) {

	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "10",
		Name: "Programming",
	}
	categoryRepository.Mock.On("FindById", "10").Return(category)
	categoryResult, err := categoryService.Get("10")
	assert.Nil(t, err)
	assert.NotNil(t, categoryResult)
	assert.Equal(t, category.Id, categoryResult.Id)
	assert.Equal(t, category.Name, categoryResult.Name)
}