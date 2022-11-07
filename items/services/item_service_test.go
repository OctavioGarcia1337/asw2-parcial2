package services

import (
	mock "github.com/stretchr/testify/mock"
	"items/model"
	"testing"
)

type ItemClientInterface struct {
	mock.Mock
}

func (m *ItemClientInterface) GetItemById(id string) model.Item {
	ret := m.Called(id)
	return ret.Get(0).(model.Item)
}

func TestGetItemById(t *testing.T) {
	//mockItemClient := new(ItemClientInterface)
	//
	//var solr model.Item
	//solr.ItemId = 1
	//solr.Name = "Test_Product"
	//solr.Description = "Test_Desc"
	//solr.Price = 500
	//solr.CurrencyId = "ARS"
	//solr.Stock = 5
	//solr.Picture = "test.png"
	//
	//var itemDto dto.ItemDto
	//itemDto.ItemId = 1
	//itemDto.Name = "Test_Product"
	//itemDto.Description = "Test_Desc"
	//itemDto.Price = 500
	//itemDto.CurrencyId = "ARS"
	//itemDto.Stock = 5
	//itemDto.Picture = "test.png"
	//
	//mockItemClient.On("GetItemById", 1).Return(solr)
	//service := initItemService(mockItemClient)
	//res, err := service.GetItemById("1")
	//assert.Nil(t, err, "Error should be nil")
	//assert.Equal(t, res, itemDto)
}
