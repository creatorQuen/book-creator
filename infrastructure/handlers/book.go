package handlers

import (
	"book-creator/app"
	"book-creator/domain"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type bookHandler struct {
	bookService app.BookService
}

func NewBookHandler(bookService app.BookService) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func (b *bookHandler) CreateBook(ctx echo.Context) error {
	var req domain.Book
	err := ctx.Bind(&req)
	if err != nil {
		log.Println("Request to bind", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := b.bookService.Create(req)
	if err != nil {
		log.Println("Error bookService.Create: ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echo.NewHTTPError(http.StatusCreated, id)
}

func (b *bookHandler) GetBooks(ctx echo.Context) error {
	books, err := b.bookService.GetAll()
	if err != nil {
		log.Println("Error bookService.GetAll: ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, books)
}

func (b *bookHandler) DeleteBook(ctx echo.Context) error {
	bookId := ctx.Param("bookId")
	if bookId == "" {
		log.Println("Error path param is empty.")
		return echo.NewHTTPError(http.StatusBadRequest, "Path parameter is empty.")
	}

	err := b.bookService.DeleteBook(bookId)
	if err != nil {
		log.Println("Invalid request to delete book by book id.")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request to delete book by book id")
	}

	return echo.NewHTTPError(http.StatusOK, "Entity is delete")
}
