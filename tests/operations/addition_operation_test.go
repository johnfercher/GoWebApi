package operations

import (
	. "../../services/operations"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestAdditionOperationAddShouldWorksProperly(t *testing.T) {
	// Arrange
	additionOperation := AdditionOperation{}
	numA := rand.Float64()
	numB := rand.Float64()

	// Act
	result := additionOperation.Add(numA, numB)

	// Assert
	assert.Equal(t, result, numA+numB)
}
