package repositories_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/postgres"
	"gotest.tools/v3/assert"
)

const toRoot = "../../../../../"

func TestGetUserPostgres(t *testing.T) {
	postgresClient, _ := postgres.SetupDatabase(t)

	userRepository := repositories.NewUserRepository(postgresClient)

	user, err := userRepository.GetUser("c2d29867-3d0b-d497-9191-18a9d8ee7830")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "joao", user.Name)
}
