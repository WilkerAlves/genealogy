package relationship_test

import (
	"context"
	"errors"
	"os"
	"path"
	"testing"

	"github.com/WilkerAlves/genealogy/application/relationship"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	rootPath := os.Getenv("ROOT_PATH")
	err := godotenv.Load(path.Join(rootPath, ".env"))
	if err != nil {
		panic(errors.New("error while load env"))
	}
	relationshipRepository, err = repository.NewRelationshipRepository(os.Getenv("CONNECTION_STRING_DB"))
	if err != nil {
		panic(errors.New("error create repository"))
	}
}

func TestGetExecute(t *testing.T) {
	uc := relationship.NewGetUseCase(relationshipRepository)
	family, err := uc.Execute(context.Background(), 5, 13)
	assert.Nil(t, err)
	assert.NotNil(t, family)
}
