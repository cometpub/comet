package app

import (
	"net/http"

	"github.com/cometpub/comet/middleware"
	"github.com/labstack/echo/v5"
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

func InitAppRoutes(e *core.ServeEvent, app core.App) {
	appGroup := e.Router.Group("/app", middleware.LoadAuthContextFromCookie(app), middleware.RequireUserAuth)

	appGroup.GET("", ToDo)

	publicationGroup := e.Router.Group("", middleware.LoadPublicationFromRequest(app), middleware.LoadPublicationEntriesFromRequest(app))
	publicationGroup.GET("/feed", PermanentRedirect("/atom.xml"))
	publicationGroup.GET("/atom.xml", AtomFeedGet)
	publicationGroup.GET("/feed.json", JsonFeedGet)
	publicationGroup.GET("/rss.xml", RSSFeedGet)
}
