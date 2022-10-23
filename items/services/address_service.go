package services

import (
	client "mvc-go/clients/address"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"
)

type addressService struct {
	addressClient client.AddressClientInterface
}

type addressServiceInterface interface {
	InsertAddress(addressDto dto.AddressDto) (dto.AddressDto, e.ApiError)
	GetAddressesByUserId(id int) (dto.AddressesDto, e.ApiError)
	GetAddressById(id int) (dto.AddressDto, e.ApiError)
}

var (
	AddressService addressServiceInterface
)

func initAddressService(addressClient client.AddressClientInterface) addressServiceInterface {
	service := new(addressService)
	service.addressClient = addressClient
	return service
}

func init() {
	AddressService = initAddressService(client.AddressClient)
}

func (s *addressService) GetAddressesByUserId(id int) (dto.AddressesDto, e.ApiError) {

	var addresses model.Addresses = s.addressClient.GetAddressesByUserId(id)
	var addressesDto dto.AddressesDto
	if addresses == nil {
		return addressesDto, e.NewBadRequestApiError("asashdfh")
	}
	if len(addresses) == 0 {
		return addressesDto, e.NewBadRequestApiError("no addresses found for user")
	}
	for _, address := range addresses {
		var addressDto dto.AddressDto
		addressDto.AddressId = address.ID
		addressDto.UserId = id
		addressDto.Street1 = address.Street1
		addressDto.Street2 = address.Street2
		addressDto.Number = address.Number
		addressDto.District = address.District
		addressDto.City = address.City
		addressDto.Country = address.Country
		addressesDto = append(addressesDto, addressDto)
	}
	return addressesDto, nil
}

func (s *addressService) InsertAddress(addressDto dto.AddressDto) (dto.AddressDto, e.ApiError) {

	var address model.Address

	address.UserId = addressDto.UserId
	address.Street1 = addressDto.Street1
	address.Street2 = addressDto.Street2
	address.Number = addressDto.Number
	address.District = addressDto.District
	address.City = addressDto.City
	address.Country = addressDto.Country
	address = s.addressClient.InsertAddress(address)

	addressDto.AddressId = address.ID

	log.Debug(address)
	return addressDto, nil
}

func (s *addressService) GetAddressById(id int) (dto.AddressDto, e.ApiError) {
	var address model.Address
	var addressDto dto.AddressDto
	address = s.addressClient.GetAddressById(id)
	if address.ID == 0 {
		return addressDto, e.NewBadRequestApiError("address not found")
	}
	addressDto.AddressId = id
	addressDto.UserId = address.UserId
	addressDto.Street1 = address.Street1
	addressDto.Street2 = address.Street2
	addressDto.Number = address.Number
	addressDto.District = address.District
	addressDto.City = address.City
	addressDto.Country = address.Country

	return addressDto, nil
}
