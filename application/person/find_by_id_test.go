package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/application/person"
	"github.com/stretchr/testify/assert"
)

func TestFindPersonById(t *testing.T) {
	ctx := context.Background()

	person1, err := creteUsers(ctx, nameNewUser)
	if err != nil {
		return
	}

	uc := person.NewFindPersonByIdUseCase(personRepository)
	p, err := uc.Execute(ctx, person1.ID)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.ID, person1.ID)
}

func TestFindPersonByIdWithInvalidId(t *testing.T) {
	ctx := context.Background()
	_, err := creteUsers(ctx, nameNewUser)
	if err != nil {
		return
	}

	newName := "NOVO NOME"
	uc := person.NewUpdatePersonUseCase(personRepository)
	err = uc.Execute(ctx, 0, newName)
	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "id invalid")
}
