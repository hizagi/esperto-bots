package services

import (
	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"github.com/hizagi/esperto-bots/internal/domain/repositories"
)

type CategoryService struct {
	categoryRepository repositories.ICategoryRepository
}

func NewCategoryService(categoryRepository repositories.ICategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository,
	}
}

func (service *CategoryService) GetCategory(id string) (*entities.Category, error) {
	return service.categoryRepository.GetCategory(id)
}

func (service *CategoryService) ListCategory(lastID string, pageSize int) ([]*entities.Category, *string, error) {
	return service.categoryRepository.ListCategory(lastID, pageSize)
}
