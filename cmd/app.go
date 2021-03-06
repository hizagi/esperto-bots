package main

import (
	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/seeds"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/pressly/goose"
)

func main() {
	configuration := viper.NewConfiguration("./")
	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	defer postgresClient.Close()

	if err := goose.Up(postgresClient.Connection, postgres.MigrationPath()); err != nil {
		panic(err)
	}

	mongoClient := mongodb.NewMongoDbClient(configuration.MongoConfiguration)

	seeds.RunSeeds(mongoClient.GetDatabase())
}
