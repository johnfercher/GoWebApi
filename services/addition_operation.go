package services

import "regexp"

type AdditionOperation struct {
}

func (op AdditionOperation) GetOperation() func(float64, float64) float64 {
	return func(a float64, b float64) float64 {
		return a + b
	}
}

func (op AdditionOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\+\d+(\.\d+)?`)
}
