package models

import (
	"time"

	"github.com/hizagi/esperto-bots/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Name      string             `bson:"name"`
	Lastname  string             `bson:"lastname"`
	Password  string             `bson:"password"`
	Document  string             `bson:"document"`
	BirthDate string             `bson:"birth_date"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at"`
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
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: &user.DeletedAt,
	}
}
