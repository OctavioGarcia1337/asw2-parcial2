package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"messages/model"
)

type MessageClient struct {
	Db *gorm.DB
}

func NewMessageInterface(DBUser string, DBPass string, DBHost string, DBPort int, DBName string) *MessageClient {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName))
	if err != nil {
		panic(fmt.Sprintf("Error initializing SQL: %v", err))
	}
	return &MessageClient{
		Db: db,
	}
}

func (s *MessageClient) StartDbEngine() {
	// We need to migrate all classes model.
	s.Db.AutoMigrate(&model.Message{})

	log.Info("Finishing Migration Database Tables")
}

func (s *MessageClient) GetMessageById(id int) model.Message {
	var message model.Message
	s.Db.Where("id = ?", id).First(&message)
	log.Debug("Message: ", message)

	return message
}

func (s *MessageClient) GetMessagesByUserId(id int) (model.Messages, error) {
	var messages model.Messages
	result := s.Db.Find(&messages).Where("user_id = ?", id)
	if result.Error != nil {
		return messages, result.Error
	}

	return messages, nil
}

func (s *MessageClient) GetMessages() model.Messages {
	var messages model.Messages
	s.Db.Find(&messages)

	log.Debug("Messages: ", messages)

	return messages
}

func (s *MessageClient) InsertMessage(message model.Message) model.Message {
	result := s.Db.Create(&message)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Message Created: ", message.ID)
	return message
}
