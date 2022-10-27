package repository

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type PersonRepository interface {
	Add(ctx context.Context, name string) (*entity.Person, error)
	Get(ctx context.Context, id int) (*entity.Person, error)
	Update(ctx context.Context, id int, name string) error
	Delete(ctx context.Context, id int) error
}
