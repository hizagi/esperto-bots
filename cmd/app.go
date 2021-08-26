package main

import (
	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/hizagi/esperto-bots/internal/infrastructure/utils"
	"github.com/pressly/goose"
)

func main() {
	configuration := viper.NewConfiguration("./")
	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	defer postgresClient.Close()

	if err := goose.Up(postgresClient.Connection, utils.FromRootPath("./", "/migrations/postgres")); err != nil {
		panic(err)
	}

}
