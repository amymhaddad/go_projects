package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = bookstore.Catalog{
	1: {ID: 1, Title: "T1", Author: "A1", Copies: 1, Price: 5},
	2: {ID: 2, Title: "T2", Author: "A2", Copies: 1, Price: 5},
}

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "T1",
		Author: "A1",
		Copies: 1,
	}
}

func TestGetBooks(t *testing.T) {
	t.Parallel()

	want := []bookstore.Book{
		{ID: 1, Title: "T1", Author: "A1", Copies: 1, Price: 5},
		{ID: 2, Title: "T2", Author: "A2", Copies: 1, Price: 5},
	}

	got := bookstore.GetBooks(catalog)

	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}

}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	ID := 2
	want := "T1, by A1"
	got, err := bookstore.GetBookDetails(catalog, ID)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
	}

}

func TestBookDetailError(t *testing.T) {
	t.Parallel()

	ID := 3
	_, err := bookstore.GetBookDetails(catalog, ID)

	if err == nil {
		t.Fatal(err)
	}
}

func TestBuyBook(t *testing.T) {
	t.Parallel()

	ID := 1
	price := 5
	got := bookstore.BuyBook(catalog, price, ID)
	want := true

	if got != want {
		t.Errorf("want: %t, got: %t ", want, got)
	}

}
