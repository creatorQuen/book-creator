package main

import (
	"book-creator/app"
	"book-creator/config"
	"book-creator/infrastructure/handlers"
	"book-creator/infrastructure/repository"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	e := echo.New()

	dbConnection := repository.NewBookRepositoryDb(&cfg)
	bookService := app.NewBookService(dbConnection)
	bookHandler := handlers.NewBookHandler(bookService)

	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books", bookHandler.GetBooks)
	e.DELETE("/books/:bookId", bookHandler.DeleteBook)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.ConfigDataBase.ListenPort), e))
}
