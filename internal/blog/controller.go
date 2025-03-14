package blog

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BlogController struct{}

func NewBlogController(e *echo.Echo) *BlogController {
	handler := &BlogController{}

	e.GET("/", handler.Home)
	e.GET("/about", handler.About)
	e.GET("/blog", handler.Blog)
	e.GET("/blog/:id", handler.View)

	return handler
}

func (uh *BlogController) Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Home")
}

func (uh *BlogController) About(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "About")
}

func (uh *BlogController) Blog(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Blog")
}

func (uh *BlogController) View(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "View")
}
