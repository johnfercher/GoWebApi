package operations

import (
	"math"
	"regexp"
)

type PowerOperation struct {
}

func (op PowerOperation) GetOperation() func(float64, float64) float64 {
	return op.Pow
}

func (op PowerOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\^\d+(\.\d+)?`)
}

func (op PowerOperation) Pow(a float64, b float64) float64 {
	return math.Pow(a, b)
}
