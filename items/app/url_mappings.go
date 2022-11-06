package app

import (
	itemController "items/controllers/item"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/items/:item_id", itemController.GetItemById)
	router.POST("/item", itemController.InsertItem)

	log.Info("Finishing mappings configurations")
}
