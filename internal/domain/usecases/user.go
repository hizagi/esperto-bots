package usecases

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type UserUsecase interface {
	GetUser(id string) (*entities.User, error)
}
