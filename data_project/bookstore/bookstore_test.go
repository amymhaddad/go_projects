package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = []bookstore.Book{
	{
		ID:     "1",
		Title:  "Title 1",
		Author: "A1",
		Copies: 1,
	},
	{
		ID:     "2",
		Title:  "Title 2",
		Author: "A2",
		Copies: 1,
	},
}

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

	// catalog := []bookstore.Book{
	// 	{
	// 		Title:  "Title 1",
	// 		Author: "A1",
	// 		Copies: 1,
	// 	},
	// 	{
	// 		Title:  "Title 2",
	// 		Author: "A2",
	// 		Copies: 1,
	// 	},
	// }

	want := catalog
	got := bookstore.GetBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}

}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	want := "Title 1, by A1"

	got := bookstore.GetBookDetails(catalog, "1")
	fmt.Println(got == want)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}

}

// func TestGetBookDetails(t *testing.T) {
// 	t.Parallel()
// 	catalog := map[int]bookstore.Book{
// 		1: {Title: "T1", Author: "A1", Copies: 1},
// 		2: {Title: "T2", Author: "A2", Copies: 1},
// 	}

// 	ID := 2

// 	want := bookstore.Book{Title: "T1", Author: "A1", Copies: 1}

// 	got := bookstore.GetBookDetails(catalog, ID)

// 	if !cmp.Equal(want, got) {
// 		cmp.Diff(want, got)
// 	}

// }
