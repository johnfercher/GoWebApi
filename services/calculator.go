package services

import (
	"fmt"
	"regexp"
	"strconv"
)

type Calculator struct {
	MathOperations   []MathOperation
	singleExpression *regexp.Regexp
}

func (c Calculator) Calculate(expression string) string {
	c.singleExpression = regexp.MustCompile(`[+|\-|*|^|/]`)

	c.MathOperations = append(c.MathOperations, PowerOperation{})
	c.MathOperations = append(c.MathOperations, MultiplicationOperation{})
	c.MathOperations = append(c.MathOperations, DivisionOperation{})
	c.MathOperations = append(c.MathOperations, AdditionOperation{})
	c.MathOperations = append(c.MathOperations, SubtractionOperation{})

	fmt.Println(expression)

	for _, item := range c.MathOperations {
		expression = c.ResolveAll(expression, item)
	}

	return expression
}

func (c Calculator) ResolveAll(expression string, mathOperation MathOperation) string {
	regex := mathOperation.GetRegex()
	operation := mathOperation.GetOperation()

	match := regex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := regex.FindAllString(expression, 1)
	leftParts := regex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := operation(numbers[0], numbers[1])

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolveAll(newExpression, mathOperation)
}

func (c Calculator) GetNumbers(singleExpression string) []float64 {
	var numbers []float64
	stringNumbers := c.singleExpression.Split(singleExpression, 3)

	for _, stringItem := range stringNumbers {
		num, _ := strconv.ParseFloat(stringItem, 64)
		numbers = append(numbers, num)
	}

	return numbers
}
