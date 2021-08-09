package repositories

import "github.com/hizagi/esperto-bots/internal/infrastructure/storage/postgres"

type UserRepository struct {
	databaseClient *postgres.PostgresClient
}

func NewUserRepository(databaseClient *postgres.PostgresClient) *UserRepository {
	return &UserRepository{
		databaseClient,
	}
}
