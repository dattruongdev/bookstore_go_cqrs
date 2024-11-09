package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
)

type FindAvailableCopies struct {
	Isbn string
}

type FindAvailableCopiesHandler struct {
	copyRepository domain.CopyRepository
}

func NewFindAvailableCopiesHandler(copyRepository domain.CopyRepository) FindAvailableCopiesHandler {
	return FindAvailableCopiesHandler{copyRepository: copyRepository}
}

func (h *FindAvailableCopiesHandler) Handle(c context.Context, isbn string) ([]domain.Copy, error) {
	return h.copyRepository.FindAvailableCopies(c, isbn)
}
