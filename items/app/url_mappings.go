package app

import (
	itemController "items/controllers/item"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/items/:item_id", itemController.GetItemById)
	router.POST("/solr", itemController.InsertItem)
	router.POST("/items", itemController.QueueItems)

	log.Info("Finishing mappings configurations")
}
