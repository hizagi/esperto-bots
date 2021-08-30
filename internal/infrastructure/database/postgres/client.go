package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"
	"github.com/hizagi/esperto-bots/projectpath"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func MigrationPath() string {
	return projectpath.Root + "/internal/infrastructure/database/postgres/migrations"
}

type PostgresClient struct {
	Connection *sql.DB
}

func NewPostgresClient(postgresConfiguration viper.PostgresConfiguration) *PostgresClient {
	db, err := sql.Open("pgx", postgresConfiguration.GetDatabaseURL())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.SetMaxOpenConns(int(postgresConfiguration.MaxConnections))
	db.SetConnMaxLifetime(time.Duration(postgresConfiguration.MaxConnectionLifetime) * time.Millisecond)

	return &PostgresClient{
		db,
	}
}

func (client *PostgresClient) Close() error {
	return client.Connection.Close()
}
