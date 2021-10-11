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
	Price  int
}

//Catalog contains a catalog of type Book
type Catalog map[int]Book

//GetBooks returns a catalog of books
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

//BuyBook returns a boolean value to determine whether a user can buy a book
func BuyBook(catalog Catalog, price int, ID int) bool {
	book, found := catalog[ID]

	if found && book.Copies > 0 && price >= book.Price {
		book.Copies--
		return true
	}

	return false
}
