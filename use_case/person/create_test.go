package person_test

import (
	"context"
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	repoDomain "github.com/WilkerAlves/genealogy/domain/repository"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var personRepository repoDomain.PersonRepository

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

func TestNewCreatePerson(t *testing.T) {
	uc := person.NewCreatePersonUseCase(personRepository)
	name := "Bruce"
	p, err := uc.Execute(context.Background(), name)

	assert.Nil(t, err)
	assert.NotNil(t, p.ID)
	assert.Equal(t, p.Name, name)
}

func TestNewCreatePersonWithInvalidName(t *testing.T) {
	uc := person.NewCreatePersonUseCase(personRepository)
	name := ""
	p, err := uc.Execute(context.Background(), name)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "name invalid", err.Error())
}
