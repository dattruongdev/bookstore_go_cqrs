package domain

import (
	"context"
)

type CopyRepository interface {
	FindByBookIsbn(c context.Context, isbn string) ([]Copy, error)
	FindByBarcode(c context.Context, barcode string) (Copy, error)
	FindAvailableCopies(c context.Context, isbn string) ([]Copy, error)
	FindFirstAvailableCopy(c context.Context, isbn string) (Copy, error)
	CreateCopy(c context.Context, copy Copy) error
	UpdateCopy(c context.Context, copy Copy) error
}
