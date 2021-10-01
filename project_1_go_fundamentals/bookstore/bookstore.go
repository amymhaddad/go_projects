package bookstore

import "errors"

//Book defines a book in the bookstore
type Book struct {
	Title  string
	Author string
	Copies int
}

//The function signature is already determined by the test
//Start with a failing test, which means you need to return the zero value of the result type is.
//In this case, it's an empty Book 
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		//conventionally return zero for the data value, since it’s to be ignored (in this case, it's Book{})
		return Book{}, errors.New("no more copies left")
	}
	b.Copies--
	return b, nil
}