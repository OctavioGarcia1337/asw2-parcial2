package solrController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wesolr/config"
	"wesolr/dto"
	"wesolr/services"
	client "wesolr/services/repositories"
	con "wesolr/utils/connections"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
		(*client.QueueClient)(con.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT)),
	)
)

func GetQuery(c *gin.Context) {
	var itemsDto dto.ItemsDto
	query := c.Param("searchQuery")

	itemsDto, err := Solr.GetQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
	}

	c.JSON(http.StatusOK, itemsDto)

}
