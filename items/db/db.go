package db

import (
	addressClient "mvc-go/clients/address"
	categoryClient "mvc-go/clients/category"
	orderClient "mvc-go/clients/order"
	orderDetailClient "mvc-go/clients/order_detail"
	productClient "mvc-go/clients/product"
	userClient "mvc-go/clients/user"
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "cart"
	DBUser := "root"
	DBPass := ""
	DBHost := "dbmysql"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db
	orderClient.Db = db
	orderDetailClient.Db = db
	categoryClient.Db = db
	addressClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderDetail{}, &model.Category{}, &model.Address{})

	log.Info("Finishing Migration Database Tables")
}
