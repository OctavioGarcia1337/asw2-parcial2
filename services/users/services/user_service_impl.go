package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"users/dto"
	"users/model"
	client "users/services/repositories"
	e "users/utils/errors"

	log "github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	userDB *client.UserClient
	queue  *client.QueueClient
}

func NewUserServiceImpl(
	userDB *client.UserClient,
	queue *client.QueueClient,
) *UserServiceImpl {
	userDB.StartDbEngine()
	return &UserServiceImpl{
		userDB: userDB,
		queue:  queue,
	}
}

func (s *UserServiceImpl) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user = s.userDB.GetUserById(id)
	var userDto dto.UserDto

	if user.ID == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Username = user.Username
	userDto.UserId = user.ID
	userDto.Email = user.Email
	return userDto, nil
}

func (s *UserServiceImpl) GetUsers() (dto.UsersDto, e.ApiError) {

	var users = s.userDB.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Username = user.Username
		userDto.Email = user.Email
		userDto.UserId = user.ID

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *UserServiceImpl) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.Password = userDto.Password
	user.Email = userDto.Email

	user = s.userDB.InsertUser(user)

	userDto.UserId = user.ID

	err := s.queue.SendMessage(userDto.UserId, "create", fmt.Sprintf("%d", userDto.UserId))
	if err != nil {
		return userDto, e.NewInternalServerApiError("error sending created message on user creation", err)
	}
	return userDto, nil
}

func (s *UserServiceImpl) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {

	var user model.User
	user, err := s.userDB.GetUserByUsername(loginDto.Username)
	var loginResponseDto dto.LoginResponseDto
	loginResponseDto.UserId = -1
	if err != nil {
		return loginResponseDto, e.NewBadRequestApiError("Usuario no encontrado")
	}
	if user.Password != loginDto.Password && loginDto.Username != "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginDto.Username == "encrypted" {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.UserId = user.ID
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

func (s *UserServiceImpl) DeleteUser(id int) e.ApiError {

	err := s.queue.SendMessage(id, "delete", fmt.Sprintf("%d", id))
	if err != nil {
		return e.NewInternalServerApiError("Error notifying other systems of user deletion. Canceling delete", err)
	}

	var user model.User
	user.ID = id
	er := s.userDB.DeleteUser(user)
	if er != nil {
		return e.NewInternalServerApiError("Error deleting user", er)
	}

	return nil
}
