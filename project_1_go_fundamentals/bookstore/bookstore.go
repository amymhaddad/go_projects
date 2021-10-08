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
	ID     int
}

//Buy reduces the copies of the book left
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no more copies left")
	}
	b.Copies--
	return b, nil
}

//GetAllBooks returns all books in the catalog
func GetAllBooks(catalog map[int]Book) []Book {
	 b := []Book{}
	 for _, book := range catalog {
		 b = append(b, book)
	 }
	return b
}

//GetBook returns a single book
func GetBook(catalog map[int]Book, bookID int) (Book, error) {
	book, found := catalog[bookID]

	if !found {
		return Book{}, errors.New("book doesn't exist")
	}
	return book, nil

}
