package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var catalog = bookstore.Catalog{
	1: {ID: 1, Title: "T1", Author: "A1", Copies: 1},
	2: {ID: 2, Title: "T2", Author: "A2", Copies: 1},
	3: {ID: 3, Title: "T3", Author: "A3", Copies: 1},
}

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:           "T1",
		Author:          "A1",
		Copies:          1,
		ID:              1,
		PriceCents:      5,
		DiscountPercent: 500,
	}

}

func TestNetPriceBook(t *testing.T) {
	t.Parallel()

	centsPerDollar := 100

	b := bookstore.Book{
		Title:           "T1",
		Author:          "A1",
		Copies:          1,
		ID:              1,
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	dollarAmt := b.PriceCents / centsPerDollar
	want := dollarAmt * b.DiscountPercent
	got := b.NetPriceBook()

	if want != got {
		t.Errorf("price %d, after %d%% discount want: net %d, got: %d", dollarAmt, b.DiscountPercent, want, got)
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	want := []bookstore.Book{
		{ID: 1, Title: "T1", Author: "A1", Copies: 1},
		{ID: 2, Title: "T2", Author: "A2", Copies: 1},
		{ID: 3, Title: "T3", Author: "A3", Copies: 1},
	}

	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBook(t *testing.T) {
	t.Parallel()

	ID := 1
	want := bookstore.Book{ID: 1, Title: "T1", Author: "A1", Copies: 1}
	got, err := catalog.GetBook(ID)

	//Return an error IF the err is NOT nil
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestBookDoesNotExist(t *testing.T) {
	t.Parallel()

	ID := 44
	_, err := catalog.GetBook(ID)

	if err == nil {
		t.Fatal(err)
	}

}

func TestSalePrice(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:      "New 1",
		PriceCents: 100,
	}

	want := .50
	got := b.SalePrice()

	if want != got {
		t.Errorf("want: %f, got: %f", want, got)
	}

}

func TestGetBookDetails(t *testing.T) {
	t.Parallel()

	id := 1
	want := "T1 by A1"

	got := catalog.GetBookDetails(id)

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestSetTitle(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "T1",
	}

	b.SetTitle("New Book")
	want := "New Book"
	got := b.Title

	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}

}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{Title: "T5", PriceCents: 4000}
	want := 3000
	err := b.SetPriceCents(want)

	//First check if there's an error. There will be an error if err is not nil. If that's the case, then terminate the program and pass in the err (the error message)
	if err != nil {
		t.Fatal(err)
	}

	got := b.PriceCents
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}

}

//The test fails UNLESS I get an error
func TestInvalidSetPriceCents(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{Title: "T5", PriceCents: 100}
	err := b.SetPriceCents(-100)
	if err == nil {
		t.Fatal(err)
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{Title: "T6"}
	want := "Autobiography"
	err := b.SetCategory("Autobiography")

	if err != nil {
		t.Fatal(err)
	}

	got := b.Category()
	if want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}

func TestInvalidSetCategory(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{Title: "T6"}
	err := b.SetCategory("Fiction")

	if err == nil {
		t.Fatal(err)
	}

}

func TestAddBook(t *testing.T) {
	t.Parallel()

	c := bookstore.Catalog{}
	book := bookstore.Book{ID: 4, Title: "T4", Author: "A4", Copies: 1}
	c.AddBook(book)

	want := []bookstore.Book{book}
	got := c.GetAllBooks()

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}
