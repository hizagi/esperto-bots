package repositories

import (
	"context"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USER_COLLECTION = "users"

type UserRepository struct {
	mongoCollection *mongo.Collection
}

func NewUserRepository(databaseClient *mongodb.MongoDBClient) *UserRepository {
	collection := databaseClient.GetCollection(USER_COLLECTION)

	return &UserRepository{
		collection,
	}
}

func (repository *UserRepository) GetUser(id string) (*entities.User, error) {
	var user models.User

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectID}}
	err := repository.mongoCollection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}
