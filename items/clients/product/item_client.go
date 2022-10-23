package product

import (
	"items/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

type itemClient struct{}

type ItemClientInterface interface {
	GetItemById(id int) model.Item
	InsertItem(item model.Item) model.Item
}

var (
	ItemClient ItemClientInterface
)

func init() {
	ItemClient = &itemClient{}
}

func (s *itemClient) GetProductById(id int) model.Item {
	var item model.Item
	Db.Where("item_id = ?", id).First(&item)
	log.Debug("Product: ", item)

	return item
}

func (s *itemClient) InsertProduct(item model.Item) model.Item {
	Db.Create(&item)
	log.Debug("Product: ", item)

	return item
}
