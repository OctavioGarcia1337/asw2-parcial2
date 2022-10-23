package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/model"
	"testing"
)

type CategoryClientInterface struct {
	mock.Mock
}

func (m *CategoryClientInterface) GetCategoryById(id int) model.Category {
	ret := m.Called(id)
	return ret.Get(0).(model.Category)
}

func (m *CategoryClientInterface) GetCategories() model.Categories {
	ret := m.Called()
	return ret.Get(0).(model.Categories)
}

func TestGetCategories(t *testing.T) {
	mockClient := new(CategoryClientInterface)

	var categories model.Categories
	var category model.Category
	categories = append(categories, category)

	mockClient.On("GetCategories").Return(categories)
	service := initCategoryService(mockClient)
	res, err := service.GetCategories()
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, len(res))

}

func TestGetCategoryById(t *testing.T) {
	mockClient := new(CategoryClientInterface)
	var category model.Category
	category.CategoryId = 1

	var categoryBad model.Category
	categoryBad.CategoryId = 0

	mockClient.On("GetCategoryById", 1).Return(category)
	mockClient.On("GetCategoryById", 0).Return(categoryBad)
	service := initCategoryService(mockClient)

	_, err := service.GetCategoryById(1)
	assert.Nil(t, err, "Error should be nil")

	_, err2 := service.GetCategoryById(0)
	assert.NotNil(t, err2, "Error should not be nil")

}
