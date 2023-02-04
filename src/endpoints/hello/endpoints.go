package helloworld

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	// register http handler and handler group
	group := server.Server.Group("/hello")
	group.GET("/world", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello World",
			"times":   1024,
		})
	})

	// show me the money
	showMeTheMoney()

	group.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(401, "Show me the money and how do you turn this on?")
	})
}

// show me the money
func showMeTheMoney() {
	fmt.Println("Show me the money")
}
