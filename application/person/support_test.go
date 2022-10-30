package person_test

import (
	"context"
	"errors"
	"os"
	"path"

	"github.com/WilkerAlves/genealogy/domain/entity"
	repoDomain "github.com/WilkerAlves/genealogy/domain/repository"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/joho/godotenv"
)

var (
	personRepository repoDomain.PersonRepository
	nameNewUser      = "Test"
)

func init() {
	rootPath := os.Getenv("ROOT_PATH")
	err := godotenv.Load(path.Join(rootPath, ".env"))
	if err != nil {
		panic(errors.New("error while load env"))
	}
	conn := os.Getenv("CONNECTION_STRING_DB_LOCAL")
	personRepository, err = repository.NewPersonRepository(conn)
	if err != nil {
		panic(errors.New("error create repository"))
	}
}

func creteUsers(ctx context.Context, name string) (*entity.Person, error) {
	add, err := personRepository.Add(ctx, name)
	if err != nil {
		return nil, err
	}

	return add, nil
}
