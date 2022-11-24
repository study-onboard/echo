package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Server *echo.Echo = echo.New()

func init() {
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
}
