package repository

import (
	"book-creator/config"
	"book-creator/domain"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type bookRepositoryDB struct {
	db *sql.DB
}

func (b *bookRepositoryDB) CreateBook(book domain.Book) (int, error) {
	var lastInsertID int
	err := b.db.QueryRow("INSERT INTO books(book_id, book_name) VALUES($1, $2) returning id;", book.BookID, book.BookName).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (pr *bookRepositoryDB) GetAll() ([]domain.Book, error) {
	// Get all books from books table that don't have bookID = "1"
	//rows, err := pr.db.Query("SELECT * FROM books where book_id <> $1", "1")
	rows, err := pr.db.Query("SELECT id, book_id, book_name FROM books")
	if err != nil {
		return nil, err
	}

	var books []domain.Book
	for rows.Next() {
		var id int
		var bookID string
		var bookName string

		err = rows.Scan(&id, &bookID, &bookName)
		if err != nil {
			return nil, err
		}

		books = append(books, domain.Book{Id: id, BookID: bookID, BookName: bookName})
	}

	return books, nil
}

func (b *bookRepositoryDB) DeleteBook(bookId string) error {
	_, err := b.db.Exec("DELETE FROM books where id = $1", bookId)
	return err
}

func (b *bookRepositoryDB) ById(bookId string) (*domain.Book, error) {
	query := `SELECT id, book_id, book_name FROM books WHERE id=$1`

	var book domain.Book
	err := b.db.QueryRow(query, bookId).Scan(&book.Id, &book.BookID, &book.BookName)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Not found row")
		} else {
			log.Println("Error while scanning book " + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &book, nil

}

func NewBookRepositoryDb(c *config.Config) *bookRepositoryDB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.ConfigDataBase.Host,
		c.ConfigDataBase.Port,
		c.ConfigDataBase.User,
		c.ConfigDataBase.Password,
		c.ConfigDataBase.NameDataBase,
	)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &bookRepositoryDB{db}
}
