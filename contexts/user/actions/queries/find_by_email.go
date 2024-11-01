package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/user/domain"
)

type FindByEmail struct {
	Email string
}

type FindByEmailHandler struct {
	userRepository domain.UserRepository
}

func NewFindByEmailHandler(userRepository domain.UserRepository) *FindByEmailHandler {
	return &FindByEmailHandler{
		userRepository,
	}
}

func (h *FindByEmailHandler) Handle(ctx context.Context, query FindByEmail) (domain.User, error) {
	return h.userRepository.FindByEmail(query.Email)
}
