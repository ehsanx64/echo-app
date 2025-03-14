package application

import (
	"echo-app/internal/user"

	"github.com/labstack/echo/v4"
)

func NewUserApplication(e *echo.Echo) {
	user.NewUserController(e)
}
