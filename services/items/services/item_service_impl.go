package services

import (
	"bytes"
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

func (s *ItemServiceImpl) GetItemsByUserId(id int) (dto.ItemsResponseDto, e.ApiError) {

	var itemsDto dto.ItemsDto
	var itemsResponseDto dto.ItemsResponseDto
	itemsDto, err := s.item.GetItemsByUserId(id)
	if err != nil {
		log.Debug("Error getting items from mongo")
		return itemsResponseDto, err
	}

	for i := range itemsDto {
		item, err := s.GetUserById(itemsDto[i].UsuarioId, itemsDto[i])
		if err != nil {
			return itemsResponseDto, e.NewBadRequestApiError("error getting user for item")
		}
		itemsResponseDto = append(itemsResponseDto, item)
	}

	return itemsResponseDto, nil

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

func CheckQueue(processed chan string, total int, userid int) {
	var complete int
	var errors int
	for loop := true; loop; {
		select {
		case data := <-processed:
			if data == "error" {
				errors++
			} else {
				complete++
			}
			if errors+complete == total {
				loop = false
			}
		default:
			log.Debugf("waiting for %d more messages", total-complete-errors)
		}
	}
	var body []byte
	var message dto.MessageDto
	message.UserId = userid
	message.System = true
	message.Body = fmt.Sprintf("Processed items total = %d, Completed: %d, Errors: %d", complete+errors, complete, errors)
	body, err := json.Marshal(&message)

	if err != nil {
		panic(e.NewInternalServerApiError("Error marshaling in sending message", err))
	}
	_, err = http.Post(fmt.Sprintf("http://%s:%d/%s", config.MESSAGESHOST, config.MESSAGESPORT, config.MESSAGESENDPOINT), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(e.NewInternalServerApiError("Error sending message to message service", err))

	}
}

func (s *ItemServiceImpl) QueueItems(itemsDto dto.ItemsDto) e.ApiError {
	total := len(itemsDto)
	processed := make(chan string, total)
	for i := range itemsDto {
		var item dto.ItemDto
		item = itemsDto[i]
		go func() {
			item, err := s.item.InsertItem(item)
			if err != nil {
				processed <- "error"
				log.Debug(err)
			}
			processed <- "complete"
			err = s.queue.SendMessage(item.ItemId, "create", item.ItemId)
			log.Debug(err)
		}()
	}

	go CheckQueue(processed, total, itemsDto[0].UsuarioId)
	return nil
}

func (s *ItemServiceImpl) DeleteUserItems(id int) e.ApiError {
	items, err := s.GetItemsByUserId(id)
	if err != nil {
		log.Error(err)
		return err
	}
	for i := range items {
		var item dto.ItemResponseDto
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

func (s *ItemServiceImpl) DeleteItemById(id string) e.ApiError {

	err := s.item.DeleteItem(id)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.memcached.DeleteItem(id)
	if err != nil {
		log.Error("Error deleting from cache", err)
	}
	err = s.queue.SendMessage(id, "delete", fmt.Sprintf("%s.delete", id))
	log.Error(err)

	return nil
}
