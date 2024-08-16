package auth

import (
	"github.com/a-h/templ"
	"github.com/cometpub/comet/lib"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type RegisterFormValue struct {
	username       string
	password       string
	passwordRepeat string
}

func (lfv RegisterFormValue) Validate() error {
	return validation.ValidateStruct(&lfv,
		validation.Field(&lfv.username, validation.Required, validation.Length(3, 50)),
		validation.Field(&lfv.password, validation.Required),
	)
}

func getRegisterFormValue(c echo.Context) RegisterFormValue {
	return RegisterFormValue{
		username:       c.FormValue("username"),
		password:       c.FormValue("password"),
		passwordRepeat: c.FormValue("passwordRepeat"),
	}
}

func RegisterRegisterRoutes(app core.App, group echo.Group) {
	group.GET("/register", func(c echo.Context) error {
		if c.Get(apis.ContextAuthRecordKey) != nil {
			return c.Redirect(302, "/app")
		}

		return lib.Render(c, 200, Register(RegisterFormValue{}, lib.RegisterError{}))
	})

	group.POST("/register", func(c echo.Context) error {
		form := getRegisterFormValue(c)
		err := form.Validate()

		registerErr := lib.RegisterError{}

		if err == nil {
			registerErr = lib.Register(app, c, form.username, form.password, form.passwordRepeat)
		} else {
			// TODO parse err
		}

		if err != nil {
			component := lib.HtmxRender(
				c,
				func() templ.Component { return RegisterForm(form, registerErr) },
				func() templ.Component { return Register(form, registerErr) },
			)
			return lib.Render(c, 200, component)
		}

		return lib.HtmxRedirect(c, "/app")
	})
}
