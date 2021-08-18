package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PostgresClient struct {
	Connection   *pgx.Conn
	databaseName string
	context      context.Context
}

func NewPostgresClient(databaseName string) *PostgresClient {
	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/test")

	return &PostgresClient{
		conn,
		databaseName,
		ctx,
	}
}

func (client *PostgresClient) Close() error {
	return client.Connection.Close(client.context)
}
