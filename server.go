package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func welcomeUser(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "USER: "+name)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!!")
	})

	e.GET("/welcome/:name", welcomeUser)

	e.Logger.Fatal(e.Start(":1323"))
}
