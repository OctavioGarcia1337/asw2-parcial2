package services

import (
	"messages/dto"
	e "messages/utils/errors"
)

type MessageService interface {
	GetMessageById(id int) (dto.MessageDto, e.ApiError)
	GetMessagesByUserId(id int) (dto.MessagesDto, e.ApiError)
	GetMessages() (dto.MessagesDto, e.ApiError)
	InsertMessage(messageDto dto.MessageDto) (dto.MessageDto, e.ApiError)
}
