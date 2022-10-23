package services

import (
	"github.com/golang-jwt/jwt"
	client "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"
)

type userService struct {
	userClient client.UserClientInterface
}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func initUserService(userClient client.UserClientInterface) userServiceInterface {
	service := new(userService)
	service.userClient = userClient
	return service
}

func init() {
	UserService = initUserService(client.UserClient)
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = s.userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.UserId == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.FirstName = user.FirstName
	userDto.LastName = user.LastName
	userDto.Username = user.Username
	userDto.UserId = user.UserId
	userDto.Email = user.Email
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = s.userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.FirstName = user.FirstName
		userDto.LastName = user.LastName
		userDto.Username = user.Username
		userDto.Email = user.Email
		userDto.UserId = user.UserId

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.FirstName = userDto.FirstName
	user.LastName = userDto.LastName
	user.Username = userDto.Username
	user.Password = userDto.Password
	user.Email = userDto.Email

	user = s.userClient.InsertUser(user)

	userDto.UserId = user.UserId

	return userDto, nil
}

func (s *userService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {

	var user model.User
	user, err := s.userClient.GetUserByUsername(loginDto.Username)
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

	loginResponseDto.UserId = user.UserId
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}
