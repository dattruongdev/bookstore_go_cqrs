package lending

import (
	"github.com/dattruongdev/bookstore_cqrs/app"
	"github.com/dattruongdev/bookstore_cqrs/contexts/lending/actions/commands"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LendingRouteHandler struct {
	app *app.Application
}

func NewLendingRouteHandler(app *app.Application) *LendingRouteHandler {
	return &LendingRouteHandler{
		app: app,
	}
}

func (h *LendingRouteHandler) LendBook(c echo.Context) error {

	//bookisbn and userid
	bookisbn := c.QueryParam("bookisbn")
	userID := c.QueryParam("userid")
	id, err := uuid.Parse(userID)

	if err != nil {
		return err
	}

	lendBook := commands.LendBook{
		BookIsbn: bookisbn,
		UserId:   id,
	}

	err = h.app.Commands.LendBook.Handle(c.Request().Context(), &lendBook)

	if err != nil {
		return err
	}

	return c.JSON(200, "Book lent successfully")
}
