package repositories

import (
	"bytes"
	"context"
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

	var body []byte
	q.Body.Read(body)
	err = json.Unmarshal(body, &itemsDto)

	return itemsDto, nil
}

func (sc *SolrClient) Update(itemDto dto.ItemDto, command string) e.ApiError {
	var addItemDto dto.AddDto
	addItemDto.Add = dto.DocDto{Doc: itemDto}
	data, err := json.Marshal(addItemDto)

	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp.Error)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	sc.Client.Commit(context.TODO(), sc.Collection)
	return nil
}
