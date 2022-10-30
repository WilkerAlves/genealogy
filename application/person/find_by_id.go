package person

import (
	"context"
	"errors"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/entity"
	"github.com/WilkerAlves/genealogy/domain/repository"
)

type FindPersonByIdUseCase struct {
	repository repository.PersonRepository
}

func (uc *FindPersonByIdUseCase) Execute(ctx context.Context, id int) (*entity.Person, error) {
	if id < 1 {
		return nil, errors.New("id invalid")
	}

	p, err := uc.repository.Get(ctx, id)
	if err != nil {
		msg := fmt.Sprintf("person %d not found", id)
		if err.Error() == msg {
			return nil, fmt.Errorf(msg)
		}

		return nil, fmt.Errorf("error for update person: %w", err)
	}

	return p, nil
}

func NewFindPersonByIdUseCase(repository repository.PersonRepository) *FindPersonByIdUseCase {
	return &FindPersonByIdUseCase{repository: repository}
}
