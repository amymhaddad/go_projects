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
	Title string 
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
	
	// updatedBookPrice := netPrice(book1.Price, book1.DiscountPercent)
	// fmt.Println(updatedBookPrice)

	//Create a slice of TYPE BOOK. Notice how I add to the slice of authors: Authors: []string{"amy"}},
	books := []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express"}, {Title: "One Hundred Years of Good Company", Authors: []string{"amy"}},
		}
	//2 ways to update the title 
	//Option 1: modify AN INDIVIDUAL SLICE element by using it's index
	books[0].Title = "Here and There"
	//Option 2: Modify a FIELD from an individual element  
	books[0] = Book{Title: "The life and times"}
	
	//Append a NEW element to a slice 
	//newBook := Book{Title: "Grapes of Wrath"}
	fmt.Println(books)

}


