package seeds

import (
	"context"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/models"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SEED_USERS = "SeedUsers"

// {
//     "_id": "6122557b844c5e9e368e7dd6",
//     "email": "joao@joao.com",
//     "name": "joao",
//     "lastname": "vitor",
//     "password": "joao",
//     "document": "06262819541",
//     "birth_date": "04/02/1997",
//     "created_at": "",
//     "updated_at": "",
//     "deleted_at": ""
//   }

func (seeder Seeder) SeedUsers() {
	var users []interface{}

	for i := 0; i < 100; i++ {

		user := models.User{
			ID:        primitive.NewObjectID(),
			Name:      faker.FirstName(),
			Lastname:  faker.LastName(),
			Email:     faker.Email(),
			Password:  faker.Password(),
			Document:  faker.UUIDDigit(),
			BirthDate: faker.Date(),
			CreatedAt: time.Now(),
		}

		if i == 90 {
			id, _ := primitive.ObjectIDFromHex("6122557b844c5e9e368e7dd6")

			user = models.User{
				ID:        id,
				Name:      "Joao",
				Lastname:  faker.LastName(),
				Email:     faker.Email(),
				Password:  faker.Password(),
				Document:  faker.UUIDDigit(),
				BirthDate: faker.Date(),
				CreatedAt: time.Now(),
			}
		}

		users = append(users, user)
	}

	result, err := seeder.mongoDatabase.Collection(repositories.USER_COLLECTION).InsertMany(context.Background(), users)

	if err != nil {
		panic(err)
	}

	log.Printf("Result affected: %d", len(result.InsertedIDs))
}
