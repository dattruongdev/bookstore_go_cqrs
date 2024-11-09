package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/google/uuid"
)

type FindById struct {
	Id uuid.UUID
}

type FindByIdHandler struct {
	userRepository domain.UserRepository
}

func NewFindByIdHandler(userRepository domain.UserRepository) FindByIdHandler {
	return FindByIdHandler{
		userRepository,
	}
}

func (h *FindByIdHandler) Handle(ctx context.Context, query FindById) (domain.User, error) {
	user, err := h.userRepository.FindById(ctx, query.Id)

	if err != nil {
		return domain.User{}, errors.NewNotFoundError(err.Error(), "user-not-found")
	}

	return user, nil
}
