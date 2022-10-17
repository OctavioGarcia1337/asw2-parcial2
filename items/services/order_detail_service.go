package services

import (
	client "mvc-go/clients/order_detail"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"
)

type orderDetailService struct {
	orderDetailClient client.OrderDetailClientInterface
}

type orderDetailServiceInterface interface {
	InsertDetail(orderDetailDto dto.OrderDetailInsertDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError)
	GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetailsByOrderId(id int) (dto.OrderDetailsDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func initOrderDetailService(orderDetailClient client.OrderDetailClientInterface) orderDetailServiceInterface {
	service := new(orderDetailService)
	service.orderDetailClient = orderDetailClient
	return service
}

func init() {
	OrderDetailService = initOrderDetailService(client.OrderDetailClient)
}

func (s *orderDetailService) GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail = s.orderDetailClient.GetOrderDetailById(id)
	var orderDetailDto dto.OrderDetailDto

	orderDetailDto.OrderDetailId = orderDetail.OrderDetailId
	orderDetailDto.ProductId = orderDetail.ProductId
	orderDetailDto.Quantity = orderDetail.Quantity
	orderDetailDto.Price = orderDetail.Price
	orderDetailDto.CurrencyId = orderDetail.CurrencyId
	orderDetailDto.Name = orderDetail.Name

	return orderDetailDto, nil
}

func (s *orderDetailService) GetOrderDetailsByOrderId(id int) (dto.OrderDetailsDto, e.ApiError) {
	var orderDetails model.OrderDetails = s.orderDetailClient.GetOrderDetailsByOrderId(id)
	var orderDetailsDto dto.OrderDetailsDto

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.OrderDetailId = orderDetail.OrderDetailId
		orderDetailDto.ProductId = orderDetail.ProductId
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.CurrencyId = orderDetail.CurrencyId
		orderDetailDto.Name = orderDetail.Name

		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

func (s *orderDetailService) InsertDetail(orderDetailDto dto.OrderDetailInsertDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError) {

	var orderDetail model.OrderDetail
	orderDetail.OrderId = orderId
	orderDetail.ProductId = orderDetailDto.ProductId
	orderDetail.Quantity = orderDetailDto.Quantity
	orderDetail.CurrencyId = "ARS"
	orderDetail.Name = orderDetailDto.Name
	orderDetail.Price = orderDetailDto.Price

	orderDetail = s.orderDetailClient.InsertOrderDetail(orderDetail)

	var orderDetailResponseDto dto.OrderDetailResponseDto
	orderDetailResponseDto.OrderDetailId = orderDetail.OrderDetailId

	log.Debug(orderDetail)

	return orderDetailResponseDto, nil
}
