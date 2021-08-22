package mongodb

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
)

func SetupDatabase(t *testing.T, rootPath string) (*mongodb.MongoDBClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(rootPath)

	host, port := EmbedMongo(t, rootPath, configuration.MongoConfiguration)

	configuration.MongoConfiguration.Host = host
	configuration.MongoConfiguration.Port = port

	mongoClient := mongodb.NewMongoDbClient(configuration.MongoConfiguration)

	return mongoClient, configuration
}
