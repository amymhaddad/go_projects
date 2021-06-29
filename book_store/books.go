package main

import "fmt"

// func author() {
// 	var authorName string
// 	authorName = "Amy"
// 	fmt.Println(authorName)
// }

func bookEdition(editionNum float64) float64 {
	return editionNum
}

func specialDiscount(isDiscount bool, amount float64) float64 {
	if isDiscount {
		return amount
	} else {
		return 0
	}
}

func main() {
	fmt.Println(bookEdition(1))
	// x := author()
	// fmt.Println(x)
	fmt.Println(specialDiscount(true, 15))
}
