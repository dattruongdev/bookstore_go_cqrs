package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Slug string `json:"slug"`
	Err  string `json:"err"`
}

func InternalError(slug, err string, c *echo.Context) {
	httpResponseWithError(slug, err, c, http.StatusInternalServerError)
}

func BadRequest(slug, err string, c *echo.Context) {
	httpResponseWithError(slug, err, c, http.StatusBadRequest)
}

func Unauthorized(slug, err string, c *echo.Context) {
	httpResponseWithError(slug, err, c, http.StatusUnauthorized)
}

func RespondWithSlugError(slugError SlugError, c *echo.Context) {
	httpResponseWithError(slugError.Slug(), slugError.Error(), c, http.StatusInternalServerError)
}

func httpResponseWithError(slug, err string, c *echo.Context, statusCode int) {
	resperr := (*c).JSON(statusCode, ErrorResponse{
		Slug: slug,
		Err:  err,
	})

	if resperr != nil {
		panic(resperr)
	}
}
