package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
)

type ProductClientInterface struct {
	mock.Mock
}

func (m *ProductClientInterface) GetProductById(id int) model.Product {
	ret := m.Called(id)
	return ret.Get(0).(model.Product)
}
func (m *ProductClientInterface) GetProducts() model.Products {
	ret := m.Called()
	return ret.Get(0).(model.Products)
}
func (m *ProductClientInterface) GetNProducts(n int) model.Products {
	ret := m.Called(n)
	return ret.Get(0).(model.Products)
}
func (m *ProductClientInterface) RemoveStock(id int, amount int) model.Product {
	ret := m.Called(id, amount)
	return ret.Get(0).(model.Product)
}
func (m *ProductClientInterface) GetProductsByCategoryId(id int) model.Products {
	ret := m.Called(id)
	return ret.Get(0).(model.Products)
}
func (m *ProductClientInterface) GetProductsBySearch(query string) model.Products {
	ret := m.Called(query)
	return ret.Get(0).(model.Products)
}

func TestGetProductById(t *testing.T) {
	mockProductClient := new(ProductClientInterface)
	mockCategoryClient := new(CategoryClientInterface)

	var product model.Product
	product.ProductId = 1
	product.CategoryId = 1
	product.Name = "Test_Product"
	product.Description = "Test_Desc"
	product.Price = 500
	product.CurrencyId = "ARS"
	product.Stock = 5
	product.Picture = "test.png"

	var categoryDto dto.CategoryDto
	categoryDto.CategoryId = 1
	categoryDto.Name = "Test_Category"
	categoryDto.Description = "Test_Description"

	var category model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var productDto dto.ProductDto
	productDto.ProductId = 1
	productDto.Category = categoryDto
	productDto.Name = "Test_Product"
	productDto.Description = "Test_Desc"
	productDto.Price = 500
	productDto.CurrencyId = "ARS"
	productDto.Stock = 5
	productDto.Picture = "test.png"

	mockProductClient.On("GetProductById", 1).Return(product)
	mockCategoryClient.On("GetCategoryById", 1).Return(category)
	service := initProductService(mockProductClient, mockCategoryClient)
	res, err := service.GetProductById(1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, res, productDto)
}

func TestGetProducts(t *testing.T) {
	mockProductClient := new(ProductClientInterface)
	mockCategoryClient := new(CategoryClientInterface)

	var category model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var product model.Product
	product.ProductId = 1
	product.CategoryId = 1
	product.Name = "Test_Product"
	product.Description = "Test_Desc"
	product.Price = 500
	product.CurrencyId = "ARS"
	product.Stock = 5
	product.Picture = "test.png"
	var products model.Products
	products = append(products, product)

	mockProductClient.On("GetProducts").Return(products)
	mockCategoryClient.On("GetCategoryById", 1).Return(category)
	service := initProductService(mockProductClient, mockCategoryClient)

	res, err := service.GetProducts()
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, len(res)) // Products shouldn't be empty
}

