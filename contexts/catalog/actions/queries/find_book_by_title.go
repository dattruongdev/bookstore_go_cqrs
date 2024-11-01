package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
)

type FindBookByTitle struct {
	Title string
}

type FindBookByTitleHandler struct {
	bookRepository domain.BookRepository
}

func NewFindBookByTitleHandler(bookRepository domain.BookRepository) FindBookByTitleHandler {
	return FindBookByTitleHandler{
		bookRepository,
	}
}

func (h *FindBookByTitleHandler) Handle(c context.Context, title string) ([]domain.Book, error) {
	books, err := h.bookRepository.FindByTitle(c, title)

	return books, err
}
