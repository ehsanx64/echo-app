package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
}

const endpoint = "/users"

func NewUserController(e *echo.Echo) *UserController {
	handler := &UserController{}

	e.GET(endpoint, handler.Index)
	e.GET("/users/:id", handler.Get)
	e.GET("/users/show", handler.View)
	e.POST("/users/save", handler.Create)

	return handler
}

func (uh *UserController) Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Users List\n")
}

func (uh *UserController) Get(ctx echo.Context) error {
	userId := ctx.Param("id")
	return ctx.String(http.StatusOK, "Getting user ID: "+userId+"\n")
}

func (uh *UserController) View(ctx echo.Context) error {
	team := ctx.QueryParam("team")
	member := ctx.QueryParam("member")
	return ctx.String(http.StatusOK, fmt.Sprintf("team: %s, member: %s\n", team, member))
}

func (uh *UserController) Create(ctx echo.Context) error {
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	return ctx.String(http.StatusOK, fmt.Sprintf("name: %s, email: %s\n", name, email))
}
