// Package calculator provides a library for // simple calculations in Go.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


type arithmeticExpression struct {
	num1 float64
	num2 float64
	operation string
}

func getArithmeticExpression() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an arithmetic expression: ")
    scanner.Scan()
   
	if scanner.Err() != nil {
       fmt.Println("Error: ", scanner.Err())
    }
    expression := scanner.Text()
	return expression
}

//Write a test that returns a struct
func parseExpression(expression string) *arithmeticExpression {
	var expr arithmeticExpression
	parsedExpression := strings.Split(expression, " ")

	for i, value := range parsedExpression{
		num, err := strconv.ParseFloat(value, 64)
		if err == nil {
			if i == 0 {
				expr.num1 = num
			} else {
				expr.num2 = num
			}
			
		} else {
			expr.operation = value
		}
	}
	return &expr
 }

func calcExpression(expression *arithmeticExpression) float64 {
	//Create switch statements to determine the which funciton to call for the string operation 
	//That's included in the struct
}

 func main() {
	x := getArithmeticExpression()	
	fmt.Println(parseExpression(x))

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
	return a/b, nil
}

func SqRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Square root of negative number is an error ")
	}	
	return math.Sqrt(a), nil

}


