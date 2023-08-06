package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/labstack/echo/v4"
)

func PropertyRoutes(e *echo.Group) {
	repo := repositories.MakeRepository(mysql.DB)
	h := handlers.HandlerProperty(repo)

	e.GET("/property/:id", h.GetProperty)
	e.GET("/properties", h.FindProperties)
	e.POST("/property", middleware.Auth(middleware.UploadFile(h.AddProperty)))

}
