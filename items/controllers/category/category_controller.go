package categoryController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(c *gin.Context) {
	log.Debug("Category id: " + c.Param("id"))

	// Get Back User

	id, _ := strconv.Atoi(c.Param("id"))
	var categoryDto dto.CategoryDto
	categoryDto, err := service.CategoryService.GetCategoryById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categoryDto)
}

func GetCategories(c *gin.Context) {
	var categoriesDto dto.CategoriesDto
	categoriesDto, err := service.CategoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categoriesDto)

}
