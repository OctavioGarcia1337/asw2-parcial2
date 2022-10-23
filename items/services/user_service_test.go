package services

import (
	"errors"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
)

type UserClientInterface struct {
	mock.Mock
}

func (m *UserClientInterface) GetUserById(id int) model.User {
	ret := m.Called(id)
	return ret.Get(0).(model.User)
}
func (m *UserClientInterface) GetUsers() model.Users {
	ret := m.Called()
	return ret.Get(0).(model.Users)
}
func (m *UserClientInterface) GetUserByUsername(username string) (model.User, error) {
	ret := m.Called(username)
	return ret.Get(0).(model.User), nil
}
func (m *UserClientInterface) InsertUser(user model.User) model.User {
	ret := m.Called(user)
	return ret.Get(0).(model.User)
}

func TestGetUserById(t *testing.T) {
	mockUserClient := new(UserClientInterface)
	var emptyAddresses model.Addresses
	var emptyAddressesDto dto.AddressesDto
	var user model.User
	user.UserId = 1
	user.Username = "test_username"
	user.Password = "test_password"
	user.FirstName = "test_firstname"
	user.LastName = "test_lastname"
	user.Email = "email@email"
	user.Addresses = emptyAddresses

	var emptyUser model.User
	emptyUser.UserId = 0

	var userDto dto.UserDto
	userDto.UserId = 1
	userDto.Username = "test_username"
	userDto.FirstName = "test_firstname"
	userDto.LastName = "test_lastname"
	userDto.Email = "email@email"
	userDto.Addresses = emptyAddressesDto

	var emptyDto dto.UserDto

	mockUserClient.On("GetUserById", 1).Return(user)
	mockUserClient.On("GetUserById", 0).Return(emptyUser)
	service := initUserService(mockUserClient)

	res, err := service.GetUserById(1)
	res2, err2 := service.GetUserById(0)

	assert.Nil(t, err, "Error should be Nil")
	assert.NotNil(t, err2, "Error should NOT be Nil")

	assert.Equal(t, res, userDto)   // Shouldn't return pass
	assert.Equal(t, res2, emptyDto) // Should be empty
}

func TestGetUsers(t *testing.T) {
	mockUserClient := new(UserClientInterface)
	var emptyAddresses model.Addresses
	var user model.User
	user.UserId = 1
	user.Username = "test_username"
	user.Password = "test_password"
	user.FirstName = "test_firstname"
	user.LastName = "test_lastname"
	user.Email = "email@email"
	user.Addresses = emptyAddresses

	var users model.Users
	users = append(users, user)

	mockUserClient.On("GetUsers").Return(users)
	service := initUserService(mockUserClient)

	res, err := service.GetUsers()

	assert.Nil(t, err, "Error should be Nil")
	assert.NotEqual(t, 0, len(res)) // Should be empty
}

func TestInsertUser(t *testing.T) {

	assert.Equal(t, 0, 0) // This is just an empty function for now
}

func TestLogin(t *testing.T) {
	mockUserClient := new(UserClientInterface)

	var emptyUser model.User

	var user model.User
	user.UserId = 1
	user.Username = "test"
	user.Password = "test"
	user.FirstName = "test_firstname"
	user.LastName = "test_lastname"
	user.Email = "email@email"

	var encryption model.User
	encryption.UserId = 2
	encryption.Username = "encrypted"
	encryption.Password = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoidGVzdCIsInVzZXJuYW1lIjoiZW5jcnlwdGVkIn0.0Bd47UDszBgDIY9jh1q07pattwOYF3zutP27oAoLlRk"
	encryption.FirstName = "test_encryption"
	encryption.LastName = "test_lastname"
	encryption.Email = "email@email"

	var correctUser dto.LoginDto
	correctUser.Username = "test"
	correctUser.Password = "test"

	var incorrectUser dto.LoginDto
	incorrectUser.Username = "testing"
	incorrectUser.Password = "test"

	var incorrectPass dto.LoginDto
	incorrectPass.Username = "test"
	incorrectPass.Password = "testing"

	var encryptionDto dto.LoginDto
	encryptionDto.Username = "encrypted"
	encryptionDto.Password = "test"

	var correctUserR dto.LoginResponseDto
	correctUserR.UserId = 1
	var incorrectUserR dto.LoginResponseDto
	incorrectUserR.UserId = -1
	var incorrectPassR dto.LoginResponseDto
	incorrectPassR.UserId = -1
	var encryptionDtoR dto.LoginResponseDto
	encryptionDtoR.UserId = 2
	encryptionDtoR.Token = encryption.Password

	mockUserClient.On("GetUserByUsername", "test").Return(user)
	mockUserClient.On("GetUserByUsername", "encrypted").Return(encryption)
	mockUserClient.On("GetUserByUsername", "testing").Return(emptyUser, errors.New("error"))
	service := initUserService(mockUserClient)

	res, err := service.Login(correctUser)

	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res.UserId, correctUserR.UserId)

	res, err = service.Login(incorrectUser)

	assert.NotNil(t, err, "Error should NOT be Nil")
	assert.Equal(t, res.UserId, incorrectUserR.UserId)

	res, err = service.Login(incorrectPass)

	assert.NotNil(t, err, "Error should NOT be Nil")
	assert.Equal(t, res.UserId, incorrectPassR.UserId)

	res, err = service.Login(encryptionDto)

	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res, encryptionDtoR)

}
