package calculator_test

import (
	"calculator"
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
			t.Errorf("want %f, got %f", tc.want, got) }
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
