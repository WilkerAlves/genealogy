package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/domain/entity"
	repo "github.com/WilkerAlves/genealogy/domain/repository"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/stretchr/testify/assert"
)

func creteUsers(ctx context.Context, r repo.PersonRepository, name string) (*entity.Person, error) {
	add, err := r.Add(ctx, name)
	if err != nil {
		return nil, err
	}

	return add, nil
}

func TestUpdatePerson(t *testing.T) {
	ctx := context.Background()
	r := new(repository.PersonRepositoryMemory)

	person1, err := creteUsers(ctx, r, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, r, "Mike")
	if err != nil {
		return
	}

	newName := "NOVO NOME"
	uc := person.NewUpdatePersonUseCase(r)
	err = uc.Execute(ctx, person1.ID, newName)
	assert.Nil(t, err)

	p, err := r.Get(ctx, person1.ID)
	assert.Nil(t, err)

	assert.Equal(t, p.Name, person1.Name)
}

func TestUpdatePersonWithInvalidName(t *testing.T) {
	ctx := context.Background()
	r := new(repository.PersonRepositoryMemory)

	person1, err := creteUsers(ctx, r, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, r, "Mike")
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
	r := new(repository.PersonRepositoryMemory)

	_, err := creteUsers(ctx, r, "Bruce")
	if err != nil {
		return
	}

	_, err = creteUsers(ctx, r, "Mike")
	if err != nil {
		return
	}

	newName := "NOVO NOME"
	uc := person.NewUpdatePersonUseCase(new(repository.PersonRepositoryMemory))
	err = uc.Execute(ctx, 0, newName)
	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "id invalid")
}
