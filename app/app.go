package app

import (
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/actions/commands"
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/actions/queries"
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/adapters"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(db *sqlx.DB) Application {
	bookRepository := adapters.NewPostgresBookRepository(db)

	commands := Commands{
		AddBookToCatalog: commands.NewAddBookToCatalogHandler(bookRepository),
	}
	queries := Queries{
		FindBookById:         queries.NewFindBookByIdHandler(bookRepository),
		FindBookByAuthorName: queries.NewFindBookByAuthorNameHandler(bookRepository),
		FindBookByTitle:      queries.NewFindBookByTitleHandler(bookRepository),
	}

	return Application{
		Commands: commands,
		Queries:  queries,
	}
}

type Commands struct {
	AddBookToCatalog commands.AddBookToCatalogHandler
}

type Queries struct {
	FindBookById         queries.FindBookByIdHandler
	FindBookByAuthorName queries.FindBookByAuthorNameHandler
	FindBookByTitle      queries.FindBookByTitleHandler
}
