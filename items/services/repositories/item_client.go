package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"items/dto"
	"items/model"
	e "items/utils/errors"
)

type ItemClient struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewItemInterface(host string, port int, collection string) *ItemClient {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://root:root@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256", host, port)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", names))

	return &ItemClient{
		Client:     client,
		Database:   client.Database("publicaciones"),
		Collection: collection,
	}
}

func (s *ItemClient) GetItemById(id string) (dto.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := s.Database.Collection(s.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dto.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dto.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dto.ItemDto{
		ItemId:      id,
		Titulo:      item.Titulo,
		Tipo:        item.Tipo,
		Ubicacion:   item.Ubicacion,
		PrecioBase:  item.PrecioBase,
		Vendedor:    item.Vendedor,
		Barrio:      item.Barrio,
		Descripcion: item.Descripcion,
		Dormitorios: item.Dormitorios,
		Banos:       item.Banos,
		Mts2:        item.Mts2,
		Ambientes:   item.Ambientes,
		UrlImg:      item.UrlImg,
		Expensas:    item.Expensas,
	}, nil

}

func (s *ItemClient) InsertItem(item dto.ItemDto) (dto.ItemDto, e.ApiError) {
	result, err := s.Database.Collection(s.Collection).InsertOne(context.TODO(), model.Item{
		ItemId:      primitive.NewObjectID(),
		Titulo:      item.Titulo,
		Tipo:        item.Tipo,
		Ubicacion:   item.Ubicacion,
		PrecioBase:  item.PrecioBase,
		Vendedor:    item.Vendedor,
		Barrio:      item.Barrio,
		Descripcion: item.Descripcion,
		Dormitorios: item.Dormitorios,
		Banos:       item.Banos,
		Mts2:        item.Mts2,
		Ambientes:   item.Ambientes,
		UrlImg:      item.UrlImg,
		Expensas:    item.Expensas,
	})

	if err != nil {
		return item, e.NewInternalServerApiError(fmt.Sprintf("error inserting to mongo %s", item.ItemId), err)
	}
	item.ItemId = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())

	return item, nil
}
