package app

import (
	messageController "messages/controllers/message"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Messages Mapping
	router.GET("/messages/:id", messageController.GetMessageById)
	router.GET("/messages", messageController.GetMessages)
	router.POST("/message", messageController.MessageInsert)
	router.GET("/users/:id/messages", messageController.GetMessagesByUserId)

	log.Info("Finishing mappings configurations")
}
