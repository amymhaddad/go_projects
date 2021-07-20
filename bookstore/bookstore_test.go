package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      "Marie Kondō",
		Description: "A tiny, cheerful Japanese woman explains tidying.",
		PriceCents:  1199,
		ID:          "1",
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	bookstore.Books = map[string]bookstore.Book{}
	bookstore.AddBook("Book1", "Problem Solving for Programmers", "Amy Haddad", "bbbb")

	want := bookstore.Books
	got := bookstore.GetAllBooks()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

// func TestNewID(t *testing.T) {
// 	t.Parallel()

// 	results := map[string]bool{}

// 	for i := 0; i < 10; i++ {
// 		id := bookstore.NewID()
// 		_, ok := results[id]
// 		if ok {
// 			t.Errorf("Invalid id: id already exists.")
// 		} else {
// 			results[id] = true
// 		}
// 	}

// }

// func TestAllByAuthor(t *testing.T) {
// 	t.Parallel()
// 	bookstore.Books = []bookstore.Book{}
// 	bookstore.AddBook("Book1", "Problem Solving for Programmers", "Amy Haddad", "bbbb")
// 	bookstore.AddBook("Book2", "Learn to PS", "Amy Haddad", "bbbb")

// 	want := "Problem Solving for Programmers\nLearn to PS"
// 	got := bookstore.GetAllByAuthor("Amy Haddad")

// 	if want != got {
// 		t.Errorf("got %s", got)
// 	}

// }

// func TestGetBookDetails(t *testing.T) {
// 	t.Parallel()
// 	want := "Title: Problem Solving for Programmers\nAuthor: Amy Haddad\nDescription: bbbb\nPriceCents: 0\nID: Book1"
// 	got := bookstore.GetBookDetails("Book1")

// 	if want != got {
// 		t.Errorf("No details for this book id:\ngot: %s\nwant: %s", got, want)

// 	}

// }