func TestGetProductsByCategoryId(t *testing.T) {
	mockProductClient := new(ProductClientInterface)
	mockCategoryClient := new(CategoryClientInterface)

	var category model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var categoryTwo model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var productCatOne model.Product
	productCatOne.ProductId = 1
	productCatOne.CategoryId = 1
	productCatOne.Name = "Test_Product_1"
	productCatOne.Description = "Test_Desc"
	productCatOne.Price = 500
	productCatOne.CurrencyId = "ARS"
	productCatOne.Stock = 5
	productCatOne.Picture = "test.png"
	var productCatTwo model.Product
	productCatTwo.ProductId = 1
	productCatTwo.CategoryId = 2
	productCatTwo.Name = "Test_Product_2"
	productCatTwo.Description = "Test_Desc"
	productCatTwo.Price = 500
	productCatTwo.CurrencyId = "ARS"
	productCatTwo.Stock = 5
	productCatTwo.Picture = "test.png"
	var productsCatOne model.Products
	productsCatOne = append(productsCatOne, productCatOne)
	var productsCatTwo model.Products
	productsCatTwo = append(productsCatTwo, productCatTwo)
	var productsCatThree model.Products

	mockCategoryClient.On("GetCategoryById", 1).Return(category)
	mockCategoryClient.On("GetCategoryById", 2).Return(categoryTwo)
	mockProductClient.On("GetProductsByCategoryId", 1).Return(productsCatOne)
	mockProductClient.On("GetProductsByCategoryId", 2).Return(productsCatTwo)
	mockProductClient.On("GetProductsByCategoryId", 3).Return(productsCatThree)
	service := initProductService(mockProductClient, mockCategoryClient)
	res, err := service.GetProductsByCategoryId(1)
	res2, err2 := service.GetProductsByCategoryId(2)
	res3, err3 := service.GetProductsByCategoryId(3)

	assert.Nil(t, err, "Error should be nil")
	assert.Nil(t, err2, "Error should be nil")
	assert.Nil(t, err3, "Error should be nil")

	assert.NotEqual(t, 0, len(res))
	assert.NotEqual(t, 0, len(res2))
	assert.Equal(t, 0, len(res3)) // Should be empty
}

func TestGetProductsBySearch(t *testing.T) {
	mockProductClient := new(ProductClientInterface)
	mockCategoryClient := new(CategoryClientInterface)

	var category model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var productWood model.Product
	productWood.ProductId = 1
	productWood.CategoryId = 1
	productWood.Name = "Test_Product_Wood"
	productWood.Description = "Test_Desc"
	productWood.Price = 500
	productWood.CurrencyId = "ARS"
	productWood.Stock = 5
	productWood.Picture = "test.png"

	var results model.Products
	results = append(results, productWood)
	var noresults model.Products

	mockCategoryClient.On("GetCategoryById", 1).Return(category)
	mockProductClient.On("GetProductsBySearch", "Results").Return(results)
	mockProductClient.On("GetProductsBySearch", "").Return(noresults)
	mockProductClient.On("GetProductsBySearch", "NoResults").Return(noresults)
	service := initProductService(mockProductClient, mockCategoryClient)

	res, err := service.GetProductsBySearch("Results")
	res2, err2 := service.GetProductsBySearch("")
	res3, err3 := service.GetProductsBySearch("NoResults")

	assert.Nil(t, err, "Error should be Nil")
	assert.Nil(t, err2, "Error should be Nil")
	assert.Nil(t, err3, "Error should be Nil")

	assert.NotEqual(t, 0, len(res))
	assert.Equal(t, 0, len(res2))
	assert.Equal(t, 0, len(res3))
}

func TestGetNProducts(t *testing.T) {
	mockProductClient := new(ProductClientInterface)
	mockCategoryClient := new(CategoryClientInterface)

	var category model.Category
	category.CategoryId = 1
	category.Name = "Test_Category"
	category.Description = "Test_Description"

	var productWood model.Product
	productWood.ProductId = 1
	productWood.CategoryId = 1
	productWood.Name = "Test_Product_Wood"
	productWood.Description = "Test_Desc"
	productWood.Price = 500
	productWood.CurrencyId = "ARS"
	productWood.Stock = 5
	productWood.Picture = "test.png"

	var results model.Products
	results = append(results, productWood)
	results = append(results, productWood)
	results = append(results, productWood)
	results = append(results, productWood)
	results = append(results, productWood)

	mockCategoryClient.On("GetCategoryById", 1).Return(category)
	mockProductClient.On("GetNProducts", 5).Return(results)
	service := initProductService(mockProductClient, mockCategoryClient)

	res, err := service.GetNProducts(5)
	assert.Nil(t, err, "Error should not be Nil")

	assert.NotEqual(t, 0, len(res)) // shouldn't be empty
	assert.Equal(t, 5, len(res))    // should be equal to the input
}
