package solrController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wesolr/dto"
	"wesolr/services"
	client "wesolr/services/repositories"
)

var (
	solr = services.NewSolrServiceImpl(
		client.NewSolrClient("localhost", 8983, "items"),
	)
)

func GetQuery(c *gin.Context) {

	var itemsDto dto.ItemsDto
	query := c.Param("searchQuery")
	strs := strings.Split(query, "=")
	itemsDto, err := solr.GetQuery(strs[1], strs[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, itemsDto)
}
