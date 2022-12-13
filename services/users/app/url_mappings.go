package app

import (
	userController "users/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/users/:id", userController.GetUserById)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	// Login Mapping
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
