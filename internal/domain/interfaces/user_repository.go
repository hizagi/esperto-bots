package interfaces

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type IUserRepository interface {
	GetUser(id string) (*entities.User, error)
}
