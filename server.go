package main

import (
	"echo-app/general"
	"echo-app/user"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	general.AddRoutes(e)
	user.AddRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
