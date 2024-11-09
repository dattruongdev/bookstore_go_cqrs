package app

import (
	auth_commands "github.com/dattruongdev/bookstore_cqrs/contexts/auth/actions/commands"
	auth_queries "github.com/dattruongdev/bookstore_cqrs/contexts/auth/actions/queries"
	auth_adapters "github.com/dattruongdev/bookstore_cqrs/contexts/auth/adapters"
	catalog_commands "github.com/dattruongdev/bookstore_cqrs/contexts/catalog/actions/commands"
	catalog_queries "github.com/dattruongdev/bookstore_cqrs/contexts/catalog/actions/queries"
	catalog_adapters "github.com/dattruongdev/bookstore_cqrs/contexts/catalog/adapters"
	lending_commands "github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/commands"
	lending_queries "github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/queries"

	lending_adapters "github.com/dattruongdev/bookstore_cqrs/contexts/lending/adapters"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func NewApplication(db *sqlx.DB) Application {
	bookRepository := catalog_adapters.NewPostgresBookRepository(db)
	borrowRepository := lending_adapters.NewPostgresBorrowRepository(db)
	copyRepository := lending_adapters.NewPostgresCopyRepository(db)
	userRepository := auth_adapters.NewPostgresUserRepository(db)

	commands := Commands{
		Login:            auth_commands.NewLoginHandler(userRepository),
		AddBookToCatalog: catalog_commands.NewAddBookToCatalogHandler(bookRepository),
		LendBook:         lending_commands.NewLendBookHandler(copyRepository, borrowRepository),
		CreateBorrow:     lending_commands.NewCreateBorrowHandler(borrowRepository),
		UpdateBorrow:     lending_commands.NewUpdateBorrowHandler(borrowRepository),
		CreateCopy:       lending_commands.NewCreateCopyHandler(copyRepository),
		UpdateCopy:       lending_commands.NewUpdateCopyHandler(copyRepository),
	}
	queries := Queries{
		FindByEmail:          auth_queries.NewFindByEmailHandler(userRepository),
		FindById:             auth_queries.NewFindByIdHandler(userRepository),
		FindByUsername:       auth_queries.NewFindByUsernameHandler(userRepository),
		FindBookById:         catalog_queries.NewFindBookByIdHandler(bookRepository),
		FindBookByAuthorName: catalog_queries.NewFindBookByAuthorNameHandler(bookRepository),
		FindBookByTitle:      catalog_queries.NewFindBookByTitleHandler(bookRepository),
		FindCopyByBarcode:    lending_queries.NewFindCopyByBarcodeHandler(copyRepository),
		FindCopiesByIsbn:     lending_queries.NewFindCopyByIsbnHandler(copyRepository),
		FindAvailableCopies:  lending_queries.NewFindAvailableCopiesHandler(copyRepository),
		FindBorrowByBarcode:  lending_queries.NewFindBorrowByBarcodeHandler(borrowRepository),
		FindBorrowByUserId:   lending_queries.NewFindBorrowByUserIdHandler(borrowRepository),
	}

	return Application{
		Commands: commands,
		Queries:  queries,
	}
}

type Commands struct {
	Register         auth_commands.RegisterHandler
	Login            auth_commands.LoginHandler
	AddBookToCatalog catalog_commands.AddBookToCatalogHandler
	LendBook         lending_commands.LendBookHandler
	CreateBorrow     lending_commands.CreateBorrowHandler
	CreateCopy       lending_commands.CreateCopyHandler
	UpdateBorrow     lending_commands.UpdateBorrowHandler
	UpdateCopy       lending_commands.UpdateCopyHandler
}

type Queries struct {
	FindByEmail          auth_queries.FindByEmailHandler
	FindById             auth_queries.FindByIdHandler
	FindByUsername       auth_queries.FindByUsernameHandler
	FindBookById         catalog_queries.FindBookByIdHandler
	FindBookByAuthorName catalog_queries.FindBookByAuthorNameHandler
	FindBookByTitle      catalog_queries.FindBookByTitleHandler
	FindCopyByBarcode    lending_queries.FindCopyByBarcodeHandler
	FindCopiesByIsbn     lending_queries.FindCopyByIsbnHandler
	FindAvailableCopies  lending_queries.FindAvailableCopiesHandler
	FindBorrowByBarcode  lending_queries.FindBorrowByBarcodeHandler
	FindBorrowByUserId   lending_queries.FindBorrowByUserIdHandler
}
