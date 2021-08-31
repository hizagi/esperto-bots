package repositories_test

import (
	"testing"

	"github.com/hizagi/esperto-bots/internal/infrastructure/database/mongodb/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/mongodb"
	"gotest.tools/v3/assert"
)

func TestGetCategoryMongo(t *testing.T) {
	mongoClient, _ := mongodb.SetupDatabase(t, []string{"category.js"})

	categoryRepository := repositories.NewCategoryRepository(mongoClient)

	category, err := categoryRepository.GetCategory("6122557b844c5e9e368e7dd6")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Category1", category.Name)
}

func TestListCategoriesMongo(t *testing.T) {
	mongoClient, _ := mongodb.SetupDatabase(t, []string{"category.js"})

	categoryRepository := repositories.NewCategoryRepository(mongoClient)

	categories, cursor, err := categoryRepository.ListCategory("6122557b844c5e9e368e7dd6", 2)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(categories))
	assert.Equal(t, categories[0].ID, "612627aee7b1dd1d463bdd65")
	assert.Equal(t, "612627d0392c056cd918c0f1", *cursor)
}
