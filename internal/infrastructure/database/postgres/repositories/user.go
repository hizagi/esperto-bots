package repositories

import "github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"

type UserRepository struct {
	databaseClient *postgres.PostgresClient
}

func NewUserRepository(databaseClient *postgres.PostgresClient) *UserRepository {
	return &UserRepository{
		databaseClient,
	}
}
