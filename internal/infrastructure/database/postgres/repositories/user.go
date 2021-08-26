package repositories

import (
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
	err := repository.databaseClient.Connection.QueryRow("SELECT * from users where id = ?", id).Scan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
