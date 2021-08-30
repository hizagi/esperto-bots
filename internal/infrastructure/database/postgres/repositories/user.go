package repositories

import (
	"context"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
)

type UserRepository struct {
	databaseClient *postgres.PostgresClient
}

func NewUserRepository(databaseClient *postgres.PostgresClient) *UserRepository {
	return &UserRepository{
		databaseClient,
	}
}

func (repository *UserRepository) GetUser(id string) (*entities.User, error) {
	var user entities.User

	err := sqlscan.Get(context.Background(), repository.databaseClient.Connection, &user, "SELECT * from users where id = $1", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
