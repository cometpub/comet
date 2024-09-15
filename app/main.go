package app

import (
	"net/http"

	"github.com/cometpub/comet/middleware"
	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func ToDo(c echo.Context) error {
	return apis.NewApiError(http.StatusNotImplemented, "not implemented", nil)
}

func PermanentRedirect(to string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, to)
	}
}

func TemporaryRedirect(to string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, to)
	}
}

func NotImplemented(c echo.Context) error {
	return apis.NewApiError(http.StatusNotImplemented, "not implemented", nil)
}

func resolveIndex(c echo.Context) error {
	publication := c.Get(middleware.ContextPublication)

	if publication != nil {
		return TemporaryRedirect("/feed")(c)
	} else {
		return ToDo(c)
	}
}

func InitAppRoutes(e *core.ServeEvent, app core.App) {
	e.Router.Use(echoMiddleware.Logger())

	// placeholder for the admin dashboard app
	appGroup := e.Router.Group("/app", middleware.LoadAuthContextFromCookie(app), middleware.RequireUserAuth)
	appGroup.GET("", ToDo)

	// distinguish between root requests for the Comet home page and publications redirected from the reverse proxy
	e.Router.GET("", resolveIndex, middleware.LoadPublicationFromRequest(app))

	feedGroup := e.Router.Group("")
	RegisterFeedRoutes(app, feedGroup)

	entryGroup := e.Router.Group("/posts")
	RegisterEntryRoutes(app, entryGroup)

	activityPubGroup := e.Router.Group("")
	RegisterActivityPubRoutes(app, activityPubGroup)
}
