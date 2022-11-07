package repositories

import (
	"bytes"
	"context"
	json "encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
	"items/dto"
	e "items/utils/errors"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

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

type docTest struct {
	Doc dto.ItemDto `json:"doc"`
}

type test struct {
	Add docTest `json:"add"`
}

func (sc *SolrClient) Update(itemDto dto.ItemDto, command string) e.ApiError {
	var addItemDto test
	addItemDto.Add = docTest{itemDto}
	data, err := json.Marshal(addItemDto)

	reader := bytes.NewReader(data)
	if err != nil {
		failOnError(err, "Failed to get json")
	}
	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp.Error)
	if err != nil {
		logger.Debug(err)
		return e.NewBadRequestApiError("Error in solr")
	}

	sc.Client.Commit(context.TODO(), sc.Collection)
	return nil
}
