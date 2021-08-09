package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	internalClient *mongo.Client
	databaseName   string
}

func NewMongoDbClient(databaseName string) *MongoDBClient {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDBClient{
		client,
		databaseName,
	}
}

func (client *MongoDBClient) GetCollection(collectionName string) *mongo.Collection {
	return client.internalClient.Database(client.databaseName).Collection(collectionName)
}
