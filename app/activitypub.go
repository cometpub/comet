package app

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"

	"github.com/cometpub/comet/activitypub"
	"github.com/cometpub/comet/middleware"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func SendActivityPubJson(app core.App, c echo.Context, v any) error {
	data, err := json.Marshal(v)

	if err != nil {
		app.Logger().Error("Error sending ActivityPub json", "error", err)
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/activity+json")
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Writer.Write(data)

	return nil
}

func RegisterActivityPubRoutes(app core.App, group *echo.Group) {
	group.Use(
		middleware.LoadPublicationFromRequest(app),
		middleware.RequirePublication(app),
	)

	group.GET("/.well-known/webfinger", WebfingerGet(app))
}

func WebfingerGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		publication := c.Get(middleware.ContextPublication).(*models.Record)
		domain, _ := url.Parse(publication.GetString("domain"))
		resource := c.Request().URL.Query().Get("resource")

		if resource == "" {
			return apis.NewBadRequestError("resource query parameter is required", nil)
		}

		re := regexp.MustCompile(`^acct:([^@]*)@(\S*)$`)
		match := re.FindStringSubmatch(resource)

		resourceDomain := match[2]
		if resourceDomain != domain.Host {
			return apis.NewNotFoundError("resource domain not found", nil)
		}

		username := match[1]

		author := publications.FindPublicationAuthor(app, publication, username)

		if author == nil {
			return apis.NewNotFoundError("resource not found", nil)
		}

		webfinger := activitypub.PublicationAuthorToWebfinger(publication, author)

		return SendActivityPubJson(app, c, webfinger)
	}
}
