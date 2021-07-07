package bookstore

type Book struct {
	Title       string
	Author      string
	Discription string
	PriceCents  int
	ID          string
}

//Create a slice of type Book
var Books = []Book{}

func GetAllBooks() []Book {
	return Books

}

func GetBookDetails() []Book {
	return Books
}
