package person_test

import (
	"context"
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/WilkerAlves/genealogy/domain/entity"
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

func creteUsers(ctx context.Context, name string) (*entity.Person, error) {
	add, err := personRepository.Add(ctx, name)
	if err != nil {
		return nil, err
	}

	return add, nil
}

func TestUpdatePerson(t *testing.T) {
	ctx := context.Background()

	person1, err := creteUsers(ctx, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, "Mike")
	if err != nil {
		return
	}

	newName := "NOVO NOME"
	uc := person.NewUpdatePersonUseCase(personRepository)
	err = uc.Execute(ctx, person1.ID, newName)
	assert.Nil(t, err)

	p, err := personRepository.Get(ctx, person1.ID)
	assert.Nil(t, err)

	assert.Equal(t, newName, p.Name)
}

func TestUpdatePersonWithInvalidName(t *testing.T) {
	ctx := context.Background()
	r := new(repository.PersonRepositoryMemory)

	person1, err := creteUsers(ctx, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, "Mike")
	if err != nil {
		return
	}

	newName := " "
	uc := person.NewUpdatePersonUseCase(r)
	err = uc.Execute(ctx, person1.ID, newName)
	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "name invalid")
}

func TestUpdatePersonWithInvalidId(t *testing.T) {
	ctx := context.Background()
	_, err := creteUsers(ctx, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, "Mike")
	if err != nil {
		return
	}

	newName := "NOVO NOME"
	uc := person.NewUpdatePersonUseCase(new(repository.PersonRepositoryMemory))
	err = uc.Execute(ctx, 0, newName)
	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "id invalid")
}
