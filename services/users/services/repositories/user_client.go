package repositories

import (
	"fmt"
	"users/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type UserClient struct {
	Db *gorm.DB
}

func NewUserInterface(DBUser string, DBPass string, DBHost string, DBPort int, DBName string) *UserClient {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName))
	if err != nil {
		panic(fmt.Sprintf("Error initializing SQL: %v", err))
	}
	return &UserClient{
		Db: db,
	}
}
func (s *UserClient) StartDbEngine() {
	// We need to migrate all classes model.
	s.Db.AutoMigrate(&model.User{})

	log.Info("Finishing Migration Database Tables")
}

func (s *UserClient) GetUserById(id int) model.User {
	var user model.User
	s.Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func (s *UserClient) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := s.Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (s *UserClient) GetUsers() model.Users {
	var users model.Users
	s.Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func (s *UserClient) InsertUser(user model.User) model.User {
	result := s.Db.Create(&user)

	if result.Error != nil {
		log.Error(result.Error)
		user.ID = 0
		return user
	}
	log.Debug("User Created: ", user.ID)
	return user
}

func (s *UserClient) DeleteUser(user model.User) error {
	result := s.Db.Delete(&user)
	return result.Error
}
