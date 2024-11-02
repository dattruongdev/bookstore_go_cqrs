package domain

import "github.com/google/uuid"

type UserRepository interface {
	FindById(userId uuid.UUID) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	CreateUser(user User) error
}
