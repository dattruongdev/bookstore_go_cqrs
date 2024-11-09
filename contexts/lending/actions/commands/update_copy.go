package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type UpdateCopy struct {
	Barcode   string
	Available bool
}

type UpdateCopyHandler struct {
	copyRepository domain.CopyRepository
}

func NewUpdateCopyHandler(copyRepository domain.CopyRepository) UpdateCopyHandler {
	return UpdateCopyHandler{
		copyRepository: copyRepository,
	}
}

func (h *UpdateCopyHandler) Handle(c context.Context, cmd UpdateCopy) error {
	copy, err := h.copyRepository.FindByBarcode(c, cmd.Barcode)
	if err != nil {
		return errors.NewNotFoundError(err.Error(), "copy-not-found")
	}

	copy.Available = cmd.Available

	err = h.copyRepository.UpdateCopy(c, copy)
	if err != nil {
		return errors.NewSlugError(err.Error(), "copy-not-updated", 500)
	}

	return nil
}
