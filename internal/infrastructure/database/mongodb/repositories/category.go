package repositories

import (
	"context"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CATEGORY_COLLECTION = "categories"

type CategoryRepository struct {
	mongoCollection *mongo.Collection
}

func NewCategoryRepository(databaseClient *mongodb.MongoDBClient) *CategoryRepository {
	collection := databaseClient.GetCollection(CATEGORY_COLLECTION)

	return &CategoryRepository{
		collection,
	}
}

func (repository *CategoryRepository) GetCategory(id string) (*entities.Category, error) {
	var category models.Category

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectID}}
	err := repository.mongoCollection.FindOne(context.Background(), filter).Decode(&category)

	if err != nil {
		return nil, err
	}

	return category.ToDomain(), nil
}

func (repository *CategoryRepository) ListCategory(lastID string, pageSize int) ([]*entities.Category, *string, error) {
	ctx := context.Background()

	filter := bson.M{}

	if lastID != "" {
		lastIdFormatted, _ := primitive.ObjectIDFromHex(lastID)
		filter = bson.M{
			"_id": bson.M{
				"$lt": lastIdFormatted,
			},
		}
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", 1}})
	findOptions.SetLimit(int64(pageSize))

	cursor, err := repository.mongoCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, nil, err
	}

	var categories []*entities.Category

	for cursor.Next(ctx) {
		var category models.Category

		err := cursor.Decode(&category)
		if err != nil {
			return nil, nil, err
		}

		categories = append(categories, category.ToDomain())
	}

	categoriesLength := len(categories) - 1

	if categoriesLength <= 0 {
		return categories, nil, nil
	}

	return categories, &categories[categoriesLength].ID, nil
}
