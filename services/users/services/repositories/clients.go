package repositories

import (
	"users/model"
)

type Client interface {
	GetUserById(id int) model.User
	GetUsers() model.Users
	DeleteUser(id int) error
	GetUserByUsername(username string) (model.User, error)
	InsertUser(user model.User) model.User
}
