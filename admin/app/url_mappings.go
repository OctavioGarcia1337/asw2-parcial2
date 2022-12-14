package app

import (
	dockerController "admin/controllers/docker"
	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.POST("/container/:image", dockerController.CreateContainer)

	log.Info("Finishing mappings configurations")
}
