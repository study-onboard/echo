package error

import (
	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	group := server.Server.Group("/error")

	group.GET("/custom", func(c echo.Context) error {
		c.Logger().Info("show me the money")
		return nil
	})
}
