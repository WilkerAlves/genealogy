package relationship

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/repository"
)

type FindUseCase struct {
	relationshipRepository repository.RelationshipRepository
}

func (uc *FindUseCase) Execute(ctx context.Context, id int) (*OutputFamily, error) {
	return genealogy(ctx, id, uc.relationshipRepository)
}

func NewFindUseCase(repository repository.RelationshipRepository) *FindUseCase {
	return &FindUseCase{relationshipRepository: repository}
}
