package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type FindBorrowByBarcode struct {
	Barcode string
}

type FindBorrowByBarcodeHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewFindBorrowByBarcodeHandler(borrowRepository domain.BorrowRepository) FindBorrowByBarcodeHandler {
	return FindBorrowByBarcodeHandler{borrowRepository}
}

func (h *FindBorrowByBarcodeHandler) Handle(c context.Context, barcode string) (domain.Borrow, error) {
	borrow, err := h.borrowRepository.FindByBarcode(c, barcode)

	if err != nil {
		return domain.Borrow{}, errors.NewNotFoundError("borrow-not-found", "Borrow not found")
	}

	return borrow, nil
}
