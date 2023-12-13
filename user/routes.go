package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!!\n")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		userId := c.Param("id")
		return c.String(http.StatusOK, "Getting user ID: "+userId+"\n")
	})

	e.GET("/users/show", func(c echo.Context) error {
		team := c.QueryParam("team")
		member := c.QueryParam("member")
		return c.String(http.StatusOK, fmt.Sprintf("team: %s, member: %s\n", team, member))
	})
}
