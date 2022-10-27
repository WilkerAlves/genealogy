package person

import (
	"context"
	"errors"
	"fmt"

	"github.com/WilkerAlves/genealogy/domain/repository"
)

type DeletePersonUseCase struct {
	repository repository.PersonRepository
}

func (uc *DeletePersonUseCase) Execute(ctx context.Context, id int) error {
	if id < 1 {
		return errors.New("id invalid")
	}

	if err := uc.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error for delete person: %w", err)
	}

	return nil
}

func NewDeletePersonUseCase(repository repository.PersonRepository) *DeletePersonUseCase {
	return &DeletePersonUseCase{repository: repository}
}
