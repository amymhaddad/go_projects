package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = map[int]bookstore.Book{
	1: {ID: 1, Title: "T1", Author: "A1", Copies: 1},
	2: {ID: 2, Title: "T2", Author: "A2", Copies: 1},
	3: {ID: 3, Title: "T3", Author: "A3", Copies: 1},
}

func TestGetBook(t *testing.T) {
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
		//note: writing a literal character: %% (ie, esc the % w/% so Go realizes this isn't a placeholder)
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

	got := bookstore.GetAllBooks(catalog)
	//Maps have random order. So I need to sort the results I get back in order to 
	//properly compare them to what I want.
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
