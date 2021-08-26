package models

import (
	"time"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ParentID  string             `json:"parent_id"`
	Name      string             `json:"name"`
	Slug      string             `json:"slug"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
}

func (category *Category) ToDomain() *entities.Category {
	return &entities.Category{
		ID:        category.ID.Hex(),
		ParentID:  category.ParentID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		DeletedAt: category.DeletedAt,
	}
}
