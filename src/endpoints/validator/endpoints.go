package validator

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
)

func init() {
	group := server.Server.Group("/validator")

	group.POST("/standard", func(c echo.Context) error {
		type standardForm struct {
			Name     string `json:"name"     validate:"required"`
			Position string `json:"position" validate:"required"`
		}

		var form standardForm
		if err := c.Bind(&form); err != nil {
			return err
		}

		if err := c.Validate(&form); err != nil {
			return err
		}

		return c.JSONPretty(http.StatusOK, form, "    ")
	})
}
