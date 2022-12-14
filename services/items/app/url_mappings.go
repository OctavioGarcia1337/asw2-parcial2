package app

import (
	itemController "items/controllers/item"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Items Mapping
	router.GET("/items/:item_id", itemController.GetItemById)
	router.POST("/item", itemController.InsertItem)
	router.POST("/items", itemController.QueueItems)

	router.DELETE("/users/:id/items", itemController.DeleteUserItems)

	log.Info("Finishing mappings configurations")
}
