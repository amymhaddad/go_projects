package bookstore

//package main

import (
	"errors"
)

//Book defines a book in the bookstore
type Book struct {
	Title  string
	Author string
	Copies int
}

//Buy reduces the copies of the book left
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no more copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
	//  b := []Book{}
	// for _, books := range books {
	// 	b = append(b, books)
	// }
	// return b
} 

// func main() {
// 	want := []Book{
// 		{Title:  "T1"}, 
// 		{Title:  "T2"},
// 	}
// 	x := GetAllBooks(want)
// 	fmt.Println(x)
// }