package usecases

import "github.com/hizagi/esperto-bots/internal/domain/entities"

type CategoryUsecase interface {
	GetCategory(id string) (*entities.Category, error)
	ListCategory(lastID string, pageSize int) ([]*entities.Category, *string, error)
}
