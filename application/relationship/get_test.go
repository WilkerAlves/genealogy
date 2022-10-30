package relationship_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/application/relationship"
	"github.com/stretchr/testify/assert"
)

func TestGetExecute(t *testing.T) {
	uc := relationship.NewGetUseCase(relationshipRepository)
	family, err := uc.Execute(context.Background(), 5, 13)
	assert.Nil(t, err)
	assert.NotNil(t, family)
}
