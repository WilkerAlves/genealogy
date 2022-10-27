package person

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/WilkerAlves/genealogy/domain/repository"
)

type UpdatePersonUseCase struct {
	repository repository.PersonRepository
}

func (uc *UpdatePersonUseCase) Execute(ctx context.Context, id int, name string) error {
	if id < 1 {
		return errors.New("id invalid")
	}

	if len(strings.Trim(name, " ")) < 1 {
		return errors.New("name invalid")
	}

	if err := uc.repository.Update(ctx, id, name); err != nil {
		return fmt.Errorf("error for update person: %w", err)
	}

	return nil
}

func NewUpdatePersonUseCase(repository repository.PersonRepository) *UpdatePersonUseCase {
	return &UpdatePersonUseCase{repository: repository}
}
