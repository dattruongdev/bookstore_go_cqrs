package adapters

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresBorrowRepository struct {
	db *sqlx.DB
}

func NewPostgresBorrowRepository(db *sqlx.DB) domain.BorrowRepository {
	return &PostgresBorrowRepository{db}
}

func (r *PostgresBorrowRepository) FindByBarcode(c context.Context, barcode string) (domain.Borrow, error) {
	var borrow domain.Borrow
	err := r.db.Get(&borrow, "SELECT * FROM borrows WHERE copy_barcode = $1", barcode)
	return borrow, err
}

func (r *PostgresBorrowRepository) FindByUserId(c context.Context, userId uuid.UUID) ([]domain.Borrow, error) {
	var borrows []domain.Borrow
	err := r.db.Select(&borrows, "SELECT * FROM borrows WHERE user_id = $1", userId)
	return borrows, err
}

func (r *PostgresBorrowRepository) CreateBorrow(c context.Context, borrow domain.Borrow) error {
	_, err := r.db.Exec("INSERT INTO borrows (copy_barcode, user_id, borrow_date, return_date) VALUES ($1, $2, $3, $4)",
		borrow.CopyBarcode, borrow.UserID, borrow.BorrowedAt, borrow.ReturnedAt)
	return err
}

func (r *PostgresBorrowRepository) UpdateBorrow(c context.Context, borrow domain.Borrow) error {
	_, err := r.db.Exec("UPDATE borrows SET borrow_date = $1, return_date = $2 WHERE copy_barcode = $3 AND user_id = $4",
		borrow.CopyBarcode, borrow.UserID, borrow.BorrowedAt, borrow.ReturnedAt)
	return err
}

func (r *PostgresBorrowRepository) BeginTransaction() (*sqlx.Tx, error) {
	return r.db.Beginx()
}
