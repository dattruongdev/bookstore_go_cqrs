package route

import (
	"github.com/dattruongdev/bookstore_cqrs/app"
	"github.com/dattruongdev/bookstore_cqrs/rest/catalog"
	"github.com/dattruongdev/bookstore_cqrs/rest/lending"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, db *sqlx.DB, app *app.Application) {
	v1 := e.Group("/api/v1")

	NewCatalogRoute(v1, catalog.NewCatalogRouteHandler(app))
	NewLendingRoute(v1, lending.NewLendingRouteHandler(app))
}
