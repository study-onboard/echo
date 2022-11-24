package error

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
	"github.com/study-onboard/echo/src/vendors/web"
)

func init() {
	group := server.Server.Group("/error")

	group.GET("/standard", func(c echo.Context) error {
		return errors.New("some standard error")
	})

	group.GET("/custom", func(c echo.Context) error {
		return &web.ApiError{
			Code:    1024,
			Message: "some error about access db",
		}
	})
}
