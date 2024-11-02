package queries

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/user/domain"
)

type FindById struct {
	Id int
}

type FindByIdHandler struct {
	userRepository domain.UserRepository
}

func NewFindByIdHandler(userRepository domain.UserRepository) *FindByIdHandler {
	return &FindByIdHandler{
		userRepository,
	}
}

func (h *FindByIdHandler) Handle(ctx context.Context, query FindById) (domain.User, error) {
	return h.userRepository.FindById(query.Id)
}
