package services

import (
	client "messages/clients/message"
	"messages/dto"
	"messages/model"
	e "messages/utils/errors"
	"time"
)

type messageService struct {
	messageClient client.MessageClientInterface
}

type messageServiceInterface interface {
	GetMessageById(id int) (dto.MessageDto, e.ApiError)
	GetMessagesByUserId(id int) (dto.MessagesDto, e.ApiError)
	GetMessages() (dto.MessagesDto, e.ApiError)
	InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError)
}

var (
	MessageService messageServiceInterface
)

func initMessageService(messageClient client.MessageClientInterface) messageServiceInterface {
	service := new(messageService)
	service.messageClient = messageClient
	return service
}

func init() {
	MessageService = initMessageService(client.MessageClient)
}

func (s *messageService) GetMessageById(id int) (dto.MessageDto, e.ApiError) {

	var message = s.messageClient.GetMessageById(id)
	var messageDto dto.MessageDto

	if message.MessageId == 0 {
		return messageDto, e.NewBadRequestApiError("message not found")
	}
	messageDto.MessageId = message.MessageId
	messageDto.UserId = message.UserId
	messageDto.ItemId = message.ItemId
	messageDto.Body = message.Body
	messageDto.CreatedAt = message.CreatedAt
	return messageDto, nil
}

func (s *messageService) GetMessagesByUserId(id int) (dto.MessagesDto, e.ApiError) {

	var messagesDto dto.MessagesDto
	var messages, err = s.messageClient.GetMessagesByUserId(id)

	if err != nil {
		return messagesDto, e.NewBadRequestApiError(err.Error())
	}

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.CreatedAt = message.CreatedAt
		messageDto.UserId = message.UserId
		messageDto.ItemId = message.ItemId
		messageDto.Body = message.Body
		messageDto.MessageId = message.MessageId

		messagesDto = append(messagesDto, messageDto)
	}
	return messagesDto, nil
}

func (s *messageService) GetMessages() (dto.MessagesDto, e.ApiError) {

	var messages = s.messageClient.GetMessages()
	var messagesDto dto.MessagesDto

	for _, message := range messages {
		var messageDto dto.MessageDto
		messageDto.CreatedAt = message.CreatedAt
		messageDto.UserId = message.UserId
		messageDto.ItemId = message.ItemId
		messageDto.Body = message.Body
		messageDto.MessageId = message.MessageId

		messagesDto = append(messagesDto, messageDto)
	}

	return messagesDto, nil
}

func (s *messageService) InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError) {

	var message model.Message

	message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	message.UserId = messageDto.UserId
	message.ItemId = messageDto.ItemId
	message.Body = messageDto.Body
	message.MessageId = messageDto.MessageId

	message = s.messageClient.InsertMessage(message)

	messageDto.MessageId = message.MessageId

	return messageDto, nil
}
