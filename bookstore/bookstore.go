package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//Book contains elements of a book
type Book struct {
	Title       string
	Author      string
	Discription string
	PriceCents  int
	ID          string
}

// Books creates a slice of type Book
var Books = []Book{}

//GetAllBooks returns a slice of Books
func GetAllBooks() []Book {
	book1 := Book{ID: "Book1", Title: "Problem Solving for Programmers", Author: "Amy M Haddad", Discription: "bbbb"}
	book2 := Book{Title: "Learn to PS", Author: "Amy M Haddad", Discription: "aaaaaa"}

	Books = append(Books, book1)
	Books = append(Books, book2)
	return Books
}

//GetBookDetails returns a slice of Books
func GetBookDetails(bookID string){
	allBooks := GetAllBooks()
	for _, book := range allBooks {
		if book.ID == bookID {
			reflectVal := reflect.ValueOf(book)
			values := make([]interface{}, reflectVal.NumField())
			for i := 0; i < reflectVal.NumField(); i++ {
				values[i] = reflectVal.Field(i).Interface()
			}

			for _, value := range values {
				fmt.Println(value)	
			}
		}
	}
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

// func genIds() {
// 	results := map[string]bool{}
//
// 	for i := 0; i < 10; i++ {
// 		id := NewID()
// 		_, ok := results[id]
// 		if ok {
// 			fmt.Println("err")
// 		} else {
// 			results[id] = true
// 		}
// 	}
// }
//
func main() {
	GetBookDetails("Book1")
}
