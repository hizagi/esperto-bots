package models

import (
	"time"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `json:"email"`
	Name      string             `json:"name"`
	Lastname  string             `json:"lastname"`
	Password  string             `json:"password"`
	Document  string             `json:"document"`
	BirthDate string             `json:"birth_date"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
}

func (user *User) ToDomain() *entities.User {
	return &entities.User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Password:  user.Password,
		Document:  user.Document,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
