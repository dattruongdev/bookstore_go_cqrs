package domain

import (
	"time"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
)

type Copy struct {
	BookIsbn  domain.Isbn `json:"book_isbn"`
	Barcode   string      `json:"barcode"`
	Available bool        `json:"available"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (c *Copy) MakeUnavailable() {
	c.Available = false
}

func (c *Copy) MakeAvailable() {
	c.Available = true
}
