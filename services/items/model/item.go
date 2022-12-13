package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ItemId      primitive.ObjectID `bson:"_id"`
	Titulo      string             `bson:"titulo"`
	Tipo        string             `bson:"tipo"`
	Ubicacion   string             `bson:"ubicacion"`
	PrecioBase  int                `bson:"precio_base"`
	Vendedor    string             `bson:"vendedor"`
	Barrio      string             `bson:"barrio"`
	Descripcion string             `bson:"descripcion"`
	Dormitorios int                `bson:"dormitorios"`
	Banos       int                `bson:"banos"`
	Mts2        int                `bson:"mts2"`
	Ambientes   int                `bson:"ambientes"`
	UrlImg      string             `bson:"url_img"`
	Expensas    int                `bson:"expensas"`
}

type Items []Item
