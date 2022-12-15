package app

import (
	messageController "messages/controllers/message"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Messages Mapping
	router.GET("/messages/:id", messageController.GetMessageById)
	router.GET("/users/:id/messages", messageController.GetMessagesByUserId)
	router.GET("/messages", messageController.GetMessages)

	router.DELETE("/messages/:id", messageController.DeleteMessageById)

	router.POST("/message", messageController.MessageInsert)

	log.Info("Finishing mappings configurations")
}
