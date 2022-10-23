package addressController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAddressById(c *gin.Context) {
	log.Debug("Address id: " + c.Param("id"))

	var addressDto dto.AddressDto
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, addressDto)
		return
	}
	addressDto, er := service.AddressService.GetAddressById(id)

	if er != nil {
		c.JSON(http.StatusBadRequest, addressDto)
		return
	}
	c.JSON(http.StatusOK, addressDto)
}

func GetAddressesByUserId(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	var addressesDto dto.AddressesDto
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, addressesDto)
		return
	}

	addressesDto, er := service.AddressService.GetAddressesByUserId(id)
	if er != nil {
		c.JSON(http.StatusBadRequest, addressesDto)
		return
	}

	if len(addressesDto) == 0 {
		c.JSON(http.StatusOK, []dto.AddressDto{})
		return
	}

	c.JSON(http.StatusOK, addressesDto)
}

func AddressInsert(c *gin.Context) {

	var addressDto dto.AddressDto
	err := c.BindJSON(&addressDto)

	log.Debug(addressDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, addressDto)
}
