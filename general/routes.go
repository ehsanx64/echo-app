package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func welcomeUser(c echo.Context) error {
	name := c.Param("*")
	return c.String(http.StatusOK, getWelcomeMessage(name))
}

func postWelcomeUser(c echo.Context) error {
	name := c.Param("*")
	return c.String(http.StatusOK, "Posting: "+getWelcomeMessage(name))
}

func AddRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!!\n")
	})

	e.GET("/welcome/*", welcomeUser)
	e.GET("/welcome", welcomeUser)
	e.POST("/welcome/*", postWelcomeUser)
}
