package app

import (
	log "github.com/sirupsen/logrus"
	solr_controller "wesolr/controllers/solr"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search=:searchQuery", solr_controller.GetQuery)
	router.GET("/searchAll=:searchQuery", solr_controller.GetQueryAllFields)
	router.GET("/items/:id", solr_controller.AddFromId)

	log.Info("Finishing mappings configurations")
}
