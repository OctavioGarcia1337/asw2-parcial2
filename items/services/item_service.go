package services

import (
	"items/dto"
	e "items/utils/errors"
)

type ItemService interface {
	GetItemById(id string) (dto.ItemDto, e.ApiError)
	InsertItem(item dto.ItemDto) (dto.ItemDto, e.ApiError)
	QueueItems(items dto.ItemsDto) e.ApiError
}
