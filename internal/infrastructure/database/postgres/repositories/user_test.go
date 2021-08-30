package repositories_test

import (
	"fmt"
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/postgres"
	"gotest.tools/v3/assert"
)

const toRoot = "../../../../../"

func TestGetUserPostgres(t *testing.T) {
	postgresClient, _ := postgres.SetupDatabase(t, []string{"user.sql"})

	result, err := postgresClient.Connection.Query(`SELECT *
	FROM pg_catalog.pg_tables
	WHERE schemaname != 'pg_catalog' AND 
		schemaname != 'information_schema';`)

	for result.Next() {
		fmt.Printf("DATA: %+v\n", result)
	}
	userRepository := repositories.NewUserRepository(postgresClient)

	user, err := userRepository.GetUser("c2d29867-3d0b-d497-9191-18a9d8ee7830")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "joao", user.Name)
}
