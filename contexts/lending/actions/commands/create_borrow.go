package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/google/uuid"
)

type CreateBorrow struct {
	CopyBarcode string    `json:"copy_barcode"`
	UserID      uuid.UUID `json:"user_id"`
}

type CreateBorrowHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewCreateBorrowHandler(borrowRepository domain.BorrowRepository) *CreateBorrowHandler {
	return &CreateBorrowHandler{borrowRepository}
}

func (h *CreateBorrowHandler) Handle(c context.Context, cmd CreateBorrow) error {
	borrow := domain.Borrow{
		CopyBarcode: cmd.CopyBarcode,
		UserID:      cmd.UserID,
	}

	return h.borrowRepository.CreateBorrow(c, borrow)
}
