package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"messages/dto"
	"messages/model"
	"testing"
)

type MessageClientInterface struct {
	mock.Mock
}

func (m *MessageClientInterface) GetMessageById(id int) model.Message {
	ret := m.Called(id)
	return ret.Get(0).(model.Message)
}
func (m *MessageClientInterface) GetMessages() model.Messages {
	ret := m.Called()
	return ret.Get(0).(model.Messages)
}
func (m *MessageClientInterface) GetMessageByMessagename(messagename string) (model.Message, error) {
	ret := m.Called(messagename)
	return ret.Get(0).(model.Message), nil
}
func (m *MessageClientInterface) InsertMessage(message model.Message) model.Message {
	ret := m.Called(message)
	return ret.Get(0).(model.Message)
}

func TestGetMessageById(t *testing.T) {
	mockMessageClient := new(MessageClientInterface)
	var message model.Message
	message.MessageId = 1
	message.Messagename = "test_messagename"
	message.Password = "test_password"
	message.FirstName = "test_firstname"
	message.LastName = "test_lastname"
	message.Email = "email@email"

	var emptyMessage model.Message
	emptyMessage.MessageId = 0

	var messageDto dto.MessageDto
	messageDto.MessageId = 1
	messageDto.Messagename = "test_messagename"
	messageDto.FirstName = "test_firstname"
	messageDto.LastName = "test_lastname"
	messageDto.Email = "email@email"

	var emptyDto dto.MessageDto

	mockMessageClient.On("GetMessageById", 1).Return(message)
	mockMessageClient.On("GetMessageById", 0).Return(emptyMessage)
	service := initMessageService(mockMessageClient)

	res, err := service.GetMessageById(1)
	res2, err2 := service.GetMessageById(0)

	assert.Nil(t, err, "Error should be Nil")
	assert.NotNil(t, err2, "Error should NOT be Nil")

	assert.Equal(t, res, messageDto) // Shouldn't return pass
	assert.Equal(t, res2, emptyDto)  // Should be empty
}

func TestGetMessages(t *testing.T) {
	mockMessageClient := new(MessageClientInterface)
	var message model.Message
	message.MessageId = 1
	message.Messagename = "test_messagename"
	message.Password = "test_password"
	message.FirstName = "test_firstname"
	message.LastName = "test_lastname"
	message.Email = "email@email"

	var messages model.Messages
	messages = append(messages, message)

	mockMessageClient.On("GetMessages").Return(messages)
	service := initMessageService(mockMessageClient)

	res, err := service.GetMessages()

	assert.Nil(t, err, "Error should be Nil")
	assert.NotEqual(t, 0, len(res)) // Should be empty
}

func TestInsertMessage(t *testing.T) {

	assert.Equal(t, 0, 0) // This is just an empty function for now
}

func TestLogin(t *testing.T) {
	mockMessageClient := new(MessageClientInterface)

	var emptyMessage model.Message

	var message model.Message
	message.MessageId = 1
	message.Messagename = "test"
	message.Password = "test"
	message.FirstName = "test_firstname"
	message.LastName = "test_lastname"
	message.Email = "email@email"

	var encryption model.Message
	encryption.MessageId = 2
	encryption.Messagename = "encrypted"
	encryption.Password = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoidGVzdCIsInVzZXJuYW1lIjoiZW5jcnlwdGVkIn0.0Bd47UDszBgDIY9jh1q07pattwOYF3zutP27oAoLlRk"
	encryption.FirstName = "test_encryption"
	encryption.LastName = "test_lastname"
	encryption.Email = "email@email"

	var correctMessage dto.LoginDto
	correctMessage.Messagename = "test"
	correctMessage.Password = "test"

	var incorrectMessage dto.LoginDto
	incorrectMessage.Messagename = "testing"
	incorrectMessage.Password = "test"

	var incorrectPass dto.LoginDto
	incorrectPass.Messagename = "test"
	incorrectPass.Password = "testing"

	var encryptionDto dto.LoginDto
	encryptionDto.Messagename = "encrypted"
	encryptionDto.Password = "test"

	var correctMessageR dto.LoginResponseDto
	correctMessageR.MessageId = 1
	var incorrectMessageR dto.LoginResponseDto
	incorrectMessageR.MessageId = -1
	var incorrectPassR dto.LoginResponseDto
	incorrectPassR.MessageId = -1
	var encryptionDtoR dto.LoginResponseDto
	encryptionDtoR.MessageId = 2
	encryptionDtoR.Token = encryption.Password

	mockMessageClient.On("GetMessageByMessagename", "test").Return(message)
	mockMessageClient.On("GetMessageByMessagename", "encrypted").Return(encryption)
	mockMessageClient.On("GetMessageByMessagename", "testing").Return(emptyMessage, errors.New("error"))
	service := initMessageService(mockMessageClient)

	res, err := service.Login(correctMessage)

	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res.MessageId, correctMessageR.MessageId)

	res, err = service.Login(incorrectMessage)

	assert.NotNil(t, err, "Error should NOT be Nil")
	assert.Equal(t, res.MessageId, incorrectMessageR.MessageId)

	res, err = service.Login(incorrectPass)

	assert.NotNil(t, err, "Error should NOT be Nil")
	assert.Equal(t, res.MessageId, incorrectPassR.MessageId)

	res, err = service.Login(encryptionDto)

	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res, encryptionDtoR)

}
