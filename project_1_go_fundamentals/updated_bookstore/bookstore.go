package bookstore

//Book represents a book in the bookstore
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
}

const centsPerDollar = 100

//NetPriceBook determines the net price of a book
func (b Book) NetPriceBook() int {
	//func NetPriceBook(b Book) int {
	dollarAmt := b.PriceCents / centsPerDollar
	return dollarAmt * b.DiscountPercent
}

//GetAllBooks returns a slice of books
func GetAllBooks(catalog map[int]Book) []Book {
	allBooks := []Book{}

	for _, book := range catalog {
		allBooks = append(allBooks, book)
	}
	return allBooks
}
