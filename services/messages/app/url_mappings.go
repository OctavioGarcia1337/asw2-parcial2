package app

import (
	messageController "messages/controllers/message"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Messages Mapping
	router.GET("/message/:id", messageController.GetMessageById)
	router.GET("/message", messageController.GetMessages)
	router.POST("/message", messageController.MessageInsert)
	router.GET("/users/:id/messages", messageController.GetMessagesByUserId)

	log.Info("Finishing mappings configurations")
}
