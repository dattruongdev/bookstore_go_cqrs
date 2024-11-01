package domain

import (
	"context"
)

type BookRepository interface {
	FindById(c context.Context, bookId Isbn) (Book, error)
	FindByTitle(c context.Context, bookTitle string) ([]Book, error)
	FindByAuthorName(c context.Context, authorName string) ([]Book, error)
	AddBook(c context.Context, book Book) error
}
