package app

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"

	"github.com/cometpub/comet/activitypub"
	"github.com/cometpub/comet/middleware"
	"github.com/cometpub/comet/publications"
	ap "github.com/go-ap/activitypub"
	"github.com/go-ap/jsonld"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func SendActivityPubJson(app core.App, c echo.Context, v any) error {
	payload, err := jsonld.WithContext(jsonld.IRI(ap.ActivityBaseURI), jsonld.IRI(ap.SecurityContextURI)).Marshal(v)

	if err != nil {
		app.Logger().Error("Error encoding ActivityPub json", "error", err)
		return apis.NewApiError(http.StatusInternalServerError, "error encoding ActivityPub JSON", nil)
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/activity+json")
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Writer.Write(payload)

	return nil
}

func RegisterActivityPubRoutes(app core.App, group *echo.Group) {
	activitypub.InitKeyStore(app)

	group.Use(
		middleware.LoadPublicationFromRequest(app),
		middleware.RequirePublication(app),
	)

	group.GET("/.well-known/webfinger", WebfingerGet(app))
	group.GET("/authors/:username", ActorGet(app))
	group.GET("/activitypub/followers/:username", FollowersGet(app))
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

		payload, _ := json.Marshal(webfinger)

		c.Response().Header().Set(echo.HeaderContentType, "application/jrd+json; charset=utf-8")
		c.Response().WriteHeader(http.StatusOK)
		c.Response().Writer.Write(payload)

		return nil
	}
}

func ActorGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		publication := c.Get(middleware.ContextPublication).(*models.Record)
		username := c.PathParam("username")
		hostBase := c.Get(middleware.ContextHostBase).(string)

		author := publications.FindPublicationAuthor(app, publication, username)

		if author == nil {
			return apis.NewNotFoundError("resource not found", nil)
		}

		actor := activitypub.AuthorToActor(hostBase, publication, author)

		return SendActivityPubJson(app, c, actor)
	}
}

func FollowersGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostBase := c.Get(middleware.ContextHostBase).(string)
		username := c.PathParam("username")

		// TODO load followers from DB

		domain, _ := url.Parse(hostBase)

		collection := ap.CollectionNew(ap.IRI(domain.JoinPath("activitypub", "followers", username).String()))
		collection.TotalItems = 0

		return SendActivityPubJson(app, c, collection)
	}
}
