package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/google/uuid"
)

type LendBook struct {
	BookIsbn string    `json:"book_id"`
	UserId   uuid.UUID `json:"user_id"`
}

type LendBookHandler struct {
	copyRepository   domain.CopyRepository
	borrowRepository domain.BorrowRepository
}

func NewLendBookHandler(copyRepo domain.CopyRepository, borrowRepo domain.BorrowRepository) LendBookHandler {
	return LendBookHandler{
		copyRepository:   copyRepo,
		borrowRepository: borrowRepo,
	}
}

func (h *LendBookHandler) Handle(ctx context.Context, cmd *LendBook) error {
	copy, err := h.copyRepository.FindFirstAvailableCopy(ctx, cmd.BookIsbn)
	if err != nil {
		return errors.NewNotFoundError(err.Error(), "book-unavailable")
	}
	tx, err := h.borrowRepository.BeginTransaction()
	if err != nil {
		return errors.NewSlugError(err.Error(), "transaction-error", 500)
	}

	// Create borrow
	borrow := domain.Borrow{
		CopyBarcode: copy.Barcode,
		UserID:      cmd.UserId,
	}

	// Save borrow
	tx.MustExec("BEGIN;")
	if err := h.borrowRepository.CreateBorrow(ctx, borrow); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.NewSlugError(err.Error(), "transaction-error", 500)
		}
		return errors.NewSlugError(err.Error(), "borrow-create-error", 500)
	}

	// Update copy status
	copy.MakeUnavailable()

	if err := h.copyRepository.UpdateCopy(ctx, copy); err != nil {
		if err := tx.Rollback(); err != nil {
			return errors.NewSlugError(err.Error(), "transaction-error", 500)
		}

		return errors.NewSlugError(err.Error(), "copy-update-error", 500)
	}

	tx.MustExec("COMMIT;")
	return nil
}
