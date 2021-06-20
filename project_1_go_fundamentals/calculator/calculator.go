// Package calculator provides a library for // simple calculations in Go.
package calculator

import (
	"errors"
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
		//Use errors.New() when I'm not using verbs to format a string. The test is only interested if the return value is Not nil. SO any message that's not nil would work here)
		//Anything can take the place of the return float64 bc we're interested in the non-nil return vaue 
		return 0, errors.New("Division by zero is undefined")
		//return 0, fmt.Errorf("bad input: %f, %f (division by zero is undefined)", a, b)	
	}
	return a/b, nil
}

//If NOT nil, then an error happened 