package postgres

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
)

func SetupDatabase(t *testing.T, toRootPath string, seedFiles []string) (*postgres.PostgresClient, *viper.Configuration) {
	configuration := viper.NewConfiguration(toRootPath)

	host, port := EmbedPostgres(t, toRootPath, seedFiles, configuration.PostgresConfiguration)

	configuration.PostgresConfiguration.Host = host
	configuration.PostgresConfiguration.Port = port

	postgresClient := postgres.NewPostgresClient(configuration.PostgresConfiguration)

	return postgresClient, configuration
}
