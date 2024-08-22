package app

import (
	"net/http"

	"github.com/cometpub/comet/auth"
	"github.com/cometpub/comet/middleware"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	accept "github.com/timewasted/go-accept-headers"
)

func RegisterProfileRoutes(app core.App, group *echo.Group) {
	group.Use(apis.RequireRecordAuth("users"))

	group.GET("/profile.json", ProfileGet)
	group.GET("/profile.xml", ProfileGet)
}

func ProfileGet(c echo.Context) error {
	record := c.Get(apis.ContextAuthRecordKey).(*models.Record)
	hostBase := c.Get(middleware.ContextHostBase).(string)

	user := auth.ParseUser(hostBase, record)

	header := c.Request().Header.Get(echo.HeaderAccept)
	headerAccepts := accept.Parse(header)

	if headerAccepts.Accepts("application/json") {
		return c.JSONPretty(http.StatusOK, user, "\t")
	}

	return c.XMLPretty(http.StatusOK, user, "\t")
}
