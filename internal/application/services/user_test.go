package services_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/application/services"
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/mongodb"
)

func TestGetUserMongo(t *testing.T) {
	var users []entities.User

	test.LoadMockData("users.json", &users)

	mongoClient, _ := mongodb.SetupDatabase(t, "users", users)
	userRespository := repositories.NewUserRepository(mongoClient)
	userService := services.NewUserService(userRespository)
}
