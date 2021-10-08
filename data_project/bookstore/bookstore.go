package bookstore

//Book contains details about a book
type Book struct {
	Title  string
	Author string
	Copies int
}

//GetBooks returns a catalog of books
func GetBooks(catalog []Book) []Book {
	return catalog
}
