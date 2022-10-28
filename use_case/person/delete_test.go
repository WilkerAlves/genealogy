package person_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
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
		personRepository, err = repository.NewPersonRepository(os.Getenv("CONNECTION_STRING_DB"))
		if err != nil {
			panic(errors.New("error create repository"))
		}

	} else {
		personRepository = new(repository.PersonRepositoryMemory)
	}
}

func TestDeletePerson(t *testing.T) {
	ctx := context.Background()

	person1, err := creteUsers(ctx, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, "Mike")
	if err != nil {
		return
	}

	uc := person.NewDeletePersonUseCase(personRepository)
	err = uc.Execute(ctx, person1.ID)
	assert.Nil(t, err)

	_, err = personRepository.Get(ctx, person1.ID)
	assert.Equal(t, fmt.Sprintf("person %d not found", person1.ID), err.Error())

}

func TestDeletePersonWithInvalidId(t *testing.T) {
	ctx := context.Background()
	r := new(repository.PersonRepositoryMemory)

	_, err := creteUsers(ctx, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, "Mike")
	if err != nil {
		return
	}

	uc := person.NewDeletePersonUseCase(r)
	err = uc.Execute(ctx, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "id invalid", err.Error())
}
