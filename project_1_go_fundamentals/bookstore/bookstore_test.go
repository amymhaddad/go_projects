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

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "T1",
		Copies: 3,
	}
	testCases := []struct {
		Title  string
		Copies int
	}{
		{Title: "T1", Copies: 2},
	}

	for _, testCase := range TestCases {
		got := bookstore.Buy(b)

	}

}
