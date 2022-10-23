package orderController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	log.Debug("Order id: " + c.Param("id"))

	var orderDto dto.OrderDto

	c.JSON(http.StatusOK, orderDto)
}

func GetOrdersByUserId(c *gin.Context) {
	log.Debug("UserId: " + c.Param("id"))

	var ordersDto dto.OrdersDto
	id, _ := strconv.Atoi(c.Param("id"))
	ordersDto, err := service.OrderService.GetOrdersByUserId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(ordersDto) == 0 {
		c.JSON(http.StatusOK, []dto.OrderDto{})
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}

func OrderInsert(c *gin.Context) {
	var orderInsertDto dto.OrderInsertDto
	var orderResponseDto dto.OrderResponseDto
	err := c.BindJSON(&orderInsertDto)

	log.Debug(orderInsertDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderResponseDto, er := service.OrderService.InsertOrder(orderInsertDto)
	if er != nil {
		log.Error(er.Error())
		c.JSON(er.Status(), er.Error())
		return
	}
	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		_, e := service.OrderDetailService.InsertDetail(orderInsertDto.OrderDetails[i], orderResponseDto.OrderId)
		if e != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	c.JSON(http.StatusCreated, orderInsertDto)
}
