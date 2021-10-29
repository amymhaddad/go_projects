package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	if !cmp.Equal(want, got) {
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
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestBookDoesNotExist(t *testing.T) {
	t.Parallel()

	ID := 4
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
