// Package calculator provides a library for // simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	//"fmt"
)

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
	return a/b, nil
}

func SqRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Square root of negative number is an error ")
	}	
	return math.Sqrt(a), nil

}
