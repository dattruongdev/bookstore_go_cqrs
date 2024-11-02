package ports

import (
	"context"

	"github.com/google/uuid"
)

type LendingService interface {
	LendCopy(c context.Context, userId uuid.UUID, barcode string) error
}
