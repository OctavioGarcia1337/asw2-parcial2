package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
)

type OrderClientInterface struct {
	mock.Mock
}

func (m *OrderClientInterface) InsertOrder(order model.Order) model.Order {
	ret := m.Called(order)
	return ret.Get(0).(model.Order)
}

func (m *OrderClientInterface) GetOrderById(id int) model.Order {
	ret := m.Called(id)
	return ret.Get(0).(model.Order)
}

func (m *OrderClientInterface) GetOrdersByUserId(id int) model.Orders {
	ret := m.Called(id)
	return ret.Get(0).(model.Orders)
}

func TestGetOrdersByUserId(t *testing.T) {
	mockOrderClient := new(OrderClientInterface)
	mockOrderDetailClient := new(OrderDetailClientInterface)
	mockProductClient := new(ProductClientInterface)
	mockAddressClient := new(AddressClientInterface)

	var address model.Address
	address.ID = 1
	address.UserId = 2
	address.Street1 = "Street1"
	address.Street2 = "Street2"
	address.Number = 123
	address.District = "District"
	address.City = "City"
	address.Country = "Country"

	var details model.OrderDetails
	var detail model.OrderDetail
	detail.OrderDetailId = 1
	detail.OrderId = 1
	detail.ProductId = 1
	detail.Quantity = 2
	detail.Price = 20
	detail.CurrencyId = "ARS"
	detail.Name = "Test_Product"
	details = append(details, detail)

	var order model.Order
	order.UserId = 2
	order.AddressId = 1
	order.Date = "2020/5/20"
	order.Total = 123
	order.CurrencyId = "ARS"
	order.ID = 1

	var empty model.Orders
	var orders model.Orders
	orders = append(orders, order)
	mockAddressClient.On("GetAddressById", 1).Return(address)
	mockOrderDetailClient.On("GetOrderDetailsByOrderId", 1).Return(details)
	mockOrderClient.On("GetOrdersByUserId", 1).Return(orders) // Not empty orders
	mockOrderClient.On("GetOrdersByUserId", 2).Return(empty)  // Empty orders

	service := initOrderService(
		mockOrderClient,
		mockOrderDetailClient,
		mockProductClient,
		mockAddressClient)
	res, err := service.GetOrdersByUserId(1)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, len(res))

	res2, err2 := service.GetOrdersByUserId(2)
	assert.Nil(t, err2, "Error should be nil")
	assert.Equal(t, 0, len(res2))
}

func TestInsertOrder(t *testing.T) {

	mockOrderClient := new(OrderClientInterface)
	mockOrderDetailClient := new(OrderDetailClientInterface)
	mockProductClient := new(ProductClientInterface)
	mockAddressClient := new(AddressClientInterface)

	var addressDto dto.AddressDto
	addressDto.AddressId = 1
	addressDto.UserId = 2
	addressDto.Street1 = "Street1"
	addressDto.Street2 = "Street2"
	addressDto.Number = 123
	addressDto.District = "District"
	addressDto.City = "City"
	addressDto.Country = "Country"

	var address model.Address
	address.ID = 1
	address.UserId = 2
	address.Street1 = "Street1"
	address.Street2 = "Street2"
	address.Number = 123
	address.District = "District"
	address.City = "City"
	address.Country = "Country"

	var product model.Product
	product.ProductId = 1
	product.CategoryId = 1
	product.Name = "Test_Product"
	product.Description = "Test_Desct"
	product.Price = 20
	product.CurrencyId = "ARS"
	product.Stock = 10
	product.Picture = "test.png"

	var lessStock model.Product
	lessStock.ProductId = 1
	lessStock.CategoryId = 1
	lessStock.Name = "Test_Product"
	lessStock.Description = "Test_Desct"
	lessStock.Price = 20
	lessStock.CurrencyId = "ARS"
	lessStock.Stock = 8
	lessStock.Picture = "test.png"

	var orderDetailsDto dto.OrderDetailsInsertDto
	var orderDetailDto dto.OrderDetailInsertDto
	orderDetailDto.ProductId = 1
	orderDetailDto.Quantity = 2
	orderDetailDto.Price = 20
	orderDetailDto.CurrencyId = "ARS"
	orderDetailDto.Name = "Test_Product"
	orderDetailsDto = append(orderDetailsDto, orderDetailDto)

	var order dto.OrderInsertDto
	order.UserId = 2
	order.CurrencyId = "ARS"
	order.Address = addressDto
	order.OrderDetails = orderDetailsDto

	var goodOrder model.Order
	goodOrder.ID = 1
	goodOrder.UserId = 2
	goodOrder.AddressId = 1
	goodOrder.Date = "2020/5/20"
	goodOrder.Total = 40
	goodOrder.CurrencyId = "ARS"

	mockOrderClient.On("InsertOrder", mock.AnythingOfType("model.Order")).Return(goodOrder)
	mockAddressClient.On("GetAddressById", 1).Return(address)
	mockProductClient.On("GetProductById", 1).Return(product)
	mockProductClient.On("RemoveStock", 1, 2).Return(lessStock)
	service := initOrderService(
		mockOrderClient,
		mockOrderDetailClient,
		mockProductClient,
		mockAddressClient)

	res, err := service.InsertOrder(order)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, res.OrderId, 1)
}
