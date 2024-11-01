package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
)

type FindBookByAuthorName struct {
	Name string
}

type FindBookByAuthorNameHandler struct {
	bookRepository domain.BookRepository
}

func NewFindBookByAuthorNameHandler(bookRepository domain.BookRepository) FindBookByAuthorNameHandler {
	return FindBookByAuthorNameHandler{
		bookRepository,
	}
}

func (h *FindBookByAuthorNameHandler) Handle(c context.Context, name string) ([]domain.Book, error) {
	books, err := h.bookRepository.FindByAuthorName(c, name)

	return books, err
}
