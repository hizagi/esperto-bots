package services

import (
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/domain/repositories"
)

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (service *UserService) GetUser(id string) (*entities.User, error) {
	return service.userRepository.GetUser(id)
}
