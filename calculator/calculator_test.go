package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b float64
		want float64
	}{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 2, want: 3},
	}

	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)

		if testCase.want != got {
			t.Errorf("Add (%f, %f): want %f, got %f", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b float64
		want float64
	}{
		{a: 6, b: 4, want: 2},
		{a: -1, b: 5, want: -6},
	}

	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract (%f, %f): want %f, got %f", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b float64
		want float64
	}{
		{a: 0, b: -1, want: 0},
		{a: 5, b: 2, want: 10.0},
	}

	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)

		if testCase.want != got {
			t.Errorf("Multiply: (%f, %f) want: %f, got: %f", testCase.a, testCase.b, testCase.want, got)
		}

	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b        float64
		want        float64
		errExpected bool
	}{
		{a: 4, b: 2, want: 2, errExpected: false},
		{a: 6, b: 0, want: 999, errExpected: true},
		{a: 1, b: 3, want: 0.333333, errExpected: false},
	}

	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)
		errReceived := err != nil

		if errReceived != testCase.errExpected {
			t.Fatalf("Division by 0 error")
		}

		if err == nil && !closeEnough(testCase.want, got, 0.001) {
			t.Errorf("Division (%f, %f) want: %f, got: %f", testCase.a, testCase.b, testCase.want, got)
		}

	}

}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a           float64
		want        float64
		errExpected bool
	}{
		{a: 4, want: 2, errExpected: false},
		{a: -4, want: 999, errExpected: true},
	}

	for _, testCase := range testCases {
		got, err := calculator.SquareRoot(testCase.a)
		errReceived := err != nil

		if errReceived != testCase.errExpected {
			t.Fatalf("Can't get the square root of a negative number")
		}

		if !testCase.errExpected && testCase.want != got {
			t.Errorf("SqureRoot (%f), want: %f, got: %f", testCase.a, testCase.want, got)
		}
	}

}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
