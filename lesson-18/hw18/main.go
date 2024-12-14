package main

import (
	"fmt"
	"hw18/data"
	"hw18/errors"
	"hw18/postgres"
	"hw18/repository"
)

func main() {
	db, err := postgres.InitDB()
	errors.HandleError(err)
	defer db.Close()
	err = db.Ping()
	errors.HandleError(err)
	fmt.Println("Connected to DB!")

	bookRepository := repository.CreateBookRepository(db)

	// 1. INSERT one book
	e := bookRepository.CreateBook(data.Books[0])
	errors.HandleError(e)

	// 2. INSERT many books
	err = bookRepository.CreateManyBooks(data.Books[1:])
	errors.HandleError(err)

	// 3. SELECT one book by ID
	b, err := bookRepository.GetBook(1)
	errors.HandleError(err)
	fmt.Printf("ID: %d\nTitle: %s\nAuthor: %s\nPublisher: %s\nISBN: %s\nCategory: %s\n", b.Id, b.Title, b.Author, b.Publisher, b.Isbn, b.Category)

	// 4. SELECT all books
	books, err := bookRepository.GetAllBooks()
	errors.HandleError(err)
	fmt.Println("All books: ", books)

	// 5. UPDATE one book
	updatedBook := data.Books[2]
	updatedBook.Isbn = "9780451524935"
	updatedBook.Publisher = "Secker & Warburg"
	err = bookRepository.UpdateBook(updatedBook)
	errors.HandleError(err)

	// 6. DELETE one book
	err = bookRepository.DeleteBook(3)
	errors.HandleError(err)
}
