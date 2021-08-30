package postgres

import (
	"fmt"
	"os"
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/seeds"
	"github.com/hizagi/esperto-bots/projectpath"
	"github.com/kristijorgji/goseeder"
	"github.com/pressly/goose"
)

func SetupDatabase(t *testing.T, seedFiles []string) (*postgres.PostgresClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(projectpath.Root)

	host, port := EmbedPostgres(t, seedFiles, configuration.PostgresConfiguration)

	configuration.PostgresConfiguration.Host = host
	configuration.PostgresConfiguration.Port = port

	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	fmt.Printf("Migration Path: %+v", postgres.MigrationPath())

	if err := goose.Up(postgresClient.Connection, postgres.MigrationPath()); err != nil {
		panic(err)
	}

	seeds.Init()
	err := goseeder.Execute(postgresClient.Connection)
	if err != nil {
		fmt.Printf("Seeding test data failed\n Err: %+v", err)
		os.Exit(-2)
	}

	return postgresClient, configuration
}
