package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/google/uuid"
)

type FindBorrowByUserIdHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewFindBorrowByUserIdHandler(borrowRepository domain.BorrowRepository) FindBorrowByUserIdHandler {
	return FindBorrowByUserIdHandler{
		borrowRepository: borrowRepository,
	}
}

func (h *FindBorrowByUserIdHandler) Handle(c context.Context, userId uuid.UUID) ([]domain.Borrow, error) {
	return h.borrowRepository.FindByUserId(c, userId)
}
