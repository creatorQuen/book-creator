package app

import (
	"book-creator/domain"
	"book-creator/infrastructure/repository"
	"errors"
	"strconv"
)

type bookService struct {
	repoBook repository.BookRepository
}

func NewBookService(repoBook repository.BookRepository) *bookService {
	return &bookService{repoBook: repoBook}
}

func (b *bookService) Create(book domain.Book) (index int, err error) {
	return b.repoBook.CreateBook(book)
}

func (s *bookService) GetAll() ([]domain.Book, error) {
	return s.repoBook.GetAll()
}

func (s *bookService) DeleteBook(param string) error {
	book, err := s.repoBook.ById(param)
	if err != nil {
		return err
	}

	inputID, _ := strconv.Atoi(param)
	if inputID != book.Id {
		return errors.New("input id and database book id is not equal")
	}
	return s.repoBook.DeleteBook(param)
}
