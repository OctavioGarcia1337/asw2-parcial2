package dto

type AddressDto struct {
	AddressId int    `json:"address_id"`
	UserId    int    `json:"user_id"`
	Street1   string `json:"street1"`
	Street2   string `json:"street2"`
	Number    int    `json:"number"`
	District  string `json:"district"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

type AddressesDto []AddressDto
