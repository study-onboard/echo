package context

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	group := server.Server.Group("/context")
	group.Use(collectUserIdMiddleware)

	group.GET("/info", func(c echo.Context) error {
		log.Println("Processing in endpoints...........")
		cc := c.(*studyContext)
		return cc.JSONPretty(http.StatusOK, map[string]any{
			"user_id": cc.userId,
		}, "  ")
	})
}

func collectUserIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &studyContext{
			Context: c,
		}
		cc.userId = "ID1314"

		log.Println("Before collect user id middleware")
		result := next(cc)
		log.Println("After collect user id middleware")
		return result
	}
}
