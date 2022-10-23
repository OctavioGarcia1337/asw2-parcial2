package dto

type OrderInsertDto struct {
	CurrencyId   string                `json:"currency_id"`
	UserId       int                   `json:"user_id"`
	Address      AddressDto            `json:"address"`
	OrderDetails OrderDetailsInsertDto `json:"details"`
}
