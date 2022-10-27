package person_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/stretchr/testify/assert"
)

func TestDeletePerson(t *testing.T) {
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

	uc := person.NewDeletePersonUseCase(r)
	err = uc.Execute(ctx, person1.ID)
	assert.Nil(t, err)

	_, err = r.Get(ctx, person1.ID)
	assert.Equal(t, fmt.Sprintf("person not found in id: %d", person1.ID), err.Error())

}

func TestDeletePersonWithInvalidId(t *testing.T) {
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

	uc := person.NewDeletePersonUseCase(r)
	err = uc.Execute(ctx, 0)
	assert.NotNil(t, err)
	assert.Equal(t, "id invalid", err.Error())
}
