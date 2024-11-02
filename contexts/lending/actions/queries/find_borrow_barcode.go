package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
)

type FindBorrowByBarcode struct {
	Barcode string
}

type FindBorrowByBarcodeHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewFindBorrowByBarcodeHandler(borrowRepository domain.BorrowRepository) *FindBorrowByBarcodeHandler {
	return &FindBorrowByBarcodeHandler{borrowRepository}
}

func (h *FindBorrowByBarcodeHandler) Handle(c context.Context, q FindBorrowByBarcode) (domain.Borrow, error) {
	return h.borrowRepository.FindByBarcode(c, q.Barcode)
}
