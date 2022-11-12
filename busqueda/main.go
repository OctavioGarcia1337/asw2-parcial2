package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	solr_controller "wesolr/controllers/solr"
)

func mapUrls() {
	// Products Mapping
	router.GET("/search=:searchQuery", solr_controller.GetQuery)

	log.Info("Finishing mappings configurations")

}

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	log.SetLevel(log.DebugLevel)
	router.Use(cors.Default())
}

func main() {
	mapUrls()
	log.Info("Starting server")
	solr_controller.Solr.QueueWorker("solr")
	router.Run(":8000")

}
