package repositories

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
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
		UsuarioId:   item.UsuarioId,
	}, nil

}

func (s *ItemClient) GetItemsByUserId(id int) (dto.ItemsDto, e.ApiError) {

	result, _ := s.Database.Collection(s.Collection).Find(context.TODO(), bson.D{
		{"usuario_id", id},
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dto.ItemsDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %d not found", id))
	}
	var items model.Items

	if err := result.All(context.TODO(), &items); err != nil {
		return dto.ItemsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %d", id), err)
	}
	var itemsDto dto.ItemsDto
	for i := range items {
		item := items[i]
		itemsDto = append(itemsDto,
			dto.ItemDto{
				ItemId:      item.ItemId.Hex(),
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
				UsuarioId:   item.UsuarioId,
			})
	}
	return itemsDto, nil

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
		UsuarioId:   item.UsuarioId,
	})

	if err != nil {
		return item, e.NewInternalServerApiError(fmt.Sprintf("error inserting to mongo %s", item.ItemId), err)
	}
	item.ItemId = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())

	return item, nil
}

func (s *ItemClient) DeleteItem(id string) e.ApiError {
	result, err := s.Database.Collection(s.Collection).DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Error(err)
		return e.NewInternalServerApiError("error deleting item", err)
	}
	log.Debug(result.DeletedCount)
	return nil
}
