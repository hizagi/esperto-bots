package repositories

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type IUserRepository interface {
	CreateUser(user entities.User, roles []*entities.Role) (*entities.User, error)
	GetUser(id string) (*entities.User, error)
}
