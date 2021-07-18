package bookstore

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//Book contains elements of a book
type Book struct {
	Title       string
	Author      string
	Description string
	PriceCents  int
	ID          string
}

// Books creates a slice of type Book
var Books = []Book{}

//AddBook adds a book to the Books slice
func AddBook(id, title, author, description string) {
	book := Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Description: description,
	}

	Books = append(Books, book)
}

//GetAllBooks returns a slice of Books
func GetAllBooks() []Book {
	return Books
}

//GetBookDetails returns a string of book details
func GetBookDetails(bookID string) string {
	allBooks := GetAllBooks()

	for _, book := range allBooks {
		if book.ID == bookID {
			return fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nPriceCents: %s\nID: %s", book.Title, book.Author, book.Description, strconv.Itoa(book.PriceCents), book.ID)
		}

	}
	return "Book is not found"

}

//GetAllByAuthor returns all of the books written by an author
func GetAllByAuthor(name string) string {
	allBooks := GetAllBooks()

	var booksWritten []string
	for _, book := range allBooks {
		if book.Author == name {
			booksWritten = append(booksWritten, book.Title)
		}
	}

	return strings.Join(booksWritten, "\n")

}

//NewID returns a book id
func NewID() string {
	rand.Seed(time.Now().UnixNano())

	id := make([]byte, 2)
	numIndex := rand.Intn(9)
	letterIndex := rand.Intn(25)

	num := '0' + numIndex
	letter := letterIndex + 'a'
	fmt.Println(num, letter)

	id[0] = byte(num)
	id[1] = byte(letter)
	return string(id)

}
