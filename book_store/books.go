/*
Practice creating a bookstore
*/

package main

import "fmt"

type Book struct {
	Edition         int
	Series          string
	Price           int
	DiscountPercent int
	Pages           string
	Title           string
	Authors         []string
	Featured        bool
}

const centsPerDollar = 100

func netPrice(price, discount int) int {
	dollars := price / centsPerDollar
	discountValue := float64(discount) * .10
	newPrice := (float64(dollars) - discountValue) * 100

	return int(newPrice)
}

func bookAuthors(book *Book) {
	book.Authors = append(book.Authors, "Sam")
	book.Authors = append(book.Authors, "Tim")
}

func allBooks() []Book {
	books := []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express", Featured: false}, {Title: "One Hundred Years of Good Company", Authors: []string{"amy"}, Featured: true},
	}
	return books
}

func addBook(book *Book) []Book {
	totalBooks := allBooks()
	totalBooks = append(totalBooks, *book)
	return totalBooks
}

func featuredBooks(allBooks []Book) []string {
	var allFeaturedBooks []string

	for _, book := range allBooks {
		if book.Featured == true {
			allFeaturedBooks = append(allFeaturedBooks, book.Title)
		}
	}

	return allFeaturedBooks
}

func main() {
	book1 := Book{
		Title:           "Computer systems",
		Edition:         2,
		Series:          "For the love of go",
		Price:           400,
		DiscountPercent: 2,
		Featured:        true,
	}
	book1.Authors = append(book1.Authors, "SS")

	bookAuthors(&book1)
	total := addBook(&book1)

	featured := featuredBooks(total)
	fmt.Println(featured)
}
