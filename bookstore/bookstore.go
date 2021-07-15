package main

import (
	"fmt"
	"math/rand"
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

//keep outside so bookIds isn't
var bookIds = make(map[int]bool)

// Books creates a slice of type Book
var Books = []Book{}

//GetAllBooks returns a slice of Books
func GetAllBooks() []Book {
	return Books
}

//GetBookDetails returns a slice of Books
func GetBookDetails() []Book {
	return Books
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

func genIds() {
	results := map[string]bool{}
	fmt.Println("r", results)

	for i := 0; i < 10; i++ {
		id := NewID()
		_, ok := results[id]
		if ok {
			fmt.Println("err")
		} else {
			results[id] = true
		}
	}
	fmt.Println(results)
}

func main() {
	genIds()
}
