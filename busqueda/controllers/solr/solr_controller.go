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

func AddFromId(c *gin.Context) {
	id := c.Param("id")
	err := Solr.AddFromId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, err)
}
