package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/application/person"
	"github.com/stretchr/testify/assert"
)

func TestNewCreatePerson(t *testing.T) {
	uc := person.NewCreatePersonUseCase(personRepository)
	name := "Test Create"
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
