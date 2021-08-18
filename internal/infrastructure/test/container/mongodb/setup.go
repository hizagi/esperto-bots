package mongodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
)

func SetupDatabase(t *testing.T, collection string, mockData []interface{}) (*mongodb.MongoDBClient, *viper.Configuration) {
	configuration := viper.NewConfiguration()

	fmt.Printf("Config: %#v", configuration)

	host, port := EmbedMongo(t, configuration.MongoConfiguration)

	configuration.MongoConfiguration.Host = host
	configuration.MongoConfiguration.Port = port

	mongoClient := mongodb.NewMongoDbClient(configuration.MongoConfiguration)

	mongoClient.GetCollection(collection).InsertMany(context.Background(), mockData)

	return mongoClient, configuration
}
