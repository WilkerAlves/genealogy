package relationship

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/entity"
	"github.com/WilkerAlves/genealogy/domain/repository"
)

type FindUseCase struct {
	relationshipRepository repository.RelationshipRepository
}

type OutputFamily struct {
	Parents      []*entity.Person `json:"parents"`
	Children     []*entity.Person `json:"children"`
	Nephews      []*entity.Person `json:"nephews"`
	Grandparents []*entity.Person `json:"grandparents"`
	Brothers     []*entity.Person `json:"brothers"`
}

func (uc *FindUseCase) Execute(ctx context.Context, id int) (*OutputFamily, error) {
	family := OutputFamily{}
	family.Parents = make([]*entity.Person, 0)
	family.Children = make([]*entity.Person, 0)

	parents, err := uc.relationshipRepository.GetParents(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(parents) > 0 {
		family.Parents = parents
	}

	children, err := uc.relationshipRepository.GetChildren(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(children) > 0 {
		family.Children = children
	}

	allparents := make(map[int][]*entity.Person)
	for i := range parents {
		id := parents[i].ID
		p, err := uc.relationshipRepository.GetParents(ctx, id)
		if err != nil {
			return nil, err
		}
		if len(p) < 1 {
			continue
		}
		allparents[id] = p
		parents = append(parents, p...)
	}

	for i := range children {
		c, err := uc.relationshipRepository.GetChildren(ctx, family.Parents[i].ID)
		if err != nil {
			return nil, err
		}
		if len(c) < 1 {
			continue
		}
		parents = append(parents, c...)

	}

	addBrother := make(map[int]string, 0)
	for i := range family.Parents {
		parent := family.Parents[i]

		children, err := uc.relationshipRepository.GetChildren(ctx, parent.ID)
		if err != nil {
			return nil, err
		}
		for c := range children {

			_, ok := addBrother[children[c].ID]
			if children[c].ID == id || ok {
				continue
			}
			family.Brothers = append(family.Brothers, children[c])
			addBrother[children[c].ID] = children[c].Name
		}

		gramas := allparents[parent.ID]
		for g := range gramas {
			family.Grandparents = append(family.Grandparents, gramas[g])
		}
	}

	for i := range family.Brothers {
		brother := family.Brothers[i]

		nephews, err := uc.relationshipRepository.GetChildren(ctx, brother.ID)
		if err != nil {
			return nil, err
		}

		family.Nephews = nephews

	}

	return &family, nil

}

func NewFindUseCase(repository repository.RelationshipRepository) *FindUseCase {
	return &FindUseCase{relationshipRepository: repository}
}
