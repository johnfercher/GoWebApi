package operations

import "regexp"

type AdditionOperation struct {
}

func (op AdditionOperation) GetOperation() func(float64, float64) float64 {
	return op.Add
}

func (op AdditionOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\+\d+(\.\d+)?`)
}

func (op AdditionOperation) Add(a float64, b float64) float64 {
	return a + b
}
