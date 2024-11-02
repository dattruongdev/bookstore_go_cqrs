package adapters

import (
	"context"
	"database/sql"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresCopyRepository struct {
	db *sqlx.DB
}

func NewPostgresCopyRepository(db *sqlx.DB) domain.CopyRepository {
	return &PostgresCopyRepository{db}
}

func (r *PostgresCopyRepository) FindByBookIsbn(c context.Context, isbn string) ([]domain.Copy, error) {
	query := `SELECT * FROM copies WHERE book_isbn=$1`
	rows, err := r.db.Query(query, isbn)

	if err != nil {
		return nil, err
	}

	return scanCopy(rows)

}

func (r *PostgresCopyRepository) FindByBarcode(c context.Context, barcode string) (domain.Copy, error) {
	query := `SELECT * FROM copies WHERE barcode=$1`

	rows, err := r.db.Query(query, barcode)

	if err != nil {
		return domain.Copy{}, err
	}

	copies, err := scanCopy(rows)

	if err != nil {
		return domain.Copy{}, err
	}

	if len(copies) == 0 {
		return domain.Copy{}, sql.ErrNoRows
	}

	return copies[0], nil
}

func (r *PostgresCopyRepository) FindAvailableCopies(c context.Context, isbn string) ([]domain.Copy, error) {
	query := `SELECT * FROM copies WHERE book_isbn=$1 AND available=true`
	rows, err := r.db.Query(query, isbn)

	if err != nil {
		return nil, err
	}

	return scanCopy(rows)
}

func (r *PostgresCopyRepository) CreateCopy(c context.Context, copy domain.Copy) error {
	query := `INSERT INTO copy(book_isbn, barcode) VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(query, copy.BookIsbn.Value, "cp-"+uuid.New().String())

	return err
}

func scanCopy(rows *sql.Rows) ([]domain.Copy, error) {
	var copies []domain.Copy

	for {
		if !rows.Next() {
			break
		}

		var copy domain.Copy

		err := rows.Scan(&copy.BookIsbn, &copy.Barcode, &copy.Available, &copy.CreatedAt, &copy.UpdatedAt)

		if err != nil {
			return nil, err
		}

		copies = append(copies, copy)
	}

	return copies, nil
}

func (r *PostgresCopyRepository) UpdateCopy(c context.Context, copy domain.Copy) error {
	query := `UPDATE copies SET available=$1 WHERE barcode=$2`

	_, err := r.db.Exec(query, copy.Available, copy.Barcode)

	return err
}
