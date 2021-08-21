package services_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/application/services"
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/mocks/repositories"
	"github.com/stretchr/testify/mock"
	"gotest.tools/v3/assert"
)

func TestGetUserMongo(t *testing.T) {
	userRepository := repositories.NewUserRepositoryMock()

	userRepository.On("GetUser", mock.AnythingOfType("string")).Return(&entities.User{
		ID:        "1",
		Name:      "joao",
		Lastname:  "vitor",
		Email:     "joao@ferraz.com",
		Password:  "123",
		Document:  "33251235",
		BirthDate: "04/02/1997",
	}, nil)

	userService := services.NewUserService(userRepository)

	user, err := userService.GetUser("1")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "joao", user.Name)
}
