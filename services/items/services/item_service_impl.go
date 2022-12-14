package services

import (
	"fmt"
	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"items/config"
	"items/dto"
	client "items/services/repositories"
	e "items/utils/errors"
	"net/http"
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

func (s *ItemServiceImpl) GetUserById(id int, itemDto dto.ItemDto) (dto.ItemResponseDto, e.ApiError) {
	var userDto dto.UserDto
	var itemRDto dto.ItemResponseDto

	var er e.ApiError
	er = nil
	itemRDto.ItemId = itemDto.ItemId
	itemRDto.Titulo = itemDto.Titulo
	itemRDto.Tipo = itemDto.Tipo
	itemRDto.Ubicacion = itemDto.Ubicacion
	itemRDto.PrecioBase = itemDto.PrecioBase
	itemRDto.Vendedor = itemDto.Vendedor
	itemRDto.Barrio = itemDto.Barrio
	itemRDto.Descripcion = itemDto.Descripcion
	itemRDto.Dormitorios = itemDto.Dormitorios
	itemRDto.Banos = itemDto.Banos
	itemRDto.Mts2 = itemDto.Mts2
	itemRDto.Ambientes = itemDto.Ambientes
	itemRDto.UrlImg = itemDto.UrlImg
	itemRDto.Expensas = itemDto.Expensas
	itemRDto.UsuarioId = itemDto.UsuarioId

	userDto, err := s.memcached.GetUserById(id)
	if err != nil {
		resp, err := http.Get(fmt.Sprintf("http://%s:%d/%s/%d", config.USERSHOST, config.USERSPORT, config.USERSENDPOINT, id))
		if err != nil {
			return itemRDto, e.NewInternalServerApiError("Error getting user from user service", err)
		}
		err = json.NewDecoder(resp.Body).Decode(&userDto)
		if err != nil {
			return itemRDto, e.NewInternalServerApiError("Error decoding userDto", err)
		}

		userDto, err = s.memcached.InsertUser(userDto)
		if err != nil {
			er = e.NewInternalServerApiError("Error inserting user to memcached", err)
		}
	}

	itemRDto.Usuario = userDto.Username
	itemRDto.UNombre = userDto.FirstName
	itemRDto.UApellido = userDto.LastName
	itemRDto.UEmail = userDto.Email
	return itemRDto, er
}

func (s *ItemServiceImpl) GetItemById(id string) (dto.ItemResponseDto, e.ApiError) {

	var itemDto dto.ItemDto
	var itemResponseDto dto.ItemResponseDto

	itemDto, err := s.memcached.GetItemById(id)
	if err != nil {
		log.Debug("Error getting item from memcached")
		itemDto, err2 := s.item.GetItemById(id)
		if err2 != nil {
			log.Debug("Error getting item from mongo")
			return itemResponseDto, err2
		}
		if itemDto.ItemId == "000000000000000000000000" {
			return itemResponseDto, e.NewBadRequestApiError("item not found")
		}
		_, err3 := s.memcached.InsertItem(itemDto)
		if err3 != nil {
			log.Debug("Error inserting in memcached")
		}
		log.Debug("mongo")
		return s.GetUserById(itemDto.UsuarioId, itemDto)
	} else {
		log.Debug("memcached")
		return s.GetUserById(itemDto.UsuarioId, itemDto)
	}
}

func (s *ItemServiceImpl) GetItemsByUserId(id int) (dto.ItemsDto, e.ApiError) {

	var itemsDto dto.ItemsDto

	itemsDto, err := s.item.GetItemsByUserId(id)
	if err != nil {
		log.Debug("Error getting item from mongo")
		return itemsDto, err
	}

	return itemsDto, nil

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

	itemDto, err = s.memcached.InsertItem(itemDto)
	if err != nil {
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
			err = s.queue.SendMessage(item.ItemId, "create", item.ItemId)
			log.Debug(err)
		}()
	}
	return nil
}

func (s *ItemServiceImpl) DeleteUserItems(id int) e.ApiError {
	items, err := s.GetItemsByUserId(id)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := range items {
		var item dto.ItemDto
		item = items[i]
		go func() {
			err := s.item.DeleteItem(item.ItemId)
			if err != nil {
				log.Error(err)
			}
			err = s.queue.SendMessage(item.ItemId, "delete", fmt.Sprintf("%s.delete", item.ItemId))
			log.Error(err)
		}()
	}
	return nil
}
