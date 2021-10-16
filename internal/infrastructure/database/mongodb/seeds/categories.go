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

const SEED_CATEGORIES = "SeedCategories"
const FIXED_CATEGORY_ID = "6122557b844c5e9e368e7dd6"
const FIXED_CATEGORY_NAME = "CATEGORY 1"

func (seeder Seeder) SeedCategories() {
	categories := make([]interface{}, 0)

	for i := 0; i < 100; i++ {

		category := models.Category{
			ID:        primitive.NewObjectID(),
			Name:      faker.Sentence(),
			Slug:      faker.Sentence(),
			CreatedAt: time.Now(),
		}

		if i == 90 {
			categoryID, _ := primitive.ObjectIDFromHex(FIXED_CATEGORY_ID)

			category = models.Category{
				ID:        categoryID,
				Name:      FIXED_CATEGORY_NAME,
				Slug:      faker.Sentence(),
				CreatedAt: time.Now(),
			}
		}

		categories = append(categories, category)
	}

	result, err := seeder.mongoDatabase.Collection(repositories.CATEGORY_COLLECTION).InsertMany(context.Background(), categories)

	if err != nil {
		panic(err)
	}

	log.Printf("Result affected: %d", len(result.InsertedIDs))
}
