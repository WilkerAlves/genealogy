package relationship

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/repository"
)

type AddUseCase struct {
	repo repository.RelationshipRepository
}

type InputRelationship struct {
	Parent   int
	Children int
}

func (uc *AddUseCase) Execute(ctx context.Context, input InputRelationship) error {
	err := uc.repo.Add(ctx, input.Parent, input.Children)
	if err != nil {
		return err
	}

	return nil
}

func NewAddUseCase(repo repository.RelationshipRepository) *AddUseCase {
	return &AddUseCase{repo: repo}
}
