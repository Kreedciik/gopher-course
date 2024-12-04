package package1

type Book struct {
	Name      string
	Author    string
	CreatedAt string
}

func GetBooks() []Book {
	book1 := Book{"To Kill a Mockingbird", "Harper Lee", "July 11, 1960"}
	book2 := Book{"1984", "George Orwell", "June 8, 1949"}
	book3 := Book{"The Great Gatsby", "F. Scott Fitzgerald", "April 10, 1925"}
	book4 := Book{"Pride and Prejudice", "Jane Austen", "January 28, 1813"}
	book5 := Book{"The Catcher in the Rye", "J.D. Salinger", "July 16, 1951"}

	return []Book{book1, book2, book3, book4, book5}
}
