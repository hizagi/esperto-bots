package postgres

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/hizagi/esperto-bots/projectpath"
)

func SetupDatabase(t *testing.T, seedFiles []string) (*postgres.PostgresClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(projectpath.Root)

	host, port := EmbedPostgres(t, seedFiles, configuration.PostgresConfiguration)

	configuration.PostgresConfiguration.Host = host
	configuration.PostgresConfiguration.Port = port

	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	return postgresClient, configuration
}
