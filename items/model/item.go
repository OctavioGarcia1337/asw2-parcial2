package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ItemId      primitive.ObjectID `bson:"_id"`
	Titulo      string             `bson:"Titulo"`
	Tipo        string             `bson:"Tipo"`
	Ubicacion   string             `bson:"Ubicacion"`
	Precio_base int                `bson:"Precio_base"`
	Vendedor    string             `bson:"Vendedor"`
	Barrio      string             `bson:"Barrio"`
	Descripcion string             `bson:"Descripcion"`
	Dormitorios int                `bson:"Dormitorios"`
	Banos       int                `bson:"Banos"`
	Mts2        int                `bson:"Mts2"`
	Ambientes   int                `bson:"Ambientes"`
	Url_Img     string             `bson:"Url_Img"`
	Expensas    int                `bson:"Expensas"`
}

type Items []Item
