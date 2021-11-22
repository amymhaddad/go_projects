package bookstore

import (
	"errors"
	"fmt"
)

//Book represents a book in the bookstore
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	discountPercent int
	category        string
}

//Catalog contains all books
type Catalog map[int]Book

const centsPerDollar = 100
const saleDiscount float64 = .50

//This method validates IF a category is valid or not, so it returns an error (or nil). Use unexported field to control what is valid as a category
//SetCategory sets the category of a book
func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		//OR use fmt.Errorf("invalid category %q", c) to use verbs to specify message
		return errors.New("invalid category")
	}
	b.category = category
	return nil
}

//Decide if method is ptr or value by asking: is the point of the method to modify the type (ie, Book)? Here, the answer is no
//I just want to get the value of it
//Category returns the book category
func (b Book) Category() string {
	//I can access the catgory field
	return b.category
}

//NetPriceBook determines the net price of a book
func (b Book) NetPriceBook() int {
	dollarAmt := b.PriceCents / centsPerDollar
	return dollarAmt * b.discountPercent
}

//AddBook adds a book to the catalog
func (c *Catalog) AddBook(b Book) {
	//Can imlicitly dereference a struct -- but dealing w/map. So I need to explicitly deference it.
	//This syntax says: (*c) -- get me the thing that catelog points to (the map). Then, set a key to it catalog[book.ID] = b
	(*c)[b.ID] = b

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
func (b *Book) SetPriceCents(price int) error {
	if price <= 0 {
		return errors.New("invalid price")
	}

	b.PriceCents = price
	return nil
}

//SalePrice returns the sale price of a book
func (b Book) SalePrice() float64 {
	discountPrice := float64(b.PriceCents) * saleDiscount
	return discountPrice / centsPerDollar
}

/*
-A user can directly set the DiscountPercent -- its an exportable field
BUT what if I don't want users to do this? I want to add some kind of parameters on the value that they can set. For ex, I don't wnat
negative discount percents.

I could create a method, BUT users aren't mandated to use it. They could use it OR they could continue to use the struct and set the value that they'd like.

The solution is accessor methods w/unexported values

*/

//SetDiscountPercent sets a discount percent on a book
func (b *Book) SetDiscountPercent(percent int) error {
	if percent <= 0 || percent > 100 {
		return fmt.Errorf("%q is an invalid percent", percent)
	}
	b.discountPercent = percent
	return nil
}

//Discount returns the discount set on a book
func (b Book) Discount() int {
	return b.discountPercent
}
