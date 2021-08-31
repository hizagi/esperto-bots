package repositories

import (
	"context"

	"github.com/Masterminds/squirrel"
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

	sql, args, err := repository.databaseClient.GetBuilder().Select("*").From("users").Where(squirrel.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	err = sqlscan.Get(context.Background(), repository.databaseClient.Connection, &user, sql, args...)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
