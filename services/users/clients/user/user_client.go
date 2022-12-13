package user

import (
	"users/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type userClient struct{}

type UserClientInterface interface {
	GetUserById(id int) model.User
	GetUsers() model.Users
	GetUserByUsername(username string) (model.User, error)
	InsertUser(user model.User) model.User
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}
func (s *userClient) GetUserById(id int) model.User {
	var user model.User
	Db.Where("user_id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func (s *userClient) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (s *userClient) GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func (s *userClient) InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.UserId)
	return user
}
