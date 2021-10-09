package bookstore

import (
	"errors"
	"fmt"
)

//Book contains details about a book
type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

type Catalog map[int]Book

//GetBooks returns a catalog of books
//NOW that I have a Catalog type, I no longer need this long param declaration: func GetBooks(catalog map[int]Book) []Book {
func GetBooks(catalog Catalog) []Book {
	var allBooks []Book
	for _, book := range catalog {
		allBooks = append(allBooks, book)
	}
	return allBooks
}

//GetBookDetails gets the details for a particular book
func GetBookDetails(catalog Catalog, ID int) (string, error) {

	book, found := catalog[ID]
	if !found {
		return "", errors.New("book isn't found")
	}

	return fmt.Sprintf("%s, by %s", book.Title, book.Author), nil
}

//GetBookDetails gets the details for a particular book
// func GetBookDetails(catalog []Book, ID string) (string, error) {
// 	for _, book := range catalog {
// 		if book.ID == ID {
// 			//Sprintf() returns a string instead of printing to terminal
// 			return fmt.Sprintf("%s, by %s", book.Title, book.Author), nil
// 		}
// 	}
// 	return "", errors.New("book isn't found")
// }
