package dto

type OrderDto struct {
	OrderId      int             `json:"order_id"`
	Address      AddressDto      `json:"address"`
	Date         string          `json:"date"`
	Total        float32         `json:"total"`
	CurrencyId   string          `json:"currency_id"`
	OrderDetails OrderDetailsDto `json:"details"`
}

type OrdersDto []OrderDto
