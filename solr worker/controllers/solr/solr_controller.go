package solrController

import (
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"io/ioutil"
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
	strs := strings.Split(query, "_")

	q, _ := http.Get("http://localhost:8983/solr/items/select?q=" + strs[0] + "%3A" + strs[1])
	body, _ := ioutil.ReadAll(q.Body)

	var test dto.SolrResponseDto
	err := json.Unmarshal(body, &test)

	itemsDto = test.Response.Docs
	// itemsDto, err := solr.GetQuery(strs[1], strs[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, itemsDto)

}
