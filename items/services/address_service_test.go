package services

import (
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"mvc-go/dto"
	"mvc-go/model"
	"testing"
)

type AddressClientInterface struct {
	mock.Mock
}

func (m *AddressClientInterface) InsertAddress(address model.Address) model.Address {
	ret := m.Called(address)
	return ret.Get(0).(model.Address)
}

func (m *AddressClientInterface) GetAddressById(id int) model.Address {
	ret := m.Called(id)
	return ret.Get(0).(model.Address)
}

func (m *AddressClientInterface) GetAddressesByUserId(id int) model.Addresses {
	ret := m.Called(id)
	return ret.Get(0).(model.Addresses)
}

func TestGetAddressesByUserId(t *testing.T) {
	mockClient := new(AddressClientInterface)

	var addresses model.Addresses
	var address model.Address
	addresses = append(addresses, address)

	var empty model.Addresses

	mockClient.On("GetAddressesByUserId", 1).Return(addresses) // Not empty addresses
	mockClient.On("GetAddressesByUserId", 2).Return(empty)     // Empty addresses
	service := initAddressService(mockClient)
	res, err := service.GetAddressesByUserId(1)
	assert.Nil(t, err, "Error should be nil")
	assert.NotEqual(t, 0, len(res))

	res2, err2 := service.GetAddressesByUserId(2)
	assert.NotNil(t, err2, "Error should not be nil")
	assert.Equal(t, 0, len(res2))

}

func TestGetAddressById(t *testing.T) {
	mockClient := new(AddressClientInterface)
	var address model.Address
	address.ID = 1

	var addressBad model.Address
	addressBad.ID = 0

	mockClient.On("GetAddressById", 1).Return(address)
	mockClient.On("GetAddressById", -1).Return(addressBad)
	service := initAddressService(mockClient)

	_, err := service.GetAddressById(1)
	assert.Nil(t, err, "Error should be nil")

	_, err2 := service.GetAddressById(-1)
	assert.NotNil(t, err2, "Error should not be nil")

}

func TestInsertAddress(t *testing.T) {

	mockClient := new(AddressClientInterface)
	var address dto.AddressDto
	address.UserId = 2
	address.Street1 = "Street1"
	address.Street2 = "Street2"
	address.Number = 123
	address.District = "District"
	address.City = "City"
	address.Country = "Country"

	var modeladdress model.Address
	modeladdress.UserId = 2
	modeladdress.Street1 = "Street1"
	modeladdress.Street2 = "Street2"
	modeladdress.Number = 123
	modeladdress.District = "District"
	modeladdress.City = "City"
	modeladdress.Country = "Country"

	var goodAddress model.Address
	goodAddress.UserId = 2
	goodAddress.Street1 = "Street1"
	goodAddress.Street2 = "Street2"
	goodAddress.Number = 123
	goodAddress.District = "District"
	goodAddress.City = "City"
	goodAddress.Country = "Country"
	goodAddress.ID = 1

	mockClient.On("InsertAddress", modeladdress).Return(goodAddress)
	service := initAddressService(mockClient)

	res, err := service.InsertAddress(address)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, res.AddressId, 1)
}
