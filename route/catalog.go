package route

import (
	"github.com/dattruongdev/bookstore_cqrs/rest/catalog"
	"github.com/labstack/echo/v4"
)

func NewCatalogRoute(group *echo.Group, handler *catalog.CatalogRouteHandler) {
	bookRoute := group.Group("/books")

	bookRoute.GET("/:isbn", handler.FindBookById)
	bookRoute.POST("/add", handler.AddBook)
	bookRoute.GET("/by-author", handler.FindBookByAuthorName)
	bookRoute.GET("/by-title", handler.FindBookByTitle)
}
