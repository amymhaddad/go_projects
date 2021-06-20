package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
        a, b float64
        want float64
    }
	testCases := []testCase{
		{ a: 2, b: 2, want: 4 },
		{ a: 1, b: 1, want: 2 },
		{ a: 5, b: 0, want: 5 },
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		} else if got > 0 {
			fmt.Printf("\n%f plus %f sum to a positive number", tc.a, tc.b)
		} else {
			fmt.Printf("\n%f plus %f sum to a negative number", tc.a, tc.b)
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
		//{a: 2.2, b: 1.5, want: 0.7},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
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
		}

	}
}

func TestDivision(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
		errExpected bool
	}

	testCases := []testCase {
		{a: 4, b: 2, want: 2, errExpected: false},
		{a: 4, b: 0, want: 0, errExpected: true},
		{a: 10, b: 10, want:1, errExpected: false},
	}

	for _, tc := range testCases {
		got, err:= calculator.Divide(tc.a, tc.b)
		if tc.want != got && err != nil {
			t.Errorf("want %f, got %f", tc.want, got)	
		}

	}	
}

