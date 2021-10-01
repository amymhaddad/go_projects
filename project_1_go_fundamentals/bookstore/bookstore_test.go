package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "T1",
		Author: "A1",
		Copies: 8,
	}

}


//Test 1, iteration 1
func TestBuy(t *testing.T) {
	t.Parallel()

	//Create an instance of the type Book 
	b := bookstore.Book{
		Title:  "T1",
		Author: "A1",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)

	if err == nil {
		t.Errorf("Want error, got nil")
	}
	// want := 9999
	// errExpected := true
	// result, err := bookstore.Buy(b)
	// got := result.Copies
	// errReceived := err != nil 
	
	// if errReceived  && errExpected {
	// 	t.Fatal("no more copies left")
	// }
	// //Check if want != to got
	// if want != got {
	// 	t.Errorf("want %d copies after buying 1 copy from a stock of 2, got %d", want, got)
	// }
}
