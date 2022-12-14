package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestPersonRepositoryAdd(t *testing.T) {
	r, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)

	person, err := r.Add(context.Background(), "Bruce")
	assert.Nil(t, err)

	assert.NotNil(t, person.ID)
	assert.Greater(t, person.ID, 0)
}

func TestPersonRepositoryGet(t *testing.T) {
	ctx := context.Background()
	r, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)

	person1, err := r.Add(ctx, "Bruce")
	assert.Nil(t, err)
	assert.NotNil(t, person1.ID)
	assert.Greater(t, person1.ID, 0)
	person2, err := r.Add(ctx, "Sonny")
	assert.Nil(t, err)
	assert.NotNil(t, person2.ID)
	assert.Greater(t, person2.ID, 0)

	person, err := r.Get(ctx, person2.ID)
	assert.Nil(t, err)
	assert.Equal(t, person.ID, person2.ID)
}

func TestPersonRepositoryUpdate(t *testing.T) {
	ctx := context.Background()
	name := "NOME ATUALIZADO"
	r, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)

	person1, err := r.Add(ctx, "Bruce")
	assert.Nil(t, err)
	assert.NotNil(t, person1.ID)
	assert.Greater(t, person1.ID, 0)
	person2, err := r.Add(ctx, "Sonny")
	assert.Nil(t, err)
	assert.NotNil(t, person2.ID)
	assert.Greater(t, person2.ID, 0)

	err = r.Update(ctx, person2.ID, name)
	assert.Nil(t, err)

	person, err := r.Get(ctx, person2.ID)
	assert.Nil(t, err)
	assert.NotNil(t, person)
	assert.Equal(t, name, person.Name)
}

func TestPersonRepositoryDelete(t *testing.T) {
	ctx := context.Background()
	r, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)

	person1, err := r.Add(ctx, "Bruce")
	assert.Nil(t, err)
	assert.NotNil(t, person1.ID)
	assert.Greater(t, person1.ID, 0)
	person2, err := r.Add(ctx, "Sonny")
	assert.Nil(t, err)
	assert.NotNil(t, person2.ID)
	assert.Greater(t, person2.ID, 0)

	err = r.Delete(ctx, person2.ID)
	assert.Nil(t, err)

	person, err := r.Get(ctx, person2.ID)
	assert.Nil(t, person)
	assert.NotNil(t, err)
	assert.Error(t, err, fmt.Sprintf("person %d not found", person2.ID))
}
