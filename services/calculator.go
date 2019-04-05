package services

import (
	. "../interfaces"
	. "../value_objects"
	. "./operations"
	"fmt"
	"regexp"
	"strconv"
)

type Calculator struct {
	MathOperations   []MathOperation
	singleExpression *regexp.Regexp
}

func (c Calculator) Calculate(expressionString string) Expression {
	var expression Expression
	c.singleExpression = regexp.MustCompile(`[+|\-|*|^|/]`)

	c.MathOperations = append(c.MathOperations, PowerOperation{})
	c.MathOperations = append(c.MathOperations, MultiplicationOperation{})
	c.MathOperations = append(c.MathOperations, DivisionOperation{})
	c.MathOperations = append(c.MathOperations, AdditionOperation{})
	c.MathOperations = append(c.MathOperations, SubtractionOperation{})

	expression.ActualExpression = expressionString
	expression.Resolution = append(expression.Resolution, expression.ActualExpression)

	for _, item := range c.MathOperations {
		expression = c.ResolveAll(expression, item)
	}

	return expression
}

func (c Calculator) ResolveAll(expression Expression, mathOperation MathOperation) Expression {
	regex := mathOperation.GetRegex()
	operation := mathOperation.GetOperation()

	match := regex.MatchString(expression.ActualExpression)

	if !match {
		return expression
	}

	toCalc := regex.FindAllString(expression.ActualExpression, 1)
	leftParts := regex.Split(expression.ActualExpression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := operation(numbers[0], numbers[1])

	expression.ActualExpression = leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	expression.Resolution = append(expression.Resolution, expression.ActualExpression)

	return c.ResolveAll(expression, mathOperation)
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
