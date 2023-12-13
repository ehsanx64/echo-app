package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!!")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		userId := c.Param("id")
		return c.String(http.StatusOK, "Getting user ID: "+userId+"\n")
	})
}
