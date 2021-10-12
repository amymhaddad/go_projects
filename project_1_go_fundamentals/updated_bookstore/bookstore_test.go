package bookstore_test

import (
	"bookstore"
	"testing"
)

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
		PriceCents:      5,
		DiscountPercent: 500,
	}

	dollarAmt := b.PriceCents / centsPerDollar
	want := dollarAmt * b.DiscountPercent
	got := bookstore.NetPriceBook(b)

	if want != got {
		//note: writing a literal character: %% (ie, esc the % w/% so Go realizes this isn't a placeholder)
		t.Errorf("price %d, after %d%% discount want: net %d, got: %d", dollarAmt, b.DiscountPercent, want, got)
	}
}
