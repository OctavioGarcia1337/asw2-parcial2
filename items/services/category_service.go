package services

import (
	client "mvc-go/clients/category"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"
)

type categoryService struct {
	categoryClient client.CategoryClientInterface
}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func initCategoryService(categoryClient client.CategoryClientInterface) categoryServiceInterface {
	service := new(categoryService)
	service.categoryClient = categoryClient
	return service
}

func init() {
	CategoryService = initCategoryService(client.CategoryClient)
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, e.ApiError) {

	var category model.Category = s.categoryClient.GetCategoryById(id)
	var categoryDto dto.CategoryDto

	if category.CategoryId <= 0 {
		return categoryDto, e.NewBadRequestApiError("Category not found")
	}
	categoryDto.Description = category.Description
	categoryDto.Name = category.Name
	categoryDto.CategoryId = category.CategoryId
	return categoryDto, nil
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = s.categoryClient.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Description = category.Description
		categoryDto.Name = category.Name
		categoryDto.CategoryId = category.CategoryId

		categoriesDto = append(categoriesDto, categoryDto)
	}

	log.Debug(categoriesDto)
	return categoriesDto, nil
}
