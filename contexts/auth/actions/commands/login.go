package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/dattruongdev/bookstore_cqrs/errors"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginHandler struct {
	userRepository domain.UserRepository
}

func NewLoginHandler(userRepo domain.UserRepository) LoginHandler {
	return LoginHandler{
		userRepo,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, cmd Login) error {
	user, err := h.userRepository.FindByUsername(ctx, cmd.Username)
	if err != nil {
		return errors.NewNotFoundError(err.Error(), "user-not-found")
	}

	if user.Password != cmd.Password {
		return errors.NewSlugError("invalid password", "invalid-password", 400)
	}

	// TODO: generate token

	return nil
}
