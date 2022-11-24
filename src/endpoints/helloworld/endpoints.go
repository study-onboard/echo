package helloworld

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	log.Println("Endpoints loaded: helloworld")
	server.Server.GET("/helloworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "Show me the money")
	})
}
