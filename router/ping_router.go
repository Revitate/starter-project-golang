package router

import (
	"github.com/labstack/echo/v4"
	"starter-project/controller"
)

func NewPingRouter(e *echo.Echo, c controller.PingService) *echo.Echo {

	e.GET("/ping", c.Ping)

	return e
}