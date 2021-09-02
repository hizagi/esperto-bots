package repositories

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type IPermissionRepository interface {
	CreatePermission(resourceName string, create bool, update bool, delete bool) (*entities.Permission, error)
	GetPermission(id string) (*entities.Permission, error)
}
