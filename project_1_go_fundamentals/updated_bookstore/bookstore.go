package bookstore

import "fmt"

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