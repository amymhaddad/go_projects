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

//for map implementation
// var v2Books = map[string]Book{}

// func v2GetAllBooks() string {
// 	v2Books["Book1"] = Book{
// 		ID: "Book1", 
// 		Title: "Problem Solving for Programmers", 
// 		Author: "Amy Haddad", 
// 		Description: "bbbb"}
	
// 	v2Books["Book2"] = Book{
// 		ID: "Book2", Title: "Learn to PS", Author: "Amy Haddad", Description: "aaaaaa",
// 	}
// 	return v2Books

// }

//GetAllBooks returns a slice of Books
func GetAllBooks() []Book {
	book1 := Book{ID: "Book1", Title: "Problem Solving for Programmers", Author: "Amy Haddad", Description: "bbbb"}
	book2 := Book{Title: "Learn to PS", Author: "Amy Haddad", Description: "aaaaaa"}
	book3 := Book{Title: "Travel 101", Author: "John Smith", Description: "aaaaaa"}

	Books = append(Books, book1, book2, book3)
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

// func main() {
// 	GetAllBooks()
// 	allBooks := Books 
	
// 	for _, v := range allBooks{
// 		fmt.Printf("%s", v.Title)
// 	}
// }
