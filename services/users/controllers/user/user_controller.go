package userController

import (
	"net/http"
	"strconv"
	"users/dto"
	service "users/services"
	client "users/services/repositories"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"users/config"
)

var (
	userService = service.NewUserServiceImpl(
		client.NewUserInterface(config.SQLUSER, config.SQLPASS, config.SQLHOST, config.SQLPORT, config.SQLDB),
		client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT),
	)
)

func GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	// Get Back User

	var userDto dto.UserDto
	id, _ := strconv.Atoi(c.Param("id"))
	userDto, err := userService.GetUserById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	var usersDto dto.UsersDto
	usersDto, err := userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	log.Debug(userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := userService.InsertUser(userDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	er := c.BindJSON(&loginDto)

	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	log.Debug(loginDto)

	var loginResponseDto dto.LoginResponseDto
	loginResponseDto, err := userService.Login(loginDto)
	if err != nil {
		if err.Status() == 400 {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponseDto)
}

func DeleteUser(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	err := userService.DeleteUser(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
