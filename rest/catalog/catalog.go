package catalog

import (
	"log"
	"net/http"

	"github.com/dattruongdev/bookstore_cqrs/app"
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/actions/commands"
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
	custom_err "github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/labstack/echo/v4"
)

type CatalogRouteHandler struct {
	app *app.Application
}

func NewCatalogRouteHandler(app *app.Application) *CatalogRouteHandler {
	return &CatalogRouteHandler{
		app: app,
	}
}

func (h *CatalogRouteHandler) AddBook(c echo.Context) error {
	postBook := commands.AddBookToCatalog{}
	err := (&echo.DefaultBinder{}).BindBody(c, &postBook)

	if err != nil {
		custom_err.BadRequest("invalid-request-body", "Invalid request body"+err.Error(), &c)
		return err
	}

	log.Println(postBook)

	err = h.app.Commands.AddBookToCatalog.Handle(c.Request().Context(), &postBook)

	if err != nil {
		custom_err.InternalError("internal-server-error", "Internal server error", &c)
		return err
	}

	return c.JSON(http.StatusCreated, "Book added successfully")
}
func (h *CatalogRouteHandler) FindBookById(c echo.Context) error {
	isbn_param := c.Param("isbn")

	isbn := domain.Isbn{
		Value: isbn_param,
	}

	book, err := h.app.Queries.FindBookById.Handle(c.Request().Context(), isbn)

	log.Println("Book getting from Rest api ", book)
	if err != nil {
		slugerr := custom_err.NewNotFoundError("Book not found"+err.Error(), "book-not-found")
		custom_err.RespondWithSlugError(slugerr, &c)
		return err
	}

	return c.JSON(http.StatusCreated, &book)
}
func (h *CatalogRouteHandler) FindBookByAuthorName(c echo.Context) error {
	name := c.Param("name")
	books, err := h.app.Queries.FindBookByAuthorName.Handle(c.Request().Context(), name)
	if err != nil {
		slugerr := custom_err.NewNotFoundError("Books with this author name not found", "books-not-found")
		custom_err.RespondWithSlugError(slugerr, &c)
		return err
	}

	return c.JSON(http.StatusOK, books)
}
func (h *CatalogRouteHandler) FindBookByTitle(c echo.Context) error {
	title := c.Param("title")
	books, err := h.app.Queries.FindBookByTitle.Handle(c.Request().Context(), title)
	if err != nil {
		slugerr := custom_err.NewNotFoundError("Books with this title not found", "books-not-found")
		custom_err.RespondWithSlugError(slugerr, &c)
		return err
	}

	return c.JSON(http.StatusOK, books)
}
