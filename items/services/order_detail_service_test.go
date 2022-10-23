package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
)

type OrderDetailClientInterface struct {
	mock.Mock
}

func (m *OrderDetailClientInterface) GetOrderDetailById(id int) model.OrderDetail {
	ret := m.Called(id)
	return ret.Get(0).(model.OrderDetail)
}

func (m *OrderDetailClientInterface) GetOrderDetailsByOrderId(id int) model.OrderDetails {
	ret := m.Called(id)
	return ret.Get(0).(model.OrderDetails)
}

func (m *OrderDetailClientInterface) InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	ret := m.Called(orderDetail)
	return ret.Get(0).(model.OrderDetail)
}

func TestInsertDetail(t *testing.T) {
	var goodOrderDetail model.OrderDetail
	goodOrderDetail.OrderDetailId = 1
	goodOrderDetail.OrderId = 1
	goodOrderDetail.ProductId = 1
	goodOrderDetail.Quantity = 10
	goodOrderDetail.Price = 10
	goodOrderDetail.CurrencyId = "ARS"
	goodOrderDetail.Name = "Test_Product"

	var orderDetailInsertDto dto.OrderDetailInsertDto
	orderDetailInsertDto.ProductId = 1
	orderDetailInsertDto.Quantity = 10
	orderDetailInsertDto.Price = 10
	orderDetailInsertDto.CurrencyId = "ARS"
	orderDetailInsertDto.Name = "Test_Product"

	mockOrderDetailClient := new(OrderDetailClientInterface)
	mockOrderDetailClient.On("InsertOrderDetail", mock.AnythingOfType("model.OrderDetail")).Return(goodOrderDetail)
	service := initOrderDetailService(mockOrderDetailClient)
	res, err := service.InsertDetail(orderDetailInsertDto, 1)
	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res.OrderDetailId, 1)
}

func TestGetOrderDetailById(t *testing.T) {
	var orderDetail model.OrderDetail
	orderDetail.OrderDetailId = 1
	orderDetail.OrderId = 1
	orderDetail.ProductId = 1
	orderDetail.Quantity = 10
	orderDetail.Price = 10
	orderDetail.CurrencyId = "ARS"
	orderDetail.Name = "Test_Product"

	var orderDetailDto dto.OrderDetailDto
	orderDetailDto.OrderDetailId = 1
	orderDetailDto.ProductId = 1
	orderDetailDto.Quantity = 10
	orderDetailDto.Price = 10
	orderDetailDto.CurrencyId = "ARS"
	orderDetailDto.Name = "Test_Product"

	mockOrderDetailClient := new(OrderDetailClientInterface)
	mockOrderDetailClient.On("GetOrderDetailById", 1).Return(orderDetail)
	service := initOrderDetailService(mockOrderDetailClient)
	res, err := service.GetOrderDetailById(1)
	assert.Nil(t, err, "Error should be Nil")
	assert.Equal(t, res, orderDetailDto)
}

func TestGetOrderDetailsByOrderId(t *testing.T) {
	var orderDetail model.OrderDetail
	orderDetail.OrderDetailId = 1
	orderDetail.OrderId = 1
	orderDetail.ProductId = 1
	orderDetail.Quantity = 10
	orderDetail.Price = 10
	orderDetail.CurrencyId = "ARS"
	orderDetail.Name = "Test_Product"

	var orderDetails model.OrderDetails
	orderDetails = append(orderDetails, orderDetail)

	var empty model.OrderDetails

	mockOrderDetailClient := new(OrderDetailClientInterface)
	mockOrderDetailClient.On("GetOrderDetailsByOrderId", 1).Return(orderDetails)
	mockOrderDetailClient.On("GetOrderDetailsByOrderId", 0).Return(empty)
	service := initOrderDetailService(mockOrderDetailClient)
	res, err := service.GetOrderDetailsByOrderId(1)
	res2, err2 := service.GetOrderDetailsByOrderId(0)
	assert.Nil(t, err, "Error should be Nil")
	assert.Nil(t, err2, "Error should be Nil")

	assert.NotEqual(t, 0, len(res))
	assert.Equal(t, 0, len(res2))
}
