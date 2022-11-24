package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/study-onboard/echo/src/vendors/web"
)

var Server *echo.Echo = echo.New()

func init() {
	Server.Debug = true
	Server.HTTPErrorHandler = web.ErrorHandler
	Server.Validator = &web.RequestValidator{
		Provider: validator.New(),
	}
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
}
