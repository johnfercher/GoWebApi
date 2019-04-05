package services

import "regexp"

type DivisionOperation struct {
}

func (op DivisionOperation) GetOperation() func(float64, float64) float64 {
	return func(a float64, b float64) float64 {
		return a / b
	}
}

func (op DivisionOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\/\d+(\.\d+)?`)
}
