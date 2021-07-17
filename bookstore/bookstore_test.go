package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"
	"github.com/google/go-cmp/cmp"
)

//ex of the compile only test
func TestBook(t *testing.T) {
	//note syntax: bookstore.Book {} --> creating an instance of Book struct
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
	//syntax says: in the bookstore module, get the Book struct and create a new intance of type Book
	book1 := bookstore.Book{ID: "Book1", Title: "Problem Solving for Programmers", Author: "Amy Haddad", Description: "bbbb"}
	book2 := bookstore.Book{Title: "Learn to PS", Author: "Amy Haddad", Description: "aaaaaa"}
	book3 := bookstore.Book{Title: "Travel 101", Author: "John Smith", Description: "aaaaaa"}

	want := append(bookstore.Books, book1, book2, book3)
	got := bookstore.GetAllBooks()

	fmt.Println("\nwant: ", want, "\ngot: ", got)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestNewID(t *testing.T) {
	t.Parallel()

	results := map[string]bool{}

	for i := 0; i < 10; i++ {
		id := bookstore.NewID()
		_, ok := results[id]
		if ok {
			t.Errorf("Invalid id: id already exists.")
		} else {
			results[id] = true
		}
	}

}

// func TestAllByAuthor(t *testing.T) {
// 	t.Parallel()

// }


func TestGetBookDetails(t *testing.T) {
	t.Parallel()
		//unclear why this test fails? this is exactly the string I want to get
	want := "Title: Problem Solving for Programmers\nAuthor: Amy Haddad\nDescription: bbbb\nPriceCents: 0\nID: Book1"
	got := bookstore.GetBookDetails("Book1")

	// fmt.Println("\nwant: ", want)
	// fmt.Println("\ngot: ", got)
	// fmt.Println("compare: ", want==got)


	if want != got {
		t.Errorf("No details for this book id")
	}

	
}



// bookstore.GetAllBooks()
// 	allBooks := bookstore.Books
// 	fmt.Println(allBooks)
	
// 	for _, v := range allBooks{
// 		fmt.Printf("%s", v.Title)
// 	}