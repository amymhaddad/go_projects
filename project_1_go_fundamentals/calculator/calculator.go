// Package calculator provides a library for // simple calculations in Go.
package calculator

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type arithmeticExpression struct {
	num1      float64
	num2      float64
	operation string
}

func parseExpression(expression string) *arithmeticExpression {
	var expr arithmeticExpression

	re := regexp.MustCompile(`([0-9]+) (.) ([0-9]+)`)
	result := re.FindStringSubmatch(expression)

	firstNum, _ := strconv.ParseFloat(result[1], 64)
	secondNum, _ := strconv.ParseFloat(result[3], 64)
	expr.num1 = firstNum
	expr.num2 = secondNum
	expr.operation = result[2]

	return &expr
}

func calcExpression(expression *arithmeticExpression) {
	switch expression.operation {
	case "*":
		fmt.Println(Multiply(expression.num1, expression.num2))
	case "+":
		fmt.Println(Add(expression.num1, expression.num2))
	case "/":
		fmt.Println(Divide(expression.num1, expression.num2))
	case "-":
		fmt.Println(Subtract(expression.num1, expression.num2))
	default:
		print(fmt.Println("Invalid expression"))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an arithmetic expression: ")
	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	userExpression := scanner.Text()

	expression := parseExpression(userExpression)
	calcExpression(expression)
}

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero is undefined")
	}
	return a / b, nil
}

func SqRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Square root of negative number is an error ")
	}
	return math.Sqrt(a), nil

}
