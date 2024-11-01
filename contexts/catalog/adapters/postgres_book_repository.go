package adapters

import (
	"context"

	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
	"github.com/jmoiron/sqlx"
)

type PostgresBookRepository struct {
	db *sqlx.DB
}

func NewPostgresBookRepository(db *sqlx.DB) *PostgresBookRepository {
	return &PostgresBookRepository{
		db,
	}
}

func (pr *PostgresBookRepository) FindById(c context.Context, bookId domain.Isbn) (domain.Book, error) {
	query := `SELECT * FROM books WHERE isbn=$1`

	row, err := pr.db.Query(query, bookId.Value)
	if err != nil {
		return domain.Book{}, err
	}

	row.Next()
	var isbn string
	var title string
	var edition string
	var author string
	var publisher string
	var source string
	var cost float64

	err = row.Scan(&isbn, &title, &edition, &author, &publisher, &source, &cost)
	if err != nil {
		return domain.Book{}, err
	}

	book := domain.NewBook(
		domain.Isbn{Value: isbn},
		title,
		edition,
		author,
		publisher,
		source,
		cost,
	)

	return *book, nil
}
func (pr *PostgresBookRepository) FindByTitle(c context.Context, bookTitle string) ([]domain.Book, error) {
	query := `SELECT * FROM books WHERE title = $1`

	var books []domain.Book

	err := pr.db.Select(&books, query, bookTitle)
	return books, err
}
func (pr *PostgresBookRepository) FindByAuthorName(c context.Context, authorName string) ([]domain.Book, error) {
	query := `SELECT * FROM books WHERE author = $1`

	var books []domain.Book

	err := pr.db.Select(&books, query, authorName)
	return books, err
}
func (pr *PostgresBookRepository) AddBook(c context.Context, book domain.Book) error {
	query := `INSERT INTO books (isbn, title, edition, author, publisher, source, cost) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := pr.db.Exec(query, book.Isbn, book.Title, book.Edition, book.Author, book.Publisher, book.Source, book.Cost)

	return err
}
