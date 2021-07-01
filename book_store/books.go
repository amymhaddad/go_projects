/*
Add suitable fields to the Book struct to represent a book that is a member of a series. For example, this book is the second in a series called “For the Love of Go”. How would you represent that? Your solution should make it possible to list a book series in order, for example.
*/

package main

import "fmt"

type Book struct {
	Name string
	Edition int
	Series string
	Price int
	DiscountPercent int
	Pages string
	//added a slice to allow multiple authors
	Authors []string
}

const centsPerDollar = 100

func netPrice(price, discount int) int {
	dollars := price / centsPerDollar
	discountValue := float64(discount) * .10
	newPrice := (float64(dollars) - discountValue) * 100

	return int(newPrice)
}

func bookAuthors(book *Book)  {
	book.Authors = append(book.Authors, "Sam")
	book.Authors = append(book.Authors, "Tim")
}

func main() {
	book1 := Book {
		Name: "Data", 
		Edition: 2, 
		Series: "For the love of go",
		Price: 400, 
		DiscountPercent: 2,
	}
	
	bookAuthors(&book1)
	
	updatedBookPrice := netPrice(book1.Price, book1.DiscountPercent)
	fmt.Println(updatedBookPrice)

	fmt.Println(book1.Authors)
	

	
}

