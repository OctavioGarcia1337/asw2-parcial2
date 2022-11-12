package services

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"wesolr/config"
	"wesolr/dto"
	client "wesolr/services/repositories"
	e "wesolr/utils/errors"
)

type SolrService struct {
	solr  *client.SolrClient
	queue *client.QueueClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
	queue *client.QueueClient,
) *SolrService {
	return &SolrService{
		solr:  solr,
		queue: queue,
	}
}
func (s *SolrService) GetQuery(query string) (dto.ItemsDto, e.ApiError) {
	var itemsDto dto.ItemsDto
	queryParams := strings.Split(query, "_")
	query, field := queryParams[0], queryParams[1]
	itemsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Solr failed")
	}
	return itemsDto, nil
}

func (s *SolrService) Add(itemDto dto.ItemDto) {
	s.solr.Update(itemDto, "add")
}

func (s *SolrService) QueueWorker(qname string) {
	s.queue.ProcessMessages(qname, func(id string) {
		var itemDto dto.ItemDto
		resp, err := http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.ITEMSHOST, config.ITEMSPORT, id))
		if err != nil {
			log.Debugf("error getting item %s", id)
			return
		}
		var body []byte
		body, _ = io.ReadAll(resp.Body)
		//resp.Body.Read(body)
		log.Debugf("%s", body)
		err = json.Unmarshal(body, &itemDto)
		if err != nil {
			log.Debugf("error in unmarshal of item %s", id)
			return
		}
		s.Add(itemDto)
	})
}
