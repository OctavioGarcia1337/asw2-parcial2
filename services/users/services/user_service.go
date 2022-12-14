package services

import (
	"users/dto"
	e "users/utils/errors"
)

type MessageService interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	DeleteUser(id int) e.ApiError
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError)
}
