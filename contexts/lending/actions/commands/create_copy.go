package commands

import (
	"context"

	bookdomain "github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
	lenddomain "github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type CreateCopy struct {
	BookIsbn bookdomain.Isbn `json:"book_isbn"`
}

type CreateCopyHandler struct {
	repo lenddomain.CopyRepository
}

func NewCreateCopyHandler(repo lenddomain.CopyRepository) CreateCopyHandler {
	return CreateCopyHandler{repo}
}

func (h *CreateCopyHandler) Handle(c context.Context, cmd CreateCopy) error {
	copy := lenddomain.Copy{
		BookIsbn: cmd.BookIsbn,
	}

	err := h.repo.CreateCopy(c, copy)

	if err != nil {
		return errors.NewSlugError(err.Error(), "create-copy", 500)
	}

	return nil
}
