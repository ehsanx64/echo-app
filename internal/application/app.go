package application

import (
	"github.com/labstack/echo/v4"
)

const (
	DefaultPort = ":1323"
)

func NewApp() {
	e := echo.New()
	NewUserApplication(e)
	NewBlogApplication(e)
	e.Logger.Fatal(e.Start(DefaultPort))
}
