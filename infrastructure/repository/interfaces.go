package repository

import "book-creator/domain"

type BookRepository interface {
	CreateBook(domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	DeleteBook(string) error
	ById(string) (*domain.Book, error)
}
