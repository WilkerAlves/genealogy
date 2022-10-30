package relationship_test

import (
	"context"
	"testing"

	"github.com/WilkerAlves/genealogy/application/relationship"

	"github.com/stretchr/testify/assert"
)

func TestFindUseCase_Execute(t *testing.T) {
	uc := relationship.NewFindUseCase(relationshipRepository)
	id := 1
	family, err := uc.Execute(context.Background(), id)
	assert.Nil(t, err)
	assert.NotNil(t, family)
}
