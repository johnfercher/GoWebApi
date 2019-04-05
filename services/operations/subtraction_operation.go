package operations

import "regexp"

type SubtractionOperation struct {
}

func (op SubtractionOperation) GetOperation() func(float64, float64) float64 {
	return func(a float64, b float64) float64 {
		return a - b
	}
}

func (op SubtractionOperation) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(`\d+(\.\d+)?\-\d+(\.\d+)?`)
}
