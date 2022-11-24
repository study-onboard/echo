package binding

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/study-onboard/echo/src/server"
	"github.com/study-onboard/echo/src/vendors/web"
)

func init() {

	group := server.Server.Group("/binding")

	// JSON
	group.POST("/json", func(c echo.Context) error {
		type registerForm struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		var form registerForm
		if err := c.Bind(&form); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, form)
	})

	// Query
	group.GET("/query", func(c echo.Context) error {
		type queryForm struct {
			Page int `query:"page" xml:"page"`
			Size int `query:"size" xml:"size"`
		}

		var form queryForm
		if err := c.Bind(&form); err != nil {
			return err
		}

		return c.XMLPretty(http.StatusOK, form, "  ")
	})

	// Form
	group.PUT("/form/:id", func(c echo.Context) error {
		type bookForm struct {
			Name   string  `form:"name"`
			Author string  `form:"author"`
			Price  float32 `form:"price"`
			// 可以混合着来干
			Id string `              param:"id"`
		}

		var form bookForm
		if err := c.Bind(&form); err != nil {
			return err
		}

		return c.JSONPretty(http.StatusOK, form, "   ")
	})

	// Header
	group.DELETE("/header/:id", func(c echo.Context) error {
		type operateForm struct {
			Id   string `param:"id"`
			From string `           header:"X-FORWARD-IP"`
		}

		var form operateForm
		binder := &echo.DefaultBinder{}

		// Bind mehtod not support for bind header
		if err := binder.BindPathParams(c, &form); err != nil {
			return err
		}

		// must special
		if err := binder.BindHeaders(c, &form); err != nil {
			return err
		}

		return c.JSONPretty(http.StatusOK, form, "    ")
	})

	group.POST("/all/:id", func(c echo.Context) error {
		type allForm struct {
			Id   string `param:"id"`
			Ip   string `           header:"X-FORWARD-IP"`
			From string `                                 query:"from"`
			Name string `                                              json:"name"`
		}

		var form allForm
		if err := web.BindForm(c, &form); err != nil {
			return err
		}

		type information struct {
			Id   string `json:"id"`
			Ip   string `json:"ip"`
			From string `json:"from"`
			Name string `json:"name"`
		}

		return c.JSONPretty(http.StatusOK, &information{
			Id:   form.Id,
			Ip:   form.Ip,
			From: form.From,
			Name: form.Name,
		}, "    ")
	})
}
