package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/study-onboard/echo/src/vendors/web"
)

// server instance
var Server *echo.Echo = echo.New()

func init() {
	// debug mode
	Server.Debug = true

	// custom HTTP error handler
	Server.HTTPErrorHandler = web.ErrorHandler

	// custom form validator
	Server.Validator = &web.RequestValidator{
		Provider: validator.New(),
	}

	// use logger middleware
	Server.Use(middleware.Logger())

	// use recover moddleware
	Server.Use(middleware.Recover())
}
