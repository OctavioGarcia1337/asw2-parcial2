package dto

type ProductDto struct {
	ProductId   int         `json:"product_id"`
	Category    CategoryDto `json:"category"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float32     `json:"base_price"`
	CurrencyId  string      `json:"currency_id"`
	Stock       int         `json:"stock"`
	Picture     string      `json:"picture_url"`
}

type ProductsDto []ProductDto
