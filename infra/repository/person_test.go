package repository_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
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
	}
}

func TestPersonRepositoryAdd(t *testing.T) {
	conn := os.Getenv("CONNECTION_STRING_DB")

	r, err := repository.NewPersonRepository(conn)
	assert.Nil(t, err)

	person, err := r.Add(context.Background(), "Bruce")
	assert.Nil(t, err)

	assert.NotNil(t, person.ID)
	assert.Greater(t, person.ID, 0)
}

func TestPersonRepositoryGet(t *testing.T) {
	conn := os.Getenv("CONNECTION_STRING_DB")
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
	conn := os.Getenv("CONNECTION_STRING_DB")
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
	conn := os.Getenv("CONNECTION_STRING_DB")
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
