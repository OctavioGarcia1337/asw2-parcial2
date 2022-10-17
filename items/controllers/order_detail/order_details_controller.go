package orderDetailController

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mvc-go/dto"
	"net/http"
)

func GetOrderDetailById(c *gin.Context) {
	log.Debug("Order Detail id: " + c.Param("id"))

	var orderDetailDto dto.OrderDetailDto
	c.JSON(http.StatusOK, orderDetailDto)
}

func OrderDetailInsert(c *gin.Context) {
	var orderDetailDto dto.OrderDetailDto
	err := c.BindJSON(&orderDetailDto)

	log.Debug(orderDetailDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, orderDetailDto)
}
