package services

import (
	cclient "mvc-go/clients/category"
	pclient "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"
)

type productService struct {
	productClient  pclient.ProductClientInterface
	categoryClient cclient.CategoryClientInterface
}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError)
	GetProductsBySearch(query string) (dto.ProductsDto, e.ApiError)
	GetNProducts(n int) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func initProductService(productClient pclient.ProductClientInterface, categoryClient cclient.CategoryClientInterface) productServiceInterface {
	service := new(productService)
	service.productClient = productClient
	service.categoryClient = categoryClient
	return service
}

func init() {
	ProductService = initProductService(pclient.ProductClient, cclient.CategoryClient)
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = s.productClient.GetProductById(id)
	var productDto dto.ProductDto

	if product.ProductId < 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	category := s.categoryClient.GetCategoryById(id)

	productDto.ProductId = product.ProductId
	productDto.Category.CategoryId = category.CategoryId
	productDto.Category.Name = category.Name
	productDto.Category.Description = category.Description
	productDto.Name = product.Name
	productDto.Description = product.Description
	productDto.Price = product.Price
	productDto.CurrencyId = product.CurrencyId
	productDto.Stock = product.Stock
	productDto.Picture = product.Picture

	return productDto, nil
}

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		category := s.categoryClient.GetCategoryById(product.CategoryId)
		productDto.ProductId = product.ProductId
		productDto.Category.CategoryId = category.CategoryId
		productDto.Category.Name = category.Name
		productDto.Category.Description = category.Description
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetNProducts(n int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetNProducts(n)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		category := s.categoryClient.GetCategoryById(product.CategoryId)

		productDto.ProductId = product.ProductId
		productDto.Category.CategoryId = category.CategoryId
		productDto.Category.Name = category.Name
		productDto.Category.Description = category.Description
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetProductsByCategoryId(id int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = s.productClient.GetProductsByCategoryId(id)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		category := s.categoryClient.GetCategoryById(id)
		productDto.ProductId = product.ProductId
		productDto.Category.CategoryId = category.CategoryId
		productDto.Category.Name = category.Name
		productDto.Category.Description = category.Description
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}

func (s *productService) GetProductsBySearch(query string) (dto.ProductsDto, e.ApiError) {
	var products model.Products
	products = s.productClient.GetProductsBySearch(query)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		category := s.categoryClient.GetCategoryById(product.CategoryId)
		productDto.ProductId = product.ProductId
		productDto.Category.CategoryId = category.CategoryId
		productDto.Category.Name = category.Name
		productDto.Category.Description = category.Description
		productDto.Name = product.Name
		productDto.Description = product.Description
		productDto.Price = product.Price
		productDto.CurrencyId = product.CurrencyId
		productDto.Stock = product.Stock
		productDto.Picture = product.Picture

		productsDto = append(productsDto, productDto)
	}

	log.Debug(productsDto)
	return productsDto, nil
}
