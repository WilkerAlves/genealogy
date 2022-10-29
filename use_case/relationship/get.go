package relationship

import (
	"context"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/repository"
)

type GetUseCase struct {
	relationshipRepository repository.RelationshipRepository
}

func (uc *GetUseCase) Execute(ctx context.Context, idCompare1, idCompare2 int) (string, error) {
	family1, err := genealogy(ctx, idCompare1, uc.relationshipRepository)
	if err != nil {
		return "", err
	}

	for i := range family1.Parents {
		if family1.Parents[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is children this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Children {
		if family1.Children[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is parents this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Brothers {
		if family1.Brothers[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is brothers this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Grandparents {
		if family1.Grandparents[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is grandparents this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Uncles {
		if family1.Uncles[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is nephews this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Nephews {
		if family1.Nephews[i].ID == idCompare2 {
			return fmt.Sprintf("this id: %d is uncles this id: %d", family1.Parents[i].ID, idCompare2), nil
		}
	}

	for i := range family1.Cousins {
		if family1.Cousins[i].ID == idCompare2 {
			return "Cousins", nil
		}
	}

	return "Not relationship", nil
}

func NewGetUseCase(repository repository.RelationshipRepository) *GetUseCase {
	return &GetUseCase{relationshipRepository: repository}
}
