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

/*
This syntax doesn't work: 	val1 := rand.Float64()
Here's why:
Remove the colon : from the second statement as you are assigning a new value to existing variable.
myArray = [...]int{11,12,14}
colon : is used when you perform the short declaration and assignment for the first time as you are doing in your first statement i.e. myArray  :=[...]int{12,14,26}.
*/
func generateRandomValues() (val1, val2, total float64) {
	rand.Seed(time.Now().UnixNano())
	val1 = rand.Float64()
	val2 = rand.Float64()
	total = val1 + val2
	return val1, val2, total
}

func TestAdd(t *testing.T) {
	t.Parallel()


	val1, val2, total := generateRandomValues()


	type testCase struct {
        a, b float64
        want float64
    }
	testCases := []testCase{
		{a: val1, b: val2, want: total},
		{a: 5, b: 0, want: 5 },
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

func TestSqRoot(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a float64
		want float64 
		errExpected bool
	}{
		{a: 4, want: 2, errExpected: false}, 
		{a: 64, want: 8, errExpected: false},
		{a: 0, want: 0, errExpected: false},
		{a: -64, want: 1111, errExpected: true},
	}
	
	for _, tc := range testCases {
		got, err := calculator.SqRoot(tc.a)
		errReceived := err != nil

		if errReceived != tc.errExpected {
			t.Fatalf("Square Root (%f) unexpected error: %v", tc.a, err)
		}

		if err == nil && tc.want != got {
			t.Errorf("want: %f, got: %f ", tc.want, got)
		}

	}
}

