package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type FindCopyByBarcode struct {
	Barcode string
}

type FindCopyByBarcodeHandler struct {
	copyRepository domain.CopyRepository
}

func NewFindCopyByBarcodeHandler(copyRepository domain.CopyRepository) FindCopyByBarcodeHandler {
	return FindCopyByBarcodeHandler{copyRepository: copyRepository}
}

func (h *FindCopyByBarcodeHandler) Handle(c context.Context, barcode string) (domain.Copy, error) {
	copy, err := h.copyRepository.FindByBarcode(c, barcode)

	if err != nil {
		return domain.Copy{}, errors.NewNotFoundError(err.Error(), "not-found")
	}

	return copy, nil
}
