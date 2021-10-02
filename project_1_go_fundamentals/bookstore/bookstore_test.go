package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "T1",
		Author: "A1",
		Copies: 8,
	}

}

func TestBuy(t *testing.T) {
	t.Parallel()

	//Create an instance of the type Book 
	b := bookstore.Book{
		Title:  "T1",
		Author: "A1",
		Copies: 2,
	}

	want := 1 
	results, err := bookstore.Buy(b)

	if err != nil {
		t.Fatal(err)
	}

	if want != results.Copies {
		t.Errorf("want %d copies after buying 1 copy from a stock of 2, got %d", want, results.Copies)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying book when zero copies left, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	//Create my own slice of books to test 
	catalog := []bookstore.Book{
		{Title:  "T1"}, 
		{Title:  "T2",},
	}

	want := []bookstore.Book{
		{Title:  "T1"}, 
		{Title:  "T2"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}


}