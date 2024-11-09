package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type Register struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type RegisterHandler struct {
	userRepository domain.UserRepository
}

func NewCreateUserHandler(userRepository domain.UserRepository) *RegisterHandler {
	return &RegisterHandler{
		userRepository,
	}
}

func (h *RegisterHandler) Handle(ctx context.Context, c Register) error {
	user := domain.User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  c.Password,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Avatar:    c.Avatar,
	}

	err := h.userRepository.CreateUser(ctx, user)
	if err != nil {
		return errors.NewSlugError(err.Error(), "create-user-error", 500)
	}

	return nil
}
