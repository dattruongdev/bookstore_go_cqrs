package domain

import (
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindById(c context.Context, userId uuid.UUID) (User, error)
	FindByEmail(c context.Context, email string) (User, error)
	FindByUsername(c context.Context, username string) (User, error)
	CreateUser(c context.Context, user User) error
}
