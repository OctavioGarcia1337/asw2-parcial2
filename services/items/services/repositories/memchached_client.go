package repositories

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	json "github.com/json-iterator/go"
	"items/dto"
	e "items/utils/errors"
	"strconv"
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

func (repo *MemcachedClient) UpdateItem(item dto.ItemDto) (dto.ItemDto, e.ApiError) {
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

func (repo *MemcachedClient) DeleteItem(id string) e.ApiError {
	err := repo.Client.Delete(id)
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}

func (repo *MemcachedClient) InsertUser(user dto.UserDto) (dto.UserDto, e.ApiError) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return dto.UserDto{}, e.NewBadRequestApiError(err.Error())
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:        strconv.Itoa(user.UserId),
		Value:      bytes,
		Expiration: 5000,
	}); err != nil {
		return dto.UserDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %d", user.UserId), err)
	}

	return user, nil
}

func (repo *MemcachedClient) UpdateUser(user dto.UserDto) (dto.UserDto, e.ApiError) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return dto.UserDto{}, e.NewBadRequestApiError(fmt.Sprintf("invalid item %s: %v", user.UserId, err))
	}

	if err := repo.Client.Set(&memcache.Item{
		Key:   strconv.Itoa(user.UserId),
		Value: bytes,
	}); err != nil {
		return dto.UserDto{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item %s", user.UserId), err)
	}

	return user, nil
}

func (repo *MemcachedClient) GetUserById(id int) (dto.UserDto, e.ApiError) {
	user, err := repo.Client.Get(strconv.Itoa(id))
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dto.UserDto{}, e.NewNotFoundApiError(fmt.Sprintf("user %d not found", id))
		}
		return dto.UserDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting user %d", id), err)
	}

	var userDto dto.UserDto
	if err := json.Unmarshal(user.Value, &userDto); err != nil {
		return dto.UserDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting user %d", id), err)
	}

	return userDto, nil
}
