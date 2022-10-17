package dto

type OrderDetailInsertDto struct {
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float32 `json:"price"`
	CurrencyId string  `json:"currency_id"`
	Name       string  `json:"name"`
}

type OrderDetailsInsertDto []OrderDetailInsertDto
