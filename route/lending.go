package route

import (
	"github.com/dattruongdev/bookstore_cqrs/rest/lending"
	"github.com/labstack/echo/v4"
)

func NewLendingRoute(g *echo.Group, handler *lending.LendingRouteHandler) {
	lendingRoute := g.Group("/lending")

	// commands
	lendingRoute.POST("/lend", handler.LendBook)
	lendingRoute.POST("/create-borrow", handler.CreateBorrow)
	lendingRoute.POST("/create-copy", handler.CreateCopy)
	lendingRoute.PUT("/update-borrow", handler.UpdateBorrow)
	lendingRoute.PUT("/update-copy", handler.UpdateCopy)

	// queries
	lendingRoute.GET("/borrow-by-barcode/:barcode", handler.FindBorrowByBarcode)
	lendingRoute.GET("/borrow-by-userid/:userid", handler.FindBorrowByUserId)
	lendingRoute.GET("/copy-by-barcode/:barcode", handler.FindCopyByBarcode)
	lendingRoute.GET("/copy-by-isbn/:isbn", handler.FindCopiesByIsbn)
	lendingRoute.GET("/copy-available/:isbn", handler.FindAvailableCopies)
}
