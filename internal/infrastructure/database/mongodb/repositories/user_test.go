package repositories_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/application/services"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/mongodb"
	"gotest.tools/v3/assert"
)

func TestGetUserMongo(t *testing.T) {
	mongoClient, _ := mongodb.SetupDatabase(t, toRoot, []string{"user.js"})

	userRepository := repositories.NewUserRepository(mongoClient)

	userService := services.NewUserService(userRepository)

	user, err := userService.GetUser("6122557b844c5e9e368e7dd6")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "joao", user.Name)
}
