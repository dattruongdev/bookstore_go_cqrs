package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BorrowRepository interface {
	FindByBarcode(c context.Context, barcode string) (Borrow, error)
	FindByUserId(c context.Context, userId uuid.UUID) ([]Borrow, error)
	CreateBorrow(c context.Context, borrow Borrow) error
	UpdateBorrow(c context.Context, borrow Borrow) error
	BeginTransaction() (*sqlx.Tx, error)
}
