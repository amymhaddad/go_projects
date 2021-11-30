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
	category        Category
}

//Category represents a book category
type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

//Catalog contains all books
type Catalog map[int]Book

const (
	centsPerDollar         = 100
	saleDiscount   float64 = .50
)

//SetCategory sets the category of a book
func (b *Book) SetCategory(category Category) error {

	if !validCategory[category] {
		return errors.New("invalid category")
	}
	b.category = category
	return nil
}


//Category returns the book category
func (b Book) GetCategory() Category {
	return b.category
}

//NetPriceBook determines the net price of a book
func (b Book) NetPriceBook() int {
	dollarAmt := b.PriceCents / centsPerDollar
	return dollarAmt * b.discountPercent
}

//AddBook adds a book to the catalog
func (c *Catalog) AddBook(b Book) {
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

//SetDiscountPercent sets a discount percent on a book
func (b *Book) SetDiscountPercent(percent int) error {
	if invalidDiscount(percent) {
		return fmt.Errorf("%q is an invalid percent", percent)
	}
	b.discountPercent = percent
	return nil
}

//Discount returns the discount set on a book
func (b Book) Discount() int {
	return b.discountPercent
}

func invalidDiscount(percent int) bool {
	return percent <= 0 || percent > 100
}

