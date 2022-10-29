package relationship

import (
	"github.com/WilkerAlves/genealogy/domain/repository"
)

type FindUseCase struct {
	repo repository.RelationshipRepository
}

type OutputFindRelationship struct {
	Parents   []string
	childrens []string
}

//func (uc *FindUseCase) Execute(ctx context.Context, id int) error {
//	output := make([]OutputFindRelationship, 0)
//
//	parents, err := uc.repo.GetParents(ctx, id)
//
//
//
//	childrens, err := uc.repo.GetParents(ctx, id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
