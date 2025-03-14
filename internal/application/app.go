package application

import (
	"echo-app/internal/infra"

	"github.com/labstack/echo/v4"
)

const (
	DefaultPort = ":1323"
)

func NewApp() {
	e := echo.New()

	// Load the database stuff
	err := infra.DatabaseLoad()
	if err != nil {
		panic(err)
	}

	NewUserApplication(e)
	NewBlogApplication(e)
	e.Logger.Fatal(e.Start(DefaultPort))
}
