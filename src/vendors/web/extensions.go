package web

import "github.com/labstack/echo/v4"

// form binder
var formBinder = &echo.DefaultBinder{}

// Bind form
//
// bind request parameters to form:
// 1. header
// 2. query
// 3. body
// 4. path
func BindForm(c echo.Context, form any) error {
	if err := formBinder.BindHeaders(c, form); err != nil {
		return err
	}

	if err := formBinder.BindQueryParams(c, form); err != nil {
		return err
	}

	if err := formBinder.BindBody(c, form); err != nil {
		return err
	}

	if err := formBinder.BindPathParams(c, form); err != nil {
		return err
	}

	return nil
}
