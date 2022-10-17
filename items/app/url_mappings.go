package app

import (
	productController "mvc-go/controllers/product"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/items/:item_id", productController.GetProductById)
	router.POST("/product", productController.InsertProduct)

	log.Info("Finishing mappings configurations")
}
