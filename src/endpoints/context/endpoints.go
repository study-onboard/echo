package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	group := server.Server.Group("/context")
	group.Use(collectUserIdMiddleware)
	group.Use(makeLoveMiddleware)

	group.GET("/info", func(c echo.Context) error {
		c.Logger().Info("Processing in endpoints...........")
		cc := c.(*studyContext)
		return cc.JSONPretty(http.StatusOK, map[string]any{
			"user_id": cc.userId,
		}, "  ")
	})
}
