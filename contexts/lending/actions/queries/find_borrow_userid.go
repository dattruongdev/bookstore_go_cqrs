package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/google/uuid"
)

type FindBorrowByUserId struct {
	UserID uuid.UUID `json:"user_id"`
}

type FindBorrowByUserIdHandler struct {
	borrowRepository domain.BorrowRepository
}

func NewFindBorrowByUserIdHandler(borrowRepository domain.BorrowRepository) *FindBorrowByUserIdHandler {
	return &FindBorrowByUserIdHandler{
		borrowRepository: borrowRepository,
	}
}

func (h *FindBorrowByUserIdHandler) Handle(c context.Context, query FindBorrowByUserId) ([]domain.Borrow, error) {
	return h.borrowRepository.FindByUserId(c, query.UserID)
}
