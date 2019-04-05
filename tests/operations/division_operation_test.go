package operations

import (
	. "../../services/operations"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDivisionOperationAddShouldWorksProperly(t *testing.T) {
	// Arrange
	divisionOperation := DivisionOperation{}
	numA := rand.Float64()
	numB := rand.Float64()

	// Act
	result := divisionOperation.Divide(numA, numB)

	// Assert
	assert.Equal(t, result, numA/numB)
}
