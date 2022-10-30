package repository_test

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	rootPath := os.Getenv("ROOT_PATH")
	err := godotenv.Load(path.Join(rootPath, ".env"))
	if err != nil {
		panic(errors.New("error while load env"))
	}
}

func TestConfigureDb(t *testing.T) {
	conn := os.Getenv("CONNECTION_STRING_DB")
	db, err := repository.ConfigureDb(conn)
	if err != nil {
		return
	}

	assert.Nil(t, err)
	assert.NotNil(t, db)
}
