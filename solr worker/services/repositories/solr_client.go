package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
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
	q := solr.NewQuery(solr.NewDisMaxQueryParser().
		Query(query).BuildParser()).
		Queries(solr.M{}).
		Limit(10).
		Fields(field)

	// Send the query
	queryResponse, err := sc.Client.Query(context.TODO(), sc.Collection, q)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Error processing query")
	}
	for item := range queryResponse.Response.Documents {
		var itemDto dto.ItemDto
		str, err := json.Marshal(queryResponse.Response.Documents[item])
		if err != nil {
			return itemsDto, e.NewBadRequestApiError("Failed marshal process")
		}
		err = json.Unmarshal(str, &itemDto)
		if err != nil {
			return itemsDto, e.NewBadRequestApiError("Failed unmarshal process")
		}
		itemsDto = append(itemsDto, itemDto)
	}
	return itemsDto, nil
}
