package app

import (
	log "github.com/sirupsen/logrus"
	solrController "wesolr/controllers/solr"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search=:searchQuery", solrController.GetQuery)
	router.GET("/searchAll=:searchQuery", solrController.GetQueryAllFields)
	router.GET("/items/:id", solrController.AddFromId)

	router.DELETE("/items/:id", solrController.Delete)

	log.Info("Finishing mappings configurations")
}
