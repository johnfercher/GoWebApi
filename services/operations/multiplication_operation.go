package operations

import "regexp"

type MultiplicationOperation struct {
}

func (op MultiplicationOperation) GetOperation() func(float64, float64) float64 {
	return op.Multiply
}

func (op MultiplicationOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\*\d+(\.\d+)?`)
}

func (op MultiplicationOperation) Multiply(a float64, b float64) float64 {
	return a * b
}
