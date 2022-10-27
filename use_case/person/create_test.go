package person_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/stretchr/testify/assert"
)

func TestNewCreatePerson(t *testing.T) {
	uc := person.NewCreatePersonUseCase(new(repository.PersonRepositoryMemory))
	name := "Bruce"
	p, err := uc.Execute(context.Background(), name)

	assert.Nil(t, err)
	assert.NotNil(t, p.ID)
	assert.Equal(t, p.Name, name)
}

func TestNewCreatePersonWithInvalidName(t *testing.T) {
	uc := person.NewCreatePersonUseCase(new(repository.PersonRepositoryMemory))
	name := ""
	p, err := uc.Execute(context.Background(), name)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "name invalid", err.Error())
}

func TestNewCreatePersonWithErrorForCreate(t *testing.T) {
	uc := person.NewCreatePersonUseCase(new(repository.PersonRepositoryMemory))
	name := "ERROR"
	p, err := uc.Execute(context.Background(), name)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, "error for create person: error for create person with name: ERROR", err.Error())
}
