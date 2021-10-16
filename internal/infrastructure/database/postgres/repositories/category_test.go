package repositories_test

import (
	"testing"
	"time"

	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/repositories"
	"github.com/hizagi/esperto-bots/internal/infrastructure/database/postgres/utils"
	"github.com/hizagi/esperto-bots/internal/infrastructure/test/container/postgres"
	"gotest.tools/v3/assert"
)

func TestGetCategoryPostgres(t *testing.T) {
	postgresClient, _ := postgres.SetupDatabase(t)

	categoryRepository := repositories.NewCategoryRepository(postgresClient)

	category, err := categoryRepository.GetCategory("c2d29867-3d0b-d497-9191-18a9d8ee7830")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Vacuum", category.Name)
}

func TestListCategoriesPostgres(t *testing.T) {
	postgresClient, _ := postgres.SetupDatabase(t)

	categoryRepository := repositories.NewCategoryRepository(postgresClient)

	time, err := time.Parse(time.RFC3339, "2016-06-20T19:10:25-07:00")

	if err != nil {
		t.Fatal(err)
	}

	lastCursor := utils.EncodeCursor(time, "c2d29867-3d0b-d497-9191-18a9d8ee7830")

	categories, cursor, err := categoryRepository.ListCategory(lastCursor, 2)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(categories))
	assert.Equal(t, "2a3eb41e-2ccf-4357-a7cb-b0d491a53b96", categories[0].ID)
	assert.Equal(t, "c2d29867-3d0b-d497-9191-18a9d8ee7830", *categories[0].ParentID)

	_, id, err := utils.DecodeCursor(*cursor)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "87169909-33c1-4a2c-8c77-76cba19369f1", id)
}
