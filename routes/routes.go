package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	UserRoutes(e)
	PropertyRoutes(e)
	TransactionRoutes(e)
	FiltersRoutes(e)
}
