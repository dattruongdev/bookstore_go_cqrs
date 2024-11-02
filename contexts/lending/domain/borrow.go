package domain

import (
	"time"

	"github.com/google/uuid"
)

type Borrow struct {
	CopyBarcode string    `json:"copy_barcode"`
	UserID      uuid.UUID `json:"user_id"`
	BorrowedAt  time.Time `json:"borrowed_at"`
	ReturnedAt  time.Time `json:"returned_at"`
}
