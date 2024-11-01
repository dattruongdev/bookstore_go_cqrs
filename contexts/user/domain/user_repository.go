package domain

type UserRepository interface {
	FindById(userId int) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	CreateUser(user User) error
}
