package dto

type CartDto struct {
	CartId   int         `json:"cart_id"`
	UserId   int         `json:"user_id"`
	Active   bool        `json:"active"`
	Products ProductsDto `json:"products"`
}

type CartsDto []CartDto
