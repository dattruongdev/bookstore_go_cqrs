package lending

import (
	"time"

	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/commands"
	"github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/dattruongdev/bookstore_cqrs/rest/response"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *LendingRouteHandler) CreateBorrow(c echo.Context) error {
	id := c.QueryParam("userid")
	barcode := c.QueryParam("barcode")

	userId, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	createBorrow := commands.CreateBorrow{
		CopyBarcode: barcode,
		UserID:      userId,
	}

	err = h.app.Commands.CreateBorrow.Handle(c.Request().Context(), createBorrow)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	return c.JSON(201, "Borrow created successfully")
}

func (h *LendingRouteHandler) UpdateBorrow(c echo.Context) error {
	barcode := c.QueryParam("barcode")

	updateBorrow := commands.UpdateBorrow{
		CopyBarcode: barcode,
		ReturnedAt:  time.Now(),
	}

	err := h.app.Commands.UpdateBorrow.Handle(c.Request().Context(), updateBorrow)
	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	return c.JSON(200, "Borrow record updated successfully")
}

func (h *LendingRouteHandler) FindBorrowByBarcode(c echo.Context) error {
	barcode := c.QueryParam("barcode")

	borrow, err := h.app.Queries.FindBorrowByBarcode.Handle(c.Request().Context(), barcode)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	return c.JSON(200, response.Response{
		Message: "Borrow record found",
		Data:    borrow,
	})
}

func (h *LendingRouteHandler) FindBorrowByUserId(c echo.Context) error {
	id := c.QueryParam("userid")

	userId, err := uuid.Parse(id)

	if err != nil {
		errors.BadRequest("invalid-user-id", "Invalid user id", &c)
		return err
	}

	borrows, err := h.app.Queries.FindBorrowByUserId.Handle(c.Request().Context(), userId)

	slugerr, ok := err.(*errors.SlugError)

	if ok {
		errors.RespondWithSlugError(slugerr, &c)
		return slugerr
	}

	return c.JSON(200, response.Response{
		Message: "Borrows found",
		Data:    borrows,
	})
}
