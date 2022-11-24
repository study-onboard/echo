package context

import "github.com/labstack/echo/v4"

func collectUserIdMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &studyContext{
			Context: c,
		}
		cc.userId = "ID1314"

		c.Logger().Info("Before collect user id middleware")
		result := next(cc)
		c.Logger().Info("After collect user id middleware")
		return result
	}
}

func makeLoveMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Warn("Begin make love.......................")
		result := next(c)
		c.Logger().Error("Make love complete!!")
		return result
	}
}
