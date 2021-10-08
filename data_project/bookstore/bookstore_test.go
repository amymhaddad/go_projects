package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Title 1",
		Author: "A1",
		Copies: 1,
	}
}

func TestGetBooks(t *testing.T) {
	t.Parallel()

	catalog := []bookstore.Book{
		{
			Title:  "Title 1",
			Author: "A1",
			Copies: 1,
		},
		{
			Title:  "Title 2",
			Author: "A2",
			Copies: 1,
		},
	}

	want := catalog
	got := bookstore.GetBooks(catalog)

	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
	}

}
