package app

import "book-creator/domain"

type BookService interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	DeleteBook(param string) error
}
