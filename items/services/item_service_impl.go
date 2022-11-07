package services

import (
	log "github.com/sirupsen/logrus"
	"items/dto"
	client "items/services/repositories"
	e "items/utils/errors"
)

type ItemServiceImpl struct {
	item      *client.ItemClient
	memcached *client.MemcachedClient
	queue     *client.QueueClient
	solr      *client.SolrClient
}

func NewItemServiceImpl(
	item *client.ItemClient,
	memcached *client.MemcachedClient,
	queue *client.QueueClient,
	solr *client.SolrClient,
) *ItemServiceImpl {
	return &ItemServiceImpl{
		item:      item,
		memcached: memcached,
		queue:     queue,
		solr:      solr,
	}
}

func (s *ItemServiceImpl) GetItemById(id string) (dto.ItemDto, e.ApiError) {

	var itemDto dto.ItemDto
	itemDto, err := s.memcached.GetItemById(id)
	if err != nil {
		log.Debug("Error getting solr from memcached")
		itemDto, err2 := s.item.GetItemById(id)
		if err2 != nil {
			log.Debug("Error getting solr from mongo")
			return itemDto, err2
		}
		if itemDto.ItemId == "000000000000000000000000" {
			return itemDto, e.NewBadRequestApiError("solr not found")
		}
		_, err3 := s.memcached.InsertItem(itemDto)
		if err3 != nil {
			log.Debug("Error inserting in memcached")
		}
		log.Debug("mongo")
		return itemDto, nil
	} else {
		log.Debug("memcached")
		return itemDto, nil
	}
}

func (s *ItemServiceImpl) InsertItem(itemDto dto.ItemDto) (dto.ItemDto, e.ApiError) {

	var insertItem dto.ItemDto

	insertItem.Name = itemDto.Name

	insertItem, err := s.item.InsertItem(itemDto)
	if err != nil {
		return itemDto, e.NewBadRequestApiError("error inserting solr")
	}

	if insertItem.ItemId == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("error in insert")
	}
	itemDto.ItemId = insertItem.ItemId

	itemDto, err2 := s.memcached.InsertItem(insertItem)
	if err2 != nil {
		return itemDto, e.NewBadRequestApiError("Error inserting in memcached")
	}

	err = s.solr.Update(itemDto, "add")
	if err != nil {
		return itemDto, e.NewBadRequestApiError("Error inserting in solar")
	}
	return itemDto, nil
}
