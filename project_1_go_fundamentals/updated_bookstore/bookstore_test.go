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
		DiscountPercent: 50,
	}

}
