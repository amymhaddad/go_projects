package bookstore

import "fmt"

//Book contains details about a book
type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

//GetBooks returns a catalog of books
func GetBooks(catalog []Book) []Book {
	return catalog
}

//GetBookDetails gets the details for a particular book
// func GetBookDetails(catalog map[int]Book, ID int) Book {
// 	return catalog[ID]
// }

//GetBookDetails gets the details for a particular book
func GetBookDetails(catalog []Book, ID string) string {

	for _, book := range catalog {
		if book.ID == ID {
			return fmt.Sprintf("%s, by %s", book.Title, book.Author)
		}
	}
	return ""
}
