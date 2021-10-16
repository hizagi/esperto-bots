package repositories_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/seeds"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/mongodb"
	"gotest.tools/v3/assert"
)

func TestGetCategoryMongo(t *testing.T) {
	mongoClient, _ := mongodb.SetupDatabase(t, seeds.SEED_CATEGORIES)

	categoryRepository := repositories.NewCategoryRepository(mongoClient)

	category, err := categoryRepository.GetCategory(seeds.FIXED_CATEGORY_ID)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, seeds.FIXED_CATEGORY_NAME, category.Name)
}

func TestListCategoriesMongo(t *testing.T) {
	mongoClient, _ := mongodb.SetupDatabase(t, seeds.SEED_CATEGORIES)
	categoryRepository := repositories.NewCategoryRepository(mongoClient)

	tests := []struct {
		name         string
		pageSize     int
		expectedSize int
		cursor       string
	}{
		{
			name:         "No cursor and with pagesize",
			pageSize:     30,
			expectedSize: 30,
			cursor:       "",
		},
		{
			name:         "With cursor and with pagesize",
			pageSize:     30,
			expectedSize: 0,
			cursor:       seeds.FIXED_CATEGORY_ID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			categories, _, err := categoryRepository.ListCategory(test.cursor, test.pageSize)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.expectedSize, len(categories))
		})
	}

}
