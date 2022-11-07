package dto

type ItemDto struct {
	ItemId     string    `json:"id"`
	Name       []string  `json:"name"`
	Price      []float32 `json:"price"`
	CurrencyId []string  `json:"currency_id"`
	Stock      []int     `json:"stock"`
	Picture    []string  `json:"picture_url"`
}

type ItemsDto []ItemDto
