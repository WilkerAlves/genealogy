package relationship

import (
	"context"

	"github.com/WilkerAlves/genealogy/domain/entity"
	"github.com/WilkerAlves/genealogy/domain/repository"
)

type OutputFamily struct {
	Parents      []*entity.Person `json:"parents"`
	Children     []*entity.Person `json:"children"`
	Nephews      []*entity.Person `json:"nephews"`
	Grandparents []*entity.Person `json:"grandparents"`
	Brothers     []*entity.Person `json:"brothers"`
	Uncles       []*entity.Person `json:"uncles"`
	Cousins      []*entity.Person `json:"cousins"`
}

func genealogy(ctx context.Context, id int, relationshipRepository repository.RelationshipRepository) (*OutputFamily, error) {
	family := OutputFamily{}
	family.Parents = make([]*entity.Person, 0)
	family.Children = make([]*entity.Person, 0)

	parents, err := relationshipRepository.GetParents(ctx, id)
	if err != nil {
		return nil, err
	}

	mapParents := make(map[int]string)
	if len(parents) > 0 {
		for i := range parents {
			family.Parents = append(family.Parents, parents[i])
			mapParents[parents[i].ID] = parents[i].Name
		}
	}

	children, err := relationshipRepository.GetChildren(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(children) > 0 {
		family.Children = children
	}

	allparents := make(map[int][]*entity.Person)
	for i := range parents {
		id := parents[i].ID
		p, err := relationshipRepository.GetParents(ctx, id)
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
		c, err := relationshipRepository.GetChildren(ctx, family.Parents[i].ID)
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

		children, err := relationshipRepository.GetChildren(ctx, parent.ID)
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

		nephews, err := relationshipRepository.GetChildren(ctx, brother.ID)
		if err != nil {
			return nil, err
		}

		family.Nephews = nephews

	}

	mapCousins := make(map[int]string)
	mapUncles := make(map[int]string)
	for i := range family.Grandparents {
		uncles, err := relationshipRepository.GetChildren(ctx, family.Grandparents[i].ID)
		if err != nil {
			return nil, err
		}

		for u := range uncles {
			id := uncles[u].ID
			_, myParent := mapParents[id]
			if myParent {
				continue
			}

			_, already := mapUncles[id]
			if already {
				continue
			}

			cousins, err := relationshipRepository.GetChildren(ctx, id)
			if err != nil {
				return nil, err
			}
			for c := range cousins {
				id := cousins[c].ID
				_, already := mapCousins[id]
				if already {
					continue
				}
				family.Cousins = append(family.Cousins, cousins[c])
				mapCousins[id] = cousins[c].Name
			}

			family.Uncles = append(family.Uncles, uncles[u])
			mapUncles[id] = uncles[u].Name
		}

	}

	return &family, nil
}
