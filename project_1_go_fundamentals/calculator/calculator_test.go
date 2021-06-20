package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	t.Parallel()

// 	type testCase struct {
//         a, b float64
//         want float64
//     }
// 	testCases := []testCase{
// 		{ a: 2, b: 2, want: 4 },
// 		{ a: 1, b: 1, want: 2 },
// 		{ a: 5, b: 0, want: 5 },
// 		{a: 0, b:0, want: 0},
// 	}

// 	for _, tc := range testCases {
// 		got := calculator.Add(tc.a, tc.b)
// 		if tc.want != got {
// 			t.Errorf("want %f, got %f", tc.want, got)
// 		} else if got > 0 {
// 			fmt.Printf("\n%f plus %f sum to a positive number", tc.a, tc.b)
// 		} else if got < 0 {
// 			fmt.Printf("\n%f plus %f sum to a negative number", tc.a, tc.b)
// 		} else {
// 			fmt.Printf("\n%f minus % f is zero", tc.a, tc.b)
// 		}

// 	}

// }

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

	type testCase struct {
		a, b float64
		want float64
		errExpected bool
	}

	testCases := []testCase {

		{a: 3, b: 3, want:1, errExpected: false},
		{a: 3, b: 0, want: 999, errExpected: true},
		// //Use an unlikely value for "want" bc If I see this value then I know something went wrong with the test
		// {a: 4, b: 0, want: 99, errExpected: true},
	
	}


	for _, tc := range testCases {
		got, err:= calculator.Divide(tc.a, tc.b)
		//Did i rec an err? Was that err val nil or not?
		//create a var to determine if there was an error, then is val is T. Otherwise there was NO error, so this val is false
		errRecieved := (err != nil)

		//IF the rec'd output doens't match my expectations, then raise an error. 
		//I always need to check what the program returns against my expectatations
		//This syntax checks 2 things: If there was an error and I didn't expect one OR I didn't get an error and I exptedted one. 
		//So I'm looking to see if there's a mismatch between the error status that I rec'd (via: errReceived) and the error status that I expected (via: tc.errExpected)
		if errRecieved != tc.errExpected {
			//Don't move to any other tests
			t.Fatalf("Divide (%.1f, %.1f): unexpected error status: %v ", tc.a, tc.b, err)	
		} 

		//alt: if !tc.errExpected && tc.want != got  -->same thing: checking my expectations
		//Only want to compare want and got (via: && tc.want != got) when we don't expect an error
		//Recall: if first val is T, then return it. Otherwise move on to the next value.
		if err == nil && tc.want != got  {
			t.Errorf("want %f, got %f", tc.want, got)	
		}
        
	}	
}

//If just printing a message: t.Error() 
//If printing w/values: t.Errorf()
/*
Don't move to any other tests:
t.Fatalf("Divide (%.1f, %.1f): want error, got nil", tc.a, tc.b)	
%v - prints out any val of any type. So I can use it to print the value of the error

*/

