package mongodb

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/seeds"
	"github.com/hizagi/esperto-bots/projectpath"
)

func SetupDatabase(t *testing.T, seedMethodNames ...string) (*mongodb.MongoDBClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(projectpath.Root)

	host, port := EmbedMongo(t, configuration.MongoConfiguration)

	configuration.MongoConfiguration.Host = host
	configuration.MongoConfiguration.Port = port

	mongoClient := mongodb.NewMongoDbClient(configuration.MongoConfiguration)

	seeds.RunSeeds(mongoClient.GetDatabase(), seedMethodNames...)

	return mongoClient, configuration
}
