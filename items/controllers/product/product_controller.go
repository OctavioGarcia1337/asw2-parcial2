package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	var productDto dto.ProductDto
	id, _ := strconv.Atoi(c.Param("product_id"))
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {

	var productsDto dto.ProductsDto
	var err error

	limit, ok := c.GetQuery("limit")
	n, _ := strconv.Atoi(limit)
	if ok {
		productsDto, err = service.ProductService.GetNProducts(n)
	} else {
		productsDto, err = service.ProductService.GetProducts()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByCategoryId(c *gin.Context) {

	var productsDto dto.ProductsDto
	id, _ := strconv.Atoi(c.Param("category_id"))
	productsDto, err := service.ProductService.GetProductsByCategoryId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsBySearch(c *gin.Context) {
	var productsDto dto.ProductsDto
	query := c.Param("searchQuery")
	productsDto, err := service.ProductService.GetProductsBySearch(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(productsDto) == 0 {
		c.JSON(http.StatusOK, []dto.ProductDto{})
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
