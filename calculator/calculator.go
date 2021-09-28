// Package calculator provides a library for // simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Add takes two numbers and returns the // result of adding them together.
func Add(a, b float64) float64 {

	return a + b
}

// Subtract takes two numbers a and b, and
// returns the result of subtracting b from a.
func Subtract(a, b float64) float64 {
	return a - b
}

//Multiply takes two numbers a and b, and
//returns the result of multiplying a and b
func Multiply(a, b float64) float64 {
	return a * b
}

//Divide takes two number a and b, and returns
//the result of dividing a and b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero error")
	}
	return a / b, nil
}

//SquareRoot takes a number and returns its square root
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Can't get the square root of a negative number")
	}
	return math.Sqrt(a), nil
}

//Expressions receives a string and returns the arithmetic evaluation of the
//expression
func Expressions(expr string) float64 {
	return 0

}

//ArithExpr is a type that contains an expression
type ArithExpr struct {
	num1      float64
	num2      float64
	operation string
}

// //escape spec chars (ie, * +) bc these values have other regex meanings; use [] to match 1 instance
// //index 0 matches everyting in regex
func parseExpression(expression string) *ArithExpr {
	var expr ArithExpr
	re := regexp.MustCompile(`(\d) ([\*\+-/]) (\d)`)
	result := re.FindStringSubmatch(expression)

	expr.num1, _ = strconv.ParseFloat(result[1], 64)
	expr.num2, _ = strconv.ParseFloat(result[3], 64)
	expr.operation = result[2]
	return &expr
}

// /*
//bad practice bc I'm relying on side effects -- I'm not returning enaythign from parseExpression.
//It's better to return value that I can pass around which is why i should use pointer semantics

// */

func calculate(expr *ArithExpr) {
	switch expr.operation {
	case "*":
		fmt.Println(Multiply(expr.num1, expr.num2))
	case "-":
		fmt.Println(Subtract(expr.num1, expr.num2))
	case "+":
		fmt.Println(Add(expr.num1, expr.num2))
	case "/":
		fmt.Println(Divide(expr.num1, expr.num2))
	default:
		fmt.Println("Invalid expression")
	}

}