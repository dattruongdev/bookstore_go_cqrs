package commands

import (
	"context"
	"time"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type UpdateBorrow struct {
	CopyBarcode string    `json:"copy_barcode"`
	ReturnedAt  time.Time `json:"returned_at"`
}

type UpdateBorrowHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewUpdateBorrowHandler(borrowRepository domain.BorrowRepository) UpdateBorrowHandler {
	return UpdateBorrowHandler{borrowRepository}
}

func (h *UpdateBorrowHandler) Handle(c context.Context, cmd UpdateBorrow) error {
	borrow, err := h.borrowRepository.FindByBarcode(c, cmd.CopyBarcode)

	if err != nil {
		return errors.NewNotFoundError("borrow-not-found", "Borrow not found")
	}

	borrow.ReturnedAt = cmd.ReturnedAt

	err = h.borrowRepository.UpdateBorrow(c, borrow)

	if err != nil {
		return err
	}

	return nil
}
