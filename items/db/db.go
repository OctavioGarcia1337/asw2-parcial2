package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database
var client *mongo.Client

func Disconect_db() {

	client.Disconnect(context.TODO())
}

func Init_db() error {

	clientOpts := options.Client().ApplyURI("mongodb://root:root@localhost:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli
	if err != nil {
		return err
	}

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	MongoDb = client.Database("publicaciones")

	fmt.Println("Available datatabases:")
	fmt.Println(dbNames)

	return nil
}
