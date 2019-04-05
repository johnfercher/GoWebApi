package services

import "regexp"

type MathOperation interface {
	GetRegex() *regexp.Regexp
	GetOperation() func(float64, float64) float64
}
