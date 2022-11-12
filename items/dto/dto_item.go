package dto

type ItemDto struct {
	ItemId      string  `json:"id"`
	Titulo      string  `json:"Titulo"`
	Tipo        string  `json:"Tipo"`
	Ubicacion   string  `json:"Ubicacion"`
	Precio_base int     `json:"Precio_base"`
	Vendedor    string  `json:"Vendedor"`
	Barrio      string  `json:"Barrio"`
	Descripcion string  `json:"Descripcion"`
	Dormitorios int     `json:"Dormitorios"`
	Baños       int     `json:"Baños"`
	Mts2        int     `json:"Mts2"`
	Ambientes   int     `json:"Ambientes"`
	Url_Img     string  `json:"Url_Img"`
	Expensas    int     `json:"Expensas"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	CurrencyId  string  `json:"currency_id"`
	Stock       int     `json:"stock"`
	Picture     string  `json:"picture_url"`
}

type ItemsDto []ItemDto
