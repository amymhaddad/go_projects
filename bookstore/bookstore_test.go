package bookstore_test

import (
	"bookstore"
	"testing"
)

//ex of the compile only test
func TestBook(t *testing.T) {
	_ = bookstore.Book {
		Title: "Spark Joy",
		Author: "Marie Kondō",
		Description: "A tiny, cheerful Japanese woman explains tidying.", 
		PriceCents: 1199,
	}
}
