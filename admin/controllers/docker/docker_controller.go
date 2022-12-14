package solrController

import (
	"admin/services"
	client "admin/services/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	Docker = services.NewDockerServiceImpl(
		client.NewDockerClient(),
	)
)

func CreateContainer(c *gin.Context) {

	image := c.Param("image")
	log.Debug(image)
	id, err := Docker.CreateContainer(image, "test")
	if err != nil {
		c.JSON(http.StatusBadRequest, id)
		return
	}

	c.JSON(http.StatusOK, id)

}
