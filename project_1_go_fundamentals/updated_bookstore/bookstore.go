package bookstore

import (
	"fmt"
)

//Book represents a book in the bookstore
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

//Catalog contains all books
type Catalog map[int]Book

const centsPerDollar = 100
const saleDiscount float64 = .50

//NetPriceBook determines the net price of a book
func (b Book) NetPriceBook() int {
	dollarAmt := b.PriceCents / centsPerDollar
	return dollarAmt * b.DiscountPercent
}

//GetAllBooks returns a slice of books
func (c Catalog) GetAllBooks() []Book {
	allBooks := []Book{}

	for _, book := range c {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

//GetBook returns a single book based on an ID
func (c Catalog) GetBook(ID int) (Book, error) {

	book, found := c[ID]
	if !found {
		//use ErrorF to print a customized error
		return Book{}, fmt.Errorf("book ID %d doesn't exist", ID)
	}
	return book, nil
}

//SalePrice returns the sale price of a book
func (b Book) SalePrice() float64 {
	discountPrice := float64(b.PriceCents) * saleDiscount
	return discountPrice / centsPerDollar
}

//GetBookDetails gets the details for a particular book based on the book ID
func (c Catalog) GetBookDetails(id int) string {
	book, found := c[id]

	if found {
		return fmt.Sprintf("%s by %s", book.Title, book.Author)
	}
	return "Book not found"

}

//SetTitle sets the title of a book
func (b *Book) SetTitle(title string) {
	b.Title = title
}

//SetPriceCents sets the price of a book
func (b *Book) SetPriceCents(value int) {
	b.PriceCents = value
}
