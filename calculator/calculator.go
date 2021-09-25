// Package calculator provides a library for // simple calculations in Go.
package calculator

import (
	"errors"
	"math"
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
