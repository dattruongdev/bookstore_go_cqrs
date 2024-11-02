package commands

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
)

type CreateUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
}

type CreateUserHandler struct {
	userRepository domain.UserRepository
}

func NewCreateUserHandler(userRepository domain.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		userRepository,
	}
}

func (h *CreateUserHandler) Handle(ctx context.Context, c CreateUser) error {
	user := domain.User{
		Username:  c.Username,
		Email:     c.Email,
		Password:  c.Password,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Avatar:    c.Avatar,
		Role:      c.Role,
	}

	return h.userRepository.CreateUser(user)
}
