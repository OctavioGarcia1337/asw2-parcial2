package db

import (
	messageClient "messages/clients/message"
	"messages/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Parameters
	DBName := "users_db"
	DBMessage := "root"
	DBPass := ""
	DBHost := "users_db"
	// ------------------------

	db, err = gorm.Open("mysql", DBMessage+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	messageClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.Message{})

	log.Info("Finishing Migration Database Tables")
}
