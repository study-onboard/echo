package context

import "github.com/labstack/echo/v4"

type studyContext struct {
	echo.Context

	// current user id
	userId string
}
