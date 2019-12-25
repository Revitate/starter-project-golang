package router

import (
	"github.com/labstack/echo/v4"
	"starter-project/controller"
)

func NewUserRouter(e *echo.Echo, c controller.UserController) *echo.Echo {

	e.GET("/user", c.GetUser)
	e.POST("/user", c.CreateUser)

	return e
}