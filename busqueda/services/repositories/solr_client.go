package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	solr "github.com/stevenferrer/solr-go"
	"io"
	"net/http"
	"wesolr/config"
	"wesolr/dto"
	e "wesolr/utils/errors"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func (sc *SolrClient) GetQuery(query string, field string) (dto.ItemsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var itemsDto dto.ItemsDto
	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/items/select?q=%s%s%s", config.SOLRHOST, config.SOLRPORT, field, "%3A", query))
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error getting from solr")
	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("error in unmarshal")
	}

	itemsDto = response.Response.Docs
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
