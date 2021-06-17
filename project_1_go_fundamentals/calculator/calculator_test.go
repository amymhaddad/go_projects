package calculator_test

import (
	"calculator"
	"testing"
)

// func TestAdd(t *testing.T) { t.Parallel()
// 	var want float64 = 4
// 	got := calculator.Add(2, 2)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got) }
// }

func TestAdd(t *testing.T) {
	t.Parallel()
	
	type testCase struct {
        a, b float64
        want float64
    }

	//create slice with literal syntax to populate immediately
	//Think: creating the 3 structs that are test cases for my inputs
	testCases := []testCase{
		{ a: 2, b: 2, want: 4 }, 
		{ a: 1, b: 1, want: 2 }, 
		{ a: 5, b: 0, want: 5 },
	}

	
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got) }
		}
	}


	/*
	1. Unclear why my function doesn't work?
	2. Also, unclear why I'm instructed to use this code:
	if tc.want != got {
		t.Errorf("want %f, got %f", tc.want, got)
	-->This code fails, too

	I declare "want" separtely, so why use: tc.want in the condintional?

	3. Want to make sure I udnerstand the struct sytnax w/slice:
	-Create a struct
	-Create a slice of TYPE testCase and insert a map as k, v pairs?

	*/


func TestSubtract(t *testing.T) { 
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2) 
	if want != got {
		t.Errorf("want %f, got %f", want, got) }
}

func TestMultiply(t *testing.T) { t.Parallel()
	var want float64 = 4
	got := calculator.Multiply(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
