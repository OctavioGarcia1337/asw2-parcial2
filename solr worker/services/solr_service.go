package services

import (
	"wesolr/dto"
	client "wesolr/services/repositories"
	e "wesolr/utils/errors"
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
) *SolrService {
	return &SolrService{
		solr: solr,
	}
}
func (s *SolrService) GetQuery(query string, field string) (dto.ItemsDto, e.ApiError) {
	var itemsDto dto.ItemsDto
	itemsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return itemsDto, e.NewBadRequestApiError("Solr failed")
	}
	return itemsDto, nil
}
