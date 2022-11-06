package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ItemId     primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Price      float32            `bson:"price"`
	CurrencyId string             `bson:"currency_id"`
	Stock      int                `bson:"stock"`
	Picture    string             `bson:"picture_url"`
}

type Items []Item
