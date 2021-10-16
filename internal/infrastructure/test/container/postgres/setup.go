package postgres

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/seeds"
	"github.com/hizagi/esperto-bots/projectpath"
	"github.com/pressly/goose"
)

func SetupDatabase(t *testing.T, seedMethodNames ...string) (*postgres.PostgresClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(projectpath.Root)

	host, port := EmbedPostgres(t, configuration.PostgresConfiguration)

	configuration.PostgresConfiguration.Host = host
	configuration.PostgresConfiguration.Port = port

	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	if err := goose.Up(postgresClient.Connection, postgres.MigrationPath()); err != nil {
		panic(err)
	}

	seeds.RunSeeds(postgresClient.Connection, seedMethodNames...)

	return postgresClient, configuration
}
