package services

import (
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/domain/interfaces"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func NewUserService(userRepository interfaces.IUserRepository) *UserService {
	return &UserService{
		userRepository,
	}
}

func (service *UserService) GetUser(id string) entities.User {
	return service.userRepository.GetUser(id)
}
