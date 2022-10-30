package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/application/person"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePerson(t *testing.T) {
	ctx := context.Background()

	person1, err := creteUsers(ctx, nameNewUser)
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

	person1, err := creteUsers(ctx, nameNewUser)
	if err != nil {
		return
	}

	newName := " "
	uc := person.NewUpdatePersonUseCase(personRepository)
	err = uc.Execute(ctx, person1.ID, newName)
	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "name invalid")
}

func TestUpdatePersonWithInvalidId(t *testing.T) {
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
