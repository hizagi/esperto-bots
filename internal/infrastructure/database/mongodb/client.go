package mongodb

import (
	"context"
	"log"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	internalClient *mongo.Client
	databaseName   string
}

func NewMongoDbClient(mongoConfiguration viper.MongoConfiguration) *MongoDBClient {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoConfiguration.GetDatabaseURL())

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
		mongoConfiguration.DatabaseName,
	}
}

func (client *MongoDBClient) GetCollection(collectionName string) *mongo.Collection {
	return client.internalClient.Database(client.databaseName).Collection(collectionName)
}
