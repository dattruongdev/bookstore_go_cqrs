package lending

import (
	"github.com/dattruongdev/bookstore_cqrs/contexts/catalog/domain"
	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/commands"
	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/queries"
	"github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/dattruongdev/bookstore_cqrs/rest/response"
	"github.com/labstack/echo/v4"
)

func (h *LendingRouteHandler) CreateCopy(c echo.Context) error {
	isbnStr := c.QueryParam("isbn")

	isbn := domain.Isbn{
		Value: isbnStr,
	}

	createCopy := commands.CreateCopy{
		BookIsbn: isbn,
	}

	err := h.app.Commands.CreateCopy.Handle(c.Request().Context(), createCopy)

	if err != nil {
		errors.InternalError("internal-server-error", err.Error(), &c)
	}

	return c.JSON(200, "Copy created successfully")
}
func (h *LendingRouteHandler) UpdateCopy(c echo.Context) error {
	barcode := c.QueryParam("barcode")

	query, err := h.app.Queries.FindCopyByBarcode.Handle(c.Request().Context(), barcode)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	updateCopy := commands.UpdateCopy{
		Barcode:   barcode,
		Available: !query.Available,
	}

	err = h.app.Commands.UpdateCopy.Handle(c.Request().Context(), updateCopy)

	slugerr, ok = err.(*errors.SlugError)

	if ok {
		errors.InternalError(slugerr.Slug(), slugerr.Error(), &c)
		return slugerr
	}

	return c.JSON(200, "Copy updated successfully")
}
func (h *LendingRouteHandler) FindAvailableCopies(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	findAvailableCopies := queries.FindAvailableCopies{
		Isbn: isbn,
	}

	copies, err := h.app.Queries.FindAvailableCopies.Handle(c.Request().Context(), findAvailableCopies.Isbn)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	return c.JSON(200, response.Response{
		Message: "Available copies found successfully",
		Data:    copies,
	})
}

// find copy by barcode
func (h *LendingRouteHandler) FindCopyByBarcode(c echo.Context) error {
	barcode := c.QueryParam("barcode")

	copy, err := h.app.Queries.FindCopyByBarcode.Handle(c.Request().Context(), barcode)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.NotFound(slugerr.Slug(), slugerr.Error(), &c)
		return slugerr
	}

	return c.JSON(200, response.Response{
		Message: "Copy found successfully",
		Data:    copy,
	})
}

// find copies by isbn
func (h *LendingRouteHandler) FindCopiesByIsbn(c echo.Context) error {
	isbn := c.QueryParam("isbn")

	copies, err := h.app.Queries.FindCopiesByIsbn.Handle(c.Request().Context(), isbn)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
	}

	return c.JSON(200, response.Response{
		Message: "Copies found successfully",
		Data:    copies,
	})
}
