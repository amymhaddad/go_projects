package bookstore_test

import (
	"bookstore"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

//ex of the compile only test
func TestBook(t *testing.T) {
	//note syntax: bookstore.Book {} --> creating an instance of Book struct
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      "Marie Kondō",
		Discription: "A tiny, cheerful Japanese woman explains tidying.",
		PriceCents:  1199,
		ID:          "1",
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	//syntax says: in the bookstore module, get the Book struct and create a new intance of type Book
	book1 := bookstore.Book{Title: "Problem Solving for Programmers"}
	book2 := bookstore.Book{Title: "Learn to PS"}

	//Create a slice of type Book
	bookstore.Books = []bookstore.Book{book1, book2}

	fmt.Println("bookstore: ", bookstore.Books)

	want := bookstore.Books
	got := bookstore.GetAllBooks()

	fmt.Println("\nwant: ", want, "\ngot: ", got)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func NewID(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)

	for _, num := range bookstore.UsedIDs {
		if num == randomNumber {
			randomNumber = rand.Intn(10)

		}
	}

}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()
	book1 := bookstore.Book{Title: "Problem Solving for Programmers", Author: "Amy M Haddad", Discription: "bbbb"}
	book2 := bookstore.Book{Title: "Learn to PS", Author: "Amy M Haddad", Discription: "aaaaaa"}

	bookstore.Books = []bookstore.Book{book1, book2}

	want := bookstore.Books
	got := bookstore.GetBookDetails()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
