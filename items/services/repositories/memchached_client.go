package repositories

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	json "github.com/json-iterator/go"
	"items/dto"
	e "items/utils/errors"
)

type MemcachedClient struct {
	Client *memcache.Client
}

func NewMemcachedInterface(host string, port int) *MemcachedClient {
	client := memcache.New(fmt.Sprintf("%s:%d", host, port))
	fmt.Println("[Memcached] Initialized connection")
	return &MemcachedClient{
		Client: client,
	}
}

func (repo *MemcachedClient) GetItemById(id string) (dto.ItemDto, e.ApiError) {
	item, err := repo.Client.Get(id)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dto.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
		}
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	var itemDto dto.ItemDto
	if err := json.Unmarshal(item.Value, &itemDto); err != nil {
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	return itemDto, nil
}

func (repo *MemcachedClient) InsertItem(item dto.ItemDto) (dto.ItemDto, e.ApiError) {
	bytes, err := json.Marshal(item)
	if err != nil {
		return dto.ItemDto{}, e.NewBadRequestApiError(err.Error())
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:        item.ItemId,
		Value:      bytes,
		Expiration: 5000,
	}); err != nil {
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.ItemId), err)
	}

	return item, nil
}

func (repo *MemcachedClient) Update(item dto.ItemDto) (dto.ItemDto, e.ApiError) {
	bytes, err := json.Marshal(item)
	if err != nil {
		return dto.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("invalid item %s: %v", item.ItemId, err))
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:   item.ItemId,
		Value: bytes,
	}); err != nil {
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item %s", item.ItemId), err)
	}

	return item, nil
}

func (repo *MemcachedClient) Delete(id string) e.ApiError {
	err := repo.Client.Delete(id)
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}
