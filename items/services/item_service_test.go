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
	//var item model.Item
	//item.ItemId = 1
	//item.Name = "Test_Product"
	//item.Description = "Test_Desc"
	//item.Price = 500
	//item.CurrencyId = "ARS"
	//item.Stock = 5
	//item.Picture = "test.png"
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
	//mockItemClient.On("GetItemById", 1).Return(item)
	//service := initItemService(mockItemClient)
	//res, err := service.GetItemById("1")
	//assert.Nil(t, err, "Error should be nil")
	//assert.Equal(t, res, itemDto)
}
