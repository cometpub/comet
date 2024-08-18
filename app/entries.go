package app

import (
	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/middleware"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func RegisterEntryRoutes(app core.App, group *echo.Group) {
	group.Use(
		middleware.LoadPublicationFromRequest(app),
		middleware.RequirePublication(app),
		middleware.LoadPublicationEntryFromRequest(app),
	)

	group.GET("/:slug", EntryGet(app))
}

func EntryGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		publication := c.Get(middleware.ContextPublication).(*models.Record)
		entry := c.Get(middleware.ContextEntry).(*models.Record)
		hostBase := c.Get(middleware.ContextHostBase).(string)

		feed := publications.PublicationToFeed(hostBase, publication, []*models.Record{entry}, &feeds.FeedPagination{})

		return XMLWithXSLT(c, &feeds.Atom{feed}, "/static/entry.xsl")
	}
}
