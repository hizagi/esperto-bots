package models

import (
	"time"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ParentID  *string            `bson:"parent_id"`
	Name      string             `bson:"name"`
	Slug      string             `bson:"slug"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at"`
}

func (category *Category) ToDomain() *entities.Category {
	return &entities.Category{
		ID:        category.ID.Hex(),
		ParentID:  category.ParentID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
		UpdatedAt: &category.UpdatedAt,
		DeletedAt: &category.DeletedAt,
	}
}
