// Package calculator provides a library for // simple calculations in Go.
package calculator

import "fmt"

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
	if b == 0{
		return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)	
	}
	return a/b, nil
}
