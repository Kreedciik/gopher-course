package repository

import (
	"database/sql"
	"hw18/model"
)

type bookRepository struct {
	Db *sql.DB
}

func CreateBookRepository(db *sql.DB) bookRepository {
	return bookRepository{Db: db}
}

func (b *bookRepository) CreateBook(book model.Book) error {
	_, err := b.Db.Exec(`INSERT INTO books (id, title, author, publisher, isbn, category)
	VALUES ($1, $2, $3, $4, $5, $6)`,
		book.Id, book.Title, book.Author,
		book.Publisher, book.Isbn, book.Category,
	)
	return err
}

func (b *bookRepository) CreateManyBooks(books []model.Book) error {
	for _, book := range books {
		err := b.CreateBook(book)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *bookRepository) GetBook(id int) (model.Book, error) {
	var book model.Book
	err := b.Db.QueryRow(`SELECT * FROM books WHERE id = $1`, id).
		Scan(&book.Id, &book.Title,
			&book.Author, &book.Publisher,
			&book.Isbn, &book.Category,
		)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *bookRepository) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	rows, err := b.Db.Query(`SELECT * FROM books`)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book model.Book
		rows.Scan(&book.Id, &book.Title,
			&book.Author, &book.Publisher,
			&book.Isbn, &book.Category,
		)
		books = append(books, book)
	}

	return books, nil
}

func (b *bookRepository) UpdateBook(book model.Book) error {
	_, err := b.Db.Exec(`UPDATE books SET 
				title = $1, author = $2, 
				publisher = $3, isbn = $4, 
				category = $5
				WHERE id = $6
				`,
		book.Title, book.Author,
		book.Publisher, book.Isbn,
		book.Category, book.Id,
	)
	return err
}

func (b *bookRepository) DeleteBook(id int) error {
	_, err := b.Db.Exec(`DELETE FROM books WHERE id = $1`, id)
	return err
}
