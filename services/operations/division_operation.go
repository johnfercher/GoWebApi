package operations

import "regexp"

type DivisionOperation struct {
}

func (op DivisionOperation) GetOperation() func(float64, float64) float64 {
	return op.Divide
}

func (op DivisionOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\/\d+(\.\d+)?`)
}

func (op DivisionOperation) Divide(a float64, b float64) float64 {
	return a / b
}
