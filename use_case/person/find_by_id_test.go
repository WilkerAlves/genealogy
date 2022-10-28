package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/stretchr/testify/assert"
)

func TestFindPersonById(t *testing.T) {
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

	uc := person.NewFindPersonByIdUseCase(r)
	p, err := uc.Execute(ctx, person1.ID)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.ID, person1.ID)
}

func TestFindPersonByIdWithInvalidId(t *testing.T) {
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
