package repositories

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type IRoleRepository interface {
	CreateRole(roleName string, permission []*entities.Permission) (*entities.Role, error)
	GetRole(id string) (*entities.Role, error)
}
