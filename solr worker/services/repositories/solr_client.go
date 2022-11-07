package repositories

import (
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
	"net/http"
	"wesolr/dto"
	e "wesolr/utils/errors"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	logger.Debug(fmt.Sprintf("%s:%d", host, port))
	Client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d", host, port))
	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}

func (sc *SolrClient) GetQuery(query string, field string) (dto.ItemsDto, e.ApiError) {
	var itemsDto dto.ItemsDto
	q, err := http.Get("http://localhost:8983/solr/items/select?q=" + field + "%3A" + query)

	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error getting from solr")
	}

	var test []byte
	logger.Debug(q.Body.Read(test))
	q.Body.Read(test)
	err = json.Unmarshal(test, &itemsDto)
	logger.Debug(itemsDto)

	return itemsDto, nil
}
