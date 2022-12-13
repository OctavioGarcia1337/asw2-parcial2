package messageController

import (
	"messages/dto"
	service "messages/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetMessageById(c *gin.Context) {
	log.Debug("Message id: " + c.Param("id"))

	// Get Back Message

	var messageDto dto.MessageDto
	id, _ := strconv.Atoi(c.Param("id"))
	messageDto, err := service.MessageService.GetMessageById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, messageDto)
}

func GetMessagesByUserId(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	// Get Back Messages

	var messagesDto dto.MessagesDto
	id, _ := strconv.Atoi(c.Param("id"))
	messagesDto, err := service.MessageService.GetMessagesByUserId(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, messagesDto)
}

func GetMessages(c *gin.Context) {

	var messagesDto dto.MessagesDto
	messagesDto, err := service.MessageService.GetMessages()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, messagesDto)
}

func MessageInsert(c *gin.Context) {
	var messageDto dto.MessageDto
	err := c.BindJSON(&messageDto)

	log.Debug(messageDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	messageDto, er := service.MessageService.InsertMessage(messageDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, messageDto)
}
