package routes

import (
	"housy/handlers"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func FiltersRoutes(e *echo.Group) {
	repo := repositories.RepositoryFilter(mysql.DB)
	h := handlers.HandlerFilter(repo)

	e.GET("/multifilter", h.MultiFilter)

}
