package application

import (
	"echo-app/internal/blog"

	"github.com/labstack/echo/v4"
)

func NewBlogApplication(e *echo.Echo) {
	blog.NewBlogController(e)
}
