package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
)

type AddBookToCatalog struct {
	Isbn      domain.Isbn
	Title     string
	Author    string
	Source    string
	Publisher string
	Edition   string
	Cost      float64
}

type AddBookToCatalogHandler struct {
	bookRepository domain.BookRepository
}

func NewAddBookToCatalogHandler(bookRepository domain.BookRepository) AddBookToCatalogHandler {
	return AddBookToCatalogHandler{
		bookRepository,
	}
}

func (h *AddBookToCatalogHandler) Handle(c context.Context, command *AddBookToCatalog) error {
	book := domain.NewBook(
		command.Isbn,
		command.Title,
		command.Edition,
		command.Author,
		command.Publisher,
		command.Source,
		command.Cost,
	)

	err := h.bookRepository.AddBook(c, *book)

	return err
}
