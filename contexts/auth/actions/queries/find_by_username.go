package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type FindByUsername struct {
	Username string
}

type FindByUsernameHandler struct {
	userRepository domain.UserRepository
}

func NewFindByUsernameHandler(userRepository domain.UserRepository) FindByUsernameHandler {
	return FindByUsernameHandler{
		userRepository,
	}
}

func (h *FindByUsernameHandler) Handle(ctx context.Context, q FindByUsername) (domain.User, error) {
	user, err := h.userRepository.FindByUsername(ctx, q.Username)

	if err != nil {
		return domain.User{}, errors.NewNotFoundError(err.Error(), "user-not-found")
	}

	return user, nil
}
