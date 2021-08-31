package repositories

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/utils"
)

type CategoryRepository struct {
	databaseClient *postgres.PostgresClient
}

func NewCategoryRepository(databaseClient *postgres.PostgresClient) *CategoryRepository {
	return &CategoryRepository{
		databaseClient,
	}
}

func (repository *CategoryRepository) GetCategory(id string) (*entities.Category, error) {
	var category entities.Category

	sql, args, err := repository.databaseClient.GetBuilder().Select("*").From("categories").Where(squirrel.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, err
	}

	err = sqlscan.Get(context.Background(), repository.databaseClient.Connection, &category, sql, args...)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repository *CategoryRepository) ListCategory(cursor string, pageSize int) ([]*entities.Category, *string, error) {
	var categories []*entities.Category
	var nextCursor string

	queryBuilder := repository.databaseClient.GetBuilder().Select("*").From("categories").OrderBy("created_at ASC, id ASC")

	if pageSize > 0 {
		queryBuilder = queryBuilder.Limit(uint64(pageSize))
	}

	if cursor != "" {
		createdCursor, id, err := utils.DecodeCursor(cursor)
		if err != nil {
			return nil, nil, err
		}
		queryBuilder = queryBuilder.Where("(created_at, id) > (?, ?)", createdCursor.Format("2006-01-02 15:04:05-07"), id)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, nil, err
	}

	err = sqlscan.Select(context.Background(), repository.databaseClient.Connection, &categories, query, args...)
	if err != nil {
		return nil, nil, err
	}

	categoriesLength := len(categories) - 1

	if categoriesLength > 0 {
		nextCursor = utils.EncodeCursor(categories[categoriesLength].CreatedAt, categories[categoriesLength].ID)
	}

	return categories, &nextCursor, nil
}
