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
}

func NewItemServiceImpl(
	item *client.ItemClient,
	memcached *client.MemcachedClient,
	queue *client.QueueClient,
) *ItemServiceImpl {
	return &ItemServiceImpl{
		item:      item,
		memcached: memcached,
		queue:     queue,
	}
}

func (s *ItemServiceImpl) GetItemById(id string) (dto.ItemDto, e.ApiError) {

	var itemDto dto.ItemDto
	itemDto, err := s.memcached.GetItemById(id)
	if err != nil {
		log.Debug("Error getting item from memcached")
		itemDto, err2 := s.item.GetItemById(id)
		if err2 != nil {
			log.Debug("Error getting item from mongo")
			return itemDto, err2
		}
		if itemDto.ItemId == "000000000000000000000000" {
			return itemDto, e.NewBadRequestApiError("item not found")
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

	insertItem, err := s.item.InsertItem(itemDto)
	if err != nil {
		return itemDto, e.NewBadRequestApiError("error inserting item")
	}

	if insertItem.ItemId == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("error in insert")
	}
	itemDto.ItemId = insertItem.ItemId

	itemDto, err2 := s.memcached.InsertItem(itemDto)
	if err2 != nil {
		return itemDto, e.NewBadRequestApiError("Error inserting in memcached")
	}
	return itemDto, nil
}

func (s *ItemServiceImpl) QueueItems(itemsDto dto.ItemsDto) e.ApiError {
	for i := range itemsDto {
		var item dto.ItemDto
		item = itemsDto[i]
		go func() {
			item, err := s.item.InsertItem(item)
			if err != nil {
				log.Debug(err)
			}
			err = s.queue.SendMessage("solr", item.ItemId)
			log.Debug(err)
		}()
	}
	return nil
}
