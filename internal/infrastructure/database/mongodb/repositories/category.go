package repositories

import (
	"context"
	"fmt"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepository struct {
	mongoCollection *mongo.Collection
}

func NewCategoryRepository(databaseClient *mongodb.MongoDBClient) *CategoryRepository {
	collection := databaseClient.GetCollection("categories")

	return &CategoryRepository{
		collection,
	}
}

func (repository *CategoryRepository) GetCategory(id string) (*entities.Category, error) {
	var category models.Category

	filter := bson.D{{Key: "_id", Value: id}}
	err := repository.mongoCollection.FindOne(context.Background(), filter).Decode(&category)

	if err != nil {
		return nil, err
	}

	return category.ToDomain(), nil
}

func (repository *CategoryRepository) ListCategory(lastID string, pageSize int) ([]*entities.Category, error) {
	ctx := context.Background()

	filter := bson.M{
		"_id": bson.M{
			"$gt": lastID,
		},
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	findOptions.SetLimit(int64(pageSize))

	cursor, err := repository.mongoCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	var categories []*entities.Category

	for cursor.Next(ctx) {
		var category models.Category

		err := cursor.Decode(&category)
		if err != nil {
			return nil, err
		}

		fmt.Printf("ToDomain: %+v", category.ToDomain())

		categories = append(categories, category.ToDomain())
	}

	return categories, nil
}
