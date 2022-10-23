package cartController

import (
	"mvc-go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCartById(c *gin.Context) {
	log.Debug("Cart id: " + c.Param("id"))

	// Get Back User
	var cartDto dto.CartDto
	c.JSON(http.StatusOK, cartDto)
}

func InsertCart(c *gin.Context) {

	return
}
