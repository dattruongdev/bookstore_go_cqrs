package adapters

import (
	"database/sql"

	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db,
	}
}

func (pr *PostgresUserRepository) FindById(userId uuid.UUID) (domain.User, error) {
	query := `SELECT * FROM users WHERE id=$1`

	rows, err := pr.db.Query(query, userId)

	if err != nil {
		return domain.User{}, err
	}

	user, err := scanUser(rows)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (pr *PostgresUserRepository) FindByEmail(email string) (domain.User, error) {
	query := `SELECT * FROM users WHERE email=$1`

	rows, err := pr.db.Query(query, email)

	if err != nil {
		return domain.User{}, err
	}

	user, err := scanUser(rows)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
func (pr *PostgresUserRepository) FindByUsername(username string) (domain.User, error) {
	query := `SELECT * FROM users WHERE username=$1`

	rows, err := pr.db.Query(query, username)

	if err != nil {
		return domain.User{}, err
	}

	user, err := scanUser(rows)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (pr *PostgresUserRepository) CreateUser(user domain.User) error {
	query := `INSERT INTO users (username, email, password, avatar, role) VALUES ($1, $2, $3, $4, $5)`

	_, err := pr.db.Exec(query, user.Username, user.Email, user.Password, user.Avatar, user.Role)

	if err != nil {
		return err
	}

	return nil
}

func scanUser(rows *sql.Rows) (domain.User, error) {
	if !rows.Next() {
		return domain.User{}, nil
	}
	var user domain.User

	err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.Role)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
