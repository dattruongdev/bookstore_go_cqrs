package auth

import (
	"github.com/dattruongdev/bookstore_cqrs/app"
	"github.com/dattruongdev/bookstore_cqrs/contexts/auth/actions/commands"
	"github.com/dattruongdev/bookstore_cqrs/errors"
	"github.com/labstack/echo/v4"
)

type AuthRouteHandler struct {
	app *app.Application
}

func NewAuthRouteHandler(app *app.Application) AuthRouteHandler {
	return AuthRouteHandler{
		app: app,
	}
}

func (h *AuthRouteHandler) Register(c echo.Context) error {
	cmd := commands.Register{}

	err := (&echo.DefaultBinder{}).BindBody(c, &cmd)

	if err != nil {
		errors.BadRequest("invalid-request-body", err.Error(), &c)
		return err
	}

	err = h.app.Commands.Register.Handle(c.Request().Context(), cmd)

	slugerror, ok := err.(*errors.SlugError)

	if ok {
		errors.InternalError(slugerror.Slug(), slugerror.Error(), &c)
		return err
	}

	return c.JSON(201, "User created successfully")
}

func (h *AuthRouteHandler) Login(e echo.Context) error {
	login := commands.Login{}

	err := (&echo.DefaultBinder{}).BindBody(e, &login)

	if err != nil {
		errors.BadRequest("invalid-request-body", err.Error(), &e)
		return err
	}

	err = h.app.Commands.Login.Handle(e.Request().Context(), login)

	slugerror, ok := err.(*errors.SlugError)

	if ok {
		errors.InternalError(slugerror.Slug(), slugerror.Error(), &e)
		return err
	}

	return e.JSON(200, "Login successfully")
}
