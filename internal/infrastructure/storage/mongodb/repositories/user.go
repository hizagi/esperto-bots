package repositories

import (
	"github.com/hizagi/esperto-bots/internal/infrastructure/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoCollection *mongo.Collection
}

func NewUserRepository(databaseClient *mongodb.MongoDBClient) *UserRepository {
	collection := databaseClient.GetCollection("users")

	return &UserRepository{
		collection,
	}
}
