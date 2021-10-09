package bookstore_test

import (
	"bookstore"
	"sort"

	//"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//Must use var for package level declarations, not the short declaration :=
var catalog = bookstore.Catalog{
	1: {ID: 1, Title: "T1", Author: "A1", Copies: 1},
	2: {ID: 2, Title: "T2", Author: "A2", Copies: 1},
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
		{ID: 1, Title: "T1", Author: "A1", Copies: 1},
		{ID: 2, Title: "T2", Author: "A2", Copies: 1},
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

//Unsure why this test doesn't work when I try an ID I know that won't work?
//Test error values w/separate test?
// func TestGetBookDetails(t *testing.T) {
// 	t.Parallel()

// 	want := "Title 1, by A1"
// 	got, err := bookstore.GetBookDetails(catalog, "3")

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if !cmp.Equal(want, got) {
// 		t.Errorf(cmp.Diff(want, got))
// 	}

// }
//This is the function:
// func GetBookDetails(catalog []Book, ID string) (string, error) {
// 	for _, book := range catalog {
// 		if book.ID == ID {
// 			return fmt.Sprintf("%s, by %s", book.Title, book.Author), nil
// 		}
// 	}
// 	return "", errors.New("book isn't found")
// }
