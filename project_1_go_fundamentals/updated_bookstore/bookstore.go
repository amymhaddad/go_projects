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

const saleDiscount int = .50

//By specifying your code w/methods you say that the code operates on some PARTICULAR data  (ie, not any old param but a type Book )
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
	discountPrice := b.PriceCents * saleDiscount
	//	x := discountPrice / centsPerDollar
	fmt.Println("x", discountPrice)
	return 2.50
	//	return math.Round(discountPrice / centsPerDollar)
}
