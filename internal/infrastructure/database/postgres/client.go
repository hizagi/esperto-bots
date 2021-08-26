package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/hizagi/esperto-bots/internal/infrastructure/config/viper"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresClient struct {
	Connection *sql.DB
}

// func NewPostgresClient(postgresConfiguration viper.PostgresConfiguration) *PostgresClient {
// 	ctx := context.Background()

// 	config, err := pgxpool.ParseConfig(postgresConfiguration.GetDatabaseURL())

// 	config.MaxConns = postgresConfiguration.MaxConnections
// 	config.MaxConnIdleTime = time.Duration(postgresConfiguration.MaxConnectionIdleTime) * time.Millisecond
// 	config.MaxConnLifetime = time.Duration(postgresConfiguration.MaxConnectionLifetime) * time.Millisecond

// 	conn, err := pgxpool.ConnectConfig(ctx, config)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return &PostgresClient{
// 		conn,
// 		ctx,
// 	}
// }

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
