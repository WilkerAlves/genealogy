package relationship_test

import (
	"context"
	"errors"
	"os"
	"path"
	"testing"

	"github.com/WilkerAlves/genealogy/application/relationship"
	domainRepo "github.com/WilkerAlves/genealogy/domain/repository"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var relationshipRepository domainRepo.RelationshipRepository

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

func TestFindUseCase_Execute(t *testing.T) {
	uc := relationship.NewFindUseCase(relationshipRepository)
	id := 6
	family, err := uc.Execute(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, family)
}
