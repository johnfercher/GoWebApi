package services

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Calculator struct {
	multiplyRegex    *regexp.Regexp
	divideRegex      *regexp.Regexp
	sumRegex         *regexp.Regexp
	minusRegex       *regexp.Regexp
	powRegex         *regexp.Regexp
	singleExpression *regexp.Regexp
}

func (c Calculator) Calculate(expression string) string {
	c.multiplyRegex = regexp.MustCompile(`\d+(\.\d+)?\*\d+(\.\d+)?`)
	c.divideRegex = regexp.MustCompile(`\d+(\.\d+)?\/\d+(\.\d+)?`)
	c.sumRegex = regexp.MustCompile(`\d+(\.\d+)?\+\d+(\.\d+)?`)
	c.minusRegex = regexp.MustCompile(`\d+(\.\d+)?\-\d+(\.\d+)?`)
	c.powRegex = regexp.MustCompile(`\d+(\.\d+)?\^\d+(\.\d+)?`)
	c.singleExpression = regexp.MustCompile(`[+|\-|*|^|/]`)

	fmt.Println(expression)

	expressionWithoutPow := c.ResolvePows(expression)
	expressionWithoutMultiplication := c.ResolveMultiplications(expressionWithoutPow)
	expressionWithoutDivision := c.ResolveDivisions(expressionWithoutMultiplication)
	expressionWithoutSum := c.ResolveSums(expressionWithoutDivision)
	expressionWithoutMinus := c.ResolveMinus(expressionWithoutSum)

	return expressionWithoutMinus
}

func (c Calculator) ResolvePows(expression string) string {
	match := c.powRegex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := c.powRegex.FindAllString(expression, 1)
	leftParts := c.powRegex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := math.Pow(numbers[0], numbers[1])

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolvePows(newExpression)
}

func (c Calculator) ResolveMultiplications(expression string) string {
	match := c.multiplyRegex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := c.multiplyRegex.FindAllString(expression, 1)
	leftParts := c.multiplyRegex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := numbers[0] * numbers[1]

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolveMultiplications(newExpression)
}

func (c Calculator) ResolveDivisions(expression string) string {
	match := c.divideRegex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := c.divideRegex.FindAllString(expression, 1)
	leftParts := c.divideRegex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := numbers[0] / numbers[1]

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolveDivisions(newExpression)
}

func (c Calculator) ResolveSums(expression string) string {
	match := c.sumRegex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := c.sumRegex.FindAllString(expression, 1)
	leftParts := c.sumRegex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := numbers[0] + numbers[1]

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolveSums(newExpression)
}

func (c Calculator) ResolveMinus(expression string) string {
	match := c.minusRegex.MatchString(expression)

	if !match {
		return expression
	}

	toCalc := c.minusRegex.FindAllString(expression, 1)
	leftParts := c.minusRegex.Split(expression, 2)

	numbers := c.GetNumbers(toCalc[0])

	result := numbers[0] - numbers[1]

	newExpression := leftParts[0] + fmt.Sprintf("%f", result) + leftParts[1]
	fmt.Println(newExpression)

	return c.ResolveMinus(newExpression)
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
