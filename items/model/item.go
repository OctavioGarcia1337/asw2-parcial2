package model

type Item struct {
	ItemId      string `bson:"Item_id"`
	Titulo      string `bson:"Titulo"`
	Tipo        string `bson:"Tipo"`
	Ubicacion   string `bson:"Ubicacion"`
	Precio_base int    `bson:"Precio_base"`
	Vendedor    string `bson:"Vendedor"`
	Barrio      string `bson:"Barrio"`
	Descripcion string `bson:"Descripcion"`
	Dormitorios int    `bson:"Dormitorios"`
	Baños       int    `bson:"Baños"`
	Mts2        int    `bson:"Mts2"`
	Ambientes   int    `bson:"Ambientes"`
	Url_Img     string `bson:"Url_Img"`
	Expensas    int    `bson:"Expensas"`
}

type Items []Item
