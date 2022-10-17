package services

import (
	aclient "mvc-go/clients/address"
	oclient "mvc-go/clients/order"
	odclient "mvc-go/clients/order_detail"
	pclient "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type orderService struct {
	orderClient       oclient.OrderClientInterface
	productClient     pclient.ProductClientInterface
	addressClient     aclient.AddressClientInterface
	orderDetailClient odclient.OrderDetailClientInterface
}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError)
	GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func initOrderService(orderClient oclient.OrderClientInterface, orderDetailClient odclient.OrderDetailClientInterface, productClient pclient.ProductClientInterface, addressClient aclient.AddressClientInterface) orderServiceInterface {
	service := new(orderService)
	service.orderClient = orderClient
	service.productClient = productClient
	service.addressClient = addressClient
	service.orderDetailClient = orderDetailClient
	return service
}

func init() {
	OrderService = initOrderService(
		oclient.OrderClient,
		odclient.OrderDetailClient,
		pclient.ProductClient,
		aclient.AddressClient)
}

func (s *orderService) GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = s.orderClient.GetOrdersByUserId(id)
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto
		var address = s.addressClient.GetAddressById(order.AddressId)
		var details = s.orderDetailClient.GetOrderDetailsByOrderId(order.ID)
		orderDto.OrderId = order.ID
		orderDto.Address.AddressId = address.ID
		orderDto.Address.UserId = address.UserId
		orderDto.Address.Street1 = address.Street1
		orderDto.Address.Street2 = address.Street2
		orderDto.Address.Number = address.Number
		orderDto.Address.District = address.District
		orderDto.Address.City = address.City
		orderDto.Address.Country = address.Country
		orderDto.Date = order.Date
		orderDto.Total = order.Total
		orderDto.CurrencyId = order.CurrencyId
		for _, orderDetail := range details {
			var d dto.OrderDetailDto
			d.OrderDetailId = orderDetail.OrderDetailId
			d.ProductId = orderDetail.ProductId
			d.Quantity = orderDetail.Quantity
			d.Price = orderDetail.Price
			d.CurrencyId = orderDetail.CurrencyId
			d.Name = orderDetail.Name
			orderDto.OrderDetails = append(orderDto.OrderDetails, d)
		}

		ordersDto = append(ordersDto, orderDto)
	}
	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderInsertDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError) {

	var order model.Order
	var addressDto dto.AddressDto
	var total float32
	var orderResponseDto dto.OrderResponseDto

	res := s.addressClient.GetAddressById(orderInsertDto.Address.AddressId)
	if res.ID == 0 {
		var address model.Address
		address.UserId = orderInsertDto.Address.UserId
		address.Street1 = orderInsertDto.Address.Street1
		address.Street2 = orderInsertDto.Address.Street2
		address.Number = orderInsertDto.Address.Number
		address.District = orderInsertDto.Address.District
		address.City = orderInsertDto.Address.City
		address.Country = orderInsertDto.Address.Country
		address = s.addressClient.InsertAddress(address)
		orderInsertDto.Address.AddressId = address.ID
	} else {
		addressDto = orderInsertDto.Address
	}

	order.AddressId = addressDto.AddressId
	total = 0
	order.UserId = orderInsertDto.UserId
	order.Date = time.Now().Format("2006.01.02 15:04:05")
	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		detail := orderInsertDto.OrderDetails[i]
		product := s.productClient.GetProductById(detail.ProductId)
		if product.Stock < detail.Quantity {
			orderResponseDto.OrderId = 0
			return orderResponseDto, e.NewConflictApiError("Not enough stock on product: " + product.Name)
		}

		total += (detail.Price * float32(detail.Quantity))

	}

	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		detail := orderInsertDto.OrderDetails[i]
		s.productClient.RemoveStock(detail.ProductId, detail.Quantity)
	}

	order.Total = total
	order.CurrencyId = "ARS"

	order = s.orderClient.InsertOrder(order)

	orderResponseDto.OrderId = order.ID

	log.Debug(order)
	return orderResponseDto, nil
}
