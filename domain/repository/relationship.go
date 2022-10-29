package repository

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/entity"
)

type RelationshipRepository interface {
	Add(ctx context.Context, idParent, idChildren int) error
	GetParents(ctx context.Context, idChildren int) ([]*entity.Person, error)
	GetChild(ctx context.Context, idParent int) ([]*entity.Person, error)
}
