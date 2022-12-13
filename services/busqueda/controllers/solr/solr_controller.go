package solrController

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
		log.Debug(itemsDto)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func GetQueryAllFields(c *gin.Context) {
	var itemsDto dto.ItemsDto
	query := c.Param("searchQuery")

	itemsDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(itemsDto)
		c.JSON(http.StatusBadRequest, itemsDto)
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func AddFromId(c *gin.Context) {
	id := c.Param("id")
	err := Solr.AddFromId(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, err)
}
