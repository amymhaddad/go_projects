package calculator_test

import (
	"calculator"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	v1 := rand.Float64()
	v2 := rand.Float64()
	total := v1 + v2

	fmt.Print("total: ", total)

	type testCase struct {
        a, b float64
        want float64
    }
	testCases := []testCase{
		{ a: v1, b: v1, want: total},
		// { a: 1, b: 1, want: 2 },
		// { a: 5, b: 0, want: 5 },
		// {a: 0, b:0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		fmt.Print("got: ", got)
		fmt.Print("want: ", tc.want)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
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
		{a: 3, b: 0, want: 999, errExpected: true},
	}
	
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)

		if errReceived != tc.errExpected {
			t.Fatalf("Division (%.1f, %.1f): unexpected error %v", tc.a, tc.b, err)
		}
	
		if err == nil && tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	 }
}

