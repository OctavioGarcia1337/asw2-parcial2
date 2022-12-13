package message

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"messages/model"
)

var Db *gorm.DB

type messageClient struct{}

type MessageClientInterface interface {
	GetMessageById(id int) model.Message
	GetMessagesByUserId(id int) (model.Messages, error)
	GetMessages() model.Messages
	InsertMessage(message model.Message) model.Message
}

var (
	MessageClient MessageClientInterface
)

func init() {
	MessageClient = &messageClient{}
}
func (s *messageClient) GetMessageById(id int) model.Message {
	var message model.Message
	Db.Where("message_id = ?", id).First(&message)
	log.Debug("Message: ", message)

	return message
}

func (s *messageClient) GetMessagesByUserId(id int) (model.Messages, error) {
	var messages model.Messages
	result := Db.Where("user_id = ?", id)
	if result.Error != nil {
		return messages, result.Error
	}

	return messages, nil
}

func (s *messageClient) GetMessages() model.Messages {
	var messages model.Messages
	Db.Find(&messages)

	log.Debug("Messages: ", messages)

	return messages
}

func (s *messageClient) InsertMessage(message model.Message) model.Message {
	result := Db.Create(&message)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Message Created: ", message.MessageId)
	return message
}
