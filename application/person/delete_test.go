package person_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/WilkerAlves/genealogy/application/person"
	"github.com/stretchr/testify/assert"
)

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
