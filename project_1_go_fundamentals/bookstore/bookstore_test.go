package bookstore_test

import (
	"bookstore"
	"sort"
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

	catalog := map[int]bookstore.Book{
		1: {Title: "T1", ID: 1},
		2: {Title: "T2", ID: 2},
		3: {Title: "T3", ID: 3},
	}
	want := []bookstore.Book{
		{Title: "T1", ID: 1},
		{Title: "T2", ID: 2},
		{Title: "T3", ID: 3},
	}

	got := bookstore.GetAllBooks(catalog)
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {Title: "T1", ID: 1},
		2: {Title: "T2", ID: 2},
	}

	want := bookstore.Book{Title: "T2", ID: 2}

	bookID := 2
	got, err := bookstore.GetBook(catalog, bookID)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want doesn't equal got")
	}
}

func TestBookWithBadIDReturnsError(t *testing.T) {
	t.Parallel()

	catalog := map[int]bookstore.Book{}
	bookID := 2
	_, err := bookstore.GetBook(catalog, bookID)

	if err == nil {
		t.Fatal("want an error for non-existing ID, but got nil")
	}

}
