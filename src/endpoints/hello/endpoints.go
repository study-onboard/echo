package helloworld

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	group := server.Server.Group("/hello")
	group.GET("/world", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello World",
		})
	})
}
