package queries

import (
	"context"

	lenddomain "github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
)

type FindCopyByIsbnHandler struct {
	copyRepository lenddomain.CopyRepository
}

func NewFindCopyByIsbnHandler(copyRepository lenddomain.CopyRepository) FindCopyByIsbnHandler {
	return FindCopyByIsbnHandler{copyRepository: copyRepository}
}

func (h *FindCopyByIsbnHandler) Handle(c context.Context, isbn string) ([]lenddomain.Copy, error) {
	return h.copyRepository.FindByBookIsbn(c, isbn)
}
