package queries

import (
	"context"

	catalogdomain "github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
	lenddomain "github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
)

type FindCopyByIsbn struct {
	Isbn catalogdomain.Isbn `json:"isbn"`
}

type FindCopyByIsbnHandler struct {
	copyRepository lenddomain.CopyRepository
}

func NewFindCopyByIsbnHandler(copyRepository lenddomain.CopyRepository) *FindCopyByIsbnHandler {
	return &FindCopyByIsbnHandler{copyRepository: copyRepository}
}

func (h *FindCopyByIsbnHandler) Handle(c context.Context, query FindCopyByIsbn) ([]lenddomain.Copy, error) {
	return h.copyRepository.FindByBookIsbn(c, query.Isbn.Value)
}
