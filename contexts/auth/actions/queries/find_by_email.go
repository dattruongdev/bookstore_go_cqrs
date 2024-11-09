package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type FindByEmail struct {
	Email string
}

type FindByEmailHandler struct {
	userRepository domain.UserRepository
}

func NewFindByEmailHandler(userRepository domain.UserRepository) FindByEmailHandler {
	return FindByEmailHandler{
		userRepository,
	}
}

func (h *FindByEmailHandler) Handle(ctx context.Context, query FindByEmail) (domain.User, error) {
	user, err := h.userRepository.FindByEmail(ctx, query.Email)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		return domain.User{}, errors.NewNotFoundError(slugerr.Error(), slugerr.Slug())
	}

	return user, nil
}
