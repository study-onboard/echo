package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/study-onboard/echo/src/vendors/web"
)

var Server *echo.Echo = echo.New()

func init() {
	Server.Debug = true
	Server.HTTPErrorHandler = web.ErrorHandler
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
}
