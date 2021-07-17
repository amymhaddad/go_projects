package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

//GetAllBooks returns a slice of Books
func GetAllBooks() []Book {
	book1 := Book{ID: "Book1", Title: "Problem Solving for Programmers", Author: "Amy Haddad", Description: "bbbb"}
	book2 := Book{Title: "Learn to PS", Author: "Amy Haddad", Description: "aaaaaa"}
	book3 := Book{Title: "Travel 101", Author: "John Smith", Description: "aaaaaa"}

	Books = append(Books, book1, book2, book3)
	return Books
}

//return a string using sprintf -- formats a string via %s
//use book.fieldName -- instead of reflectVal...
//GetBookDetails returns a slice of Books
func GetBookDetails(bookID string) string {
	allBooks := GetAllBooks()

	for _, book := range allBooks {
	
		if book.ID == bookID {
			return fmt.Sprintf("\nTitle: %s \nAuthor: %s \nDescription: %s \nPriceCents: %s \nID: %s", book.Title, book.Author, book.Description, strconv.Itoa(book.PriceCents), book.ID)
			//fmt.Sprintf("%s", book.ID)
		}

	}
	return "amuy"
	
}

func GetAllByAuthor(name string) []string {
	allBooks := GetAllBooks()

	var booksWritten []string
	for _, book := range allBooks {
		if book.Author == name {
			booksWritten = append(booksWritten, book.Title)
		}
	}
	return booksWritten

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

func main() {
	a := GetBookDetails("Book1")
	fmt.Println(a)
	//fmt.Println(GetAllByAuthor("Amy Haddad"))

}
