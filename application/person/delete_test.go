package person_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/WilkerAlves/genealogy/application/person"
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
	personRepository, err = repository.NewPersonRepository(os.Getenv("CONNECTION_STRING_DB"))
	if err != nil {
		panic(errors.New("error create repository"))
	}
}

func TestDeletePerson(t *testing.T) {
	ctx := context.Background()

	person1, err := creteUsers(ctx, "Test")
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

	_, err := creteUsers(ctx, "Test")
	if err != nil {
		return
	}

	uc := person.NewDeletePersonUseCase(personRepository)
	err = uc.Execute(ctx, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "id invalid", err.Error())
}
