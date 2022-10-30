package person

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/WilkerAlves/genealogy/domain/entity"
	"github.com/WilkerAlves/genealogy/domain/repository"
)

type CreatePersonUseCase struct {
	repository repository.PersonRepository
}

func (uc *CreatePersonUseCase) Execute(ctx context.Context, name string) (*entity.Person, error) {

	if len(strings.Trim(name, " ")) < 1 {
		return nil, errors.New("name invalid")
	}

	person, err := uc.repository.Add(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("error for create person: %w", err)
	}

	return person, nil
}

func NewCreatePersonUseCase(repository repository.PersonRepository) *CreatePersonUseCase {
	return &CreatePersonUseCase{repository: repository}
}
