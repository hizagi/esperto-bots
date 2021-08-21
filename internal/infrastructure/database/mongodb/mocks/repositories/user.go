package repositories

import (
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (repository *UserRepositoryMock) GetUser(id string) (*entities.User, error) {
	args := repository.Called(id)

	return args.Get(0).(*entities.User), args.Error(1)
}
