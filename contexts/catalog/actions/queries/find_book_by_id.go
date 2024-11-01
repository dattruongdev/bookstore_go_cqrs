package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
)

type FindBookById struct {
	Isbn domain.Isbn
}

type FindBookByIdHandler struct {
	bookRepository domain.BookRepository
}

func NewFindBookByIdHandler(bookRepository domain.BookRepository) FindBookByIdHandler {
	return FindBookByIdHandler{
		bookRepository,
	}
}

func (h *FindBookByIdHandler) Handle(c context.Context, isbn domain.Isbn) (domain.Book, error) {
	book, err := h.bookRepository.FindById(c, isbn)
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}
