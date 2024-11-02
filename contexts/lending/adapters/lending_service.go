package adapters

import (
	"context"
	"errors"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/ports"
	"github.com/google/uuid"
)

type lendingService struct {
	copyRespository  domain.CopyRepository
	borrowRepository domain.BorrowRepository
}

func NewLendingService(copyRepo domain.CopyRepository, borrowRepo domain.BorrowRepository) ports.LendingService {
	return &lendingService{
		copyRespository:  copyRepo,
		borrowRepository: borrowRepo,
	}
}

func (s *lendingService) LendCopy(c context.Context, userId uuid.UUID, barcode string) error {
	tx, err := s.borrowRepository.BeginTransaction()
	if err != nil {
		return err
	}

	// Find copy by barcode
	copy, err := s.copyRespository.FindByBarcode(c, barcode)
	if err != nil {
		return err
	}

	// Check if copy is available
	if !copy.Available {
		return errors.New("copy is not available")
	}

	// Create borrow
	borrow := domain.Borrow{
		CopyBarcode: copy.Barcode,
		UserID:      userId,
	}

	// Save borrow
	tx.MustExec("BEGIN;")
	if err := s.borrowRepository.CreateBorrow(c, borrow); err != nil {
		if err := tx.Rollback(); err != nil {
			return err

		}
		return err
	}

	// Update copy status
	copy.MakeUnavailable()

	if err := s.copyRespository.UpdateCopy(c, copy); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	tx.MustExec("COMMIT;")
	return nil
}
