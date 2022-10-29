package repository_test

import (
	"context"
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/relationship"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	rootPath := os.Getenv("ROOT_PATH")
	if len(strings.Trim(rootPath, " ")) > 0 {
		err := godotenv.Load(path.Join(rootPath, ".env"))
		if err != nil {
			panic(errors.New("error while load env"))
		}
	}
}

func TestRelationshipRepositoryAdd(t *testing.T) {
	conn := os.Getenv("CONNECTION_STRING_DB")
	ctx := context.Background()

	perRepo, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)
	relaRepo, err := repository.NewRelationshipRepository(conn)
	assert.Nil(t, err)

	bruce, err := perRepo.Add(ctx, "Bruce")
	assert.Nil(t, err)

	mike, err := perRepo.Add(ctx, "Mike")
	assert.Nil(t, err)

	sonny, err := perRepo.Add(ctx, "Sonny")
	assert.Nil(t, err)

	phoebe, err := perRepo.Add(ctx, "Phoebe")
	assert.Nil(t, err)

	anastasia, err := perRepo.Add(ctx, "Anastasia")
	assert.Nil(t, err)

	martin, err := perRepo.Add(ctx, "Martin")
	assert.Nil(t, err)

	inputs := make([]relationship.InputRelationship, 0)

	inputs = append(
		inputs,
		relationship.InputRelationship{
			Parent:   mike.ID,
			Children: bruce.ID,
		},
		relationship.InputRelationship{
			Parent:   phoebe.ID,
			Children: bruce.ID,
		},
		relationship.InputRelationship{
			Parent:   sonny.ID,
			Children: mike.ID,
		},
		relationship.InputRelationship{
			Parent:   anastasia.ID,
			Children: phoebe.ID,
		},
		relationship.InputRelationship{
			Parent:   martin.ID,
			Children: phoebe.ID,
		},
	)

	for i := range inputs {
		err = relaRepo.Add(ctx, inputs[i].Parent, inputs[i].Children)
		assert.Nil(t, err)

	}

}
