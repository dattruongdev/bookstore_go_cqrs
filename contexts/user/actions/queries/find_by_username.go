package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/user/domain"
)

type FindByUsername struct {
	Username string
}

type FindByUsernameHandler struct {
	userRepository domain.UserRepository
}

func NewFindByUsernameHandler(userRepository domain.UserRepository) *FindByUsernameHandler {
	return &FindByUsernameHandler{
		userRepository,
	}
}

func (h *FindByUsernameHandler) Handle(ctx context.Context, q FindByUsername) (domain.User, error) {
	return h.userRepository.FindByUsername(q.Username)
}
