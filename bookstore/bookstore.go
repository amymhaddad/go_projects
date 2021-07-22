package main

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
var Books = map[string]Book{}

// BooksByAuthor keeps track of each book an author writes
var BooksByAuthor = make(map[string]map[string]string)

//var BooksByAuthor = make(map[string]Book)

//AddBook adds a book to the Books slice
func AddBook(id, title, author, description string) {
	book := Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Description: description,
	}

	Books[id] = book
}

// AddAuthorBooks adds an author's book, id, and title to the BooksByAuthor map
func AddAuthorBooks(name, id, title string) {

	//Checking to see if the name exists in the map. IF it does,
	//then I want add another id/title to it (i don't want to overwrite what's
	//already there)
	if _,ok := BooksByAuthor[name]; ok {
		fmt.Println(BooksByAuthor
		author := BooksByAuthor[name]
		author[id] = currVal
	}

	idTitle := map[string]string{id: title}
	BooksByAuthor[name] = idTitle
}

func main() {
	AddAuthorBooks("Aamy", "id1", "title")
}

//GetAllBooks returns a map of Books
func GetAllBooks() map[string]Book {
	return Books
}

//GetBookDetails returns a string of book details
func GetBookDetails(bookID string) string {
	allBooks := GetAllBooks()

	book, ok := allBooks[bookID]

	if !ok {
		return "Book is not found"
	}
	return fmt.Sprintf("Title: %s\nAuthor: %s\nDescription: %s\nPriceCents: %s\nID: %s", book.Title, book.Author, book.Description, strconv.Itoa(book.PriceCents), book.ID)

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

	id[0] = byte(num)
	id[1] = byte(letter)
	return string(id)

}

// func main() {
// 	AddBook("Book1", "Problem Solving for Programmers", "Amy Haddad", "bbbb")
// 	fmt.Println(GetBookDetails("Book1"))
// }
