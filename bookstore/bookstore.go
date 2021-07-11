package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Book blah
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

func GetAllBooks() []Book {
	return Books
}

func GetBookDetails() []Book {
	return Books
}

//ask:
//Bc dealing w/small nums and letters it makes sense to use bytes -- not runes

func NewID()  {
	rand.Seed(time.Now().UnixNano())

	letters := "abcdefghijklmnopqrstuvwxyz"
	var arr [2]byte
	index := rand.Intn(25)

	num := '0' + index
	letter := letters[index] + 'a'

	arr[0] = byte(num)
	arr[1] = byte(letter)
	//return string(arr)
	

}

func main() {
	fmt.Println(NewID)
}
