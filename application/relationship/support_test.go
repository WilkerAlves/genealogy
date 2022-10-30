package relationship_test

import (
	"errors"
	"fmt"
	"os"
	"path"

	domainRepo "github.com/WilkerAlves/genealogy/domain/repository"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/joho/godotenv"
)

var (
	relationshipRepository domainRepo.RelationshipRepository
	conn                   string
)

func init() {
	rootPath := os.Getenv("ROOT_PATH")
	err := godotenv.Load(path.Join(rootPath, ".env"))
	if err != nil {
		panic(errors.New("error while load env"))
	}

	conn = getConnectionString()
	relationshipRepository, err = repository.NewRelationshipRepository(conn)
	if err != nil {
		panic(errors.New("error while create repository"))
	}
}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)
}
