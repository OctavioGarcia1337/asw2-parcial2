package repositories

import (
	"items/dto"
	"items/utils/errors"
)

type Client interface {
	GetItemById(id string) (dto.ItemDto, errors.ApiError)
	InsertItem(book dto.ItemDto) (dto.ItemDto, errors.ApiError)
	Update(book dto.ItemDto) (dto.ItemDto, errors.ApiError)
	Delete(id string) errors.ApiError
}
