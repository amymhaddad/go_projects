package calculator_test

import (
	"calculator"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func closeEnough(a, b, tolerance float64) bool { 
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	//make as part of function
	v1 := rand.Float64()
	v2 := rand.Float64()
	total := v1 + v2

	type testCase struct {
        a, b float64
        want float64
    }
	testCases := []testCase{
		//Ask about using multiple random values -- why error out? Each time I use v1 and v2 I should get a random number?
		{ a: v1, b: v2, want: total},
		//{a: v1, b: v2, want: total },
		 { a: 5, b: 0, want: 5 },
		{a: 0, b:0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("want: %f, got: %f ", tc.want, got)
		} else if got > 0 {
			fmt.Printf("\n%f plus %f sum to a positive number", tc.a, tc.b)
		} else if got < 0 {
			fmt.Printf("\n%f plus %f sum to a negative number", tc.a, tc.b)
		} else {
			fmt.Printf("\n%f minus % f is zero", tc.a, tc.b)
		}

	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase {
		{a: 4, b: 2, want: 2},
		{a: -6, b: 1, want: -7},
		{a: 0, b: 0, want: 0},
		//{a: 2.2, b: 1.5, want: 0.7},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		} else if got > 0 {
			fmt.Printf("\n%f minus %f is a positive number", tc.a, tc.b)
		} else if got < 0 {
			fmt.Printf("\n%f minus %f is a negative number", tc.a, tc.b)
		} else {
			fmt.Printf("\n%f minus % f is zero", tc.a, tc.b)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase {
		{a: 8, b: 8, want: 64},
		{a: 0, b: 200, want: 0},
		{a: -5, b: 5, want: -25},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		} else if got > 0{
			fmt.Printf("\n%f multiplied by %f results in a positive number", tc.a, tc.b)
		} else if got < 0{
			fmt.Printf("\n%f multiplied by %f results in a negative number", tc.a, tc.b)
		} else {
			fmt.Printf("\n%f multiplied by %f is zero\n", tc.a, tc.b)
		}

	}
}

func TestDivision(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b float64
		want float64
		errExpected bool
	}{
		{a: 3, b: 3, want:1, errExpected: false},
		//{a: 3, b: 0, want: 999, errExpected: true},
		{a: 1, b: 3, want: 0.333333, errExpected: false},
	}
	
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)
		fmt.Println(errReceived)

		if errReceived != tc.errExpected {
			fmt.Println(errReceived)
			t.Fatalf("Division (%.1f, %.1f): unexpected error %v", tc.a, tc.b, err)
		}
		//This was my original: if err == nil && !tc.errExpected && tc.want != got {...} -- this does NOT work w/decimals. I'll get a precision error
		//SO instead of comparing want and got (due to precision error), I want to see if my values are CLOSE ENOUGH

		if err == nil && !closeEnough(tc.want, got, 0.001)  {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

// func TestSqRoot(t *testing.T) {
// 	t.Parallel()

// 	testCases := []struct {
// 		a float64
// 		want float64 
// 	}{
// 		{a: 4, want: 2}, 
// 		{a: 64, want: 8},
// 	}
	
// 	for _, tc := range testCases {
// 		got := calculator.SqRoot(tc.a)

// 		if got != tc.want {
// 			fmt.Printf("want %f, got %f ", tc.want, got)
// 		}	
// 	}
// }

