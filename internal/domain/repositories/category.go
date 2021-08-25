package repositories

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type ICategoryRepository interface {
	GetCategory(id string) (*entities.Category, error)
	ListCategory(lastID string, pageSize int) ([]*entities.Category, error)
}
