package controller

import "github.com/labstack/echo/v4"

type pingService struct {}

type PingService interface {
	Ping(ctx echo.Context) error
}

func NewPingService() PingService {
	return &pingService{}
}

func (p *pingService) Ping(ctx echo.Context) error {
	return ctx.String(200, "pong")
}
