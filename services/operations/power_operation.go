package operations

import (
	"math"
	"regexp"
)

type PowerOperation struct {
}

func (op PowerOperation) GetOperation() func(float64, float64) float64 {
	return func(a float64, b float64) float64 {
		return math.Pow(a, b)
	}
}

func (op PowerOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\^\d+(\.\d+)?`)
}
