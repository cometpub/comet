package app

import (
	"encoding/json"
	"fmt"
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

	group.GET("/.well-known/host-meta", WellKnownMetaGet(app))
	group.GET("/.well-known/nodeinfo", WellKnownNodeInfoGet)
	group.GET("/nodeinfo", ServerInfoGet(app))
	group.GET("/.well-known/webfinger", WebfingerGet(app))
	group.GET("/activitypub/authors/:username", ActorGet(app), middleware.LoadActivityPubAuthorForRequest(app))
	group.GET("/activitypub/followers/:username", FollowersGet(app), middleware.LoadActivityPubAuthorForRequest(app))
	group.GET("/activitypub/outbox/:username", NotImplemented, middleware.LoadActivityPubAuthorForRequest(app))
}

func WellKnownMetaGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostBase := c.Get(middleware.ContextHostBase).(string)

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)

		c.Response().Write([]byte(fmt.Sprintf(`<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0"><Link rel="lrdd" type="application/xrd+xml" template="%s/.well-known/webfinger?resource={uri}"/></XRD>`, hostBase)))

		return nil
	}
}

func WellKnownNodeInfoGet(c echo.Context) error {
	hostBase := c.Get(middleware.ContextHostBase).(string)
	domain, _ := url.Parse(hostBase)

	info := map[string][]activitypub.WebfingerLink{
		"links": {
			{
				Href: domain.JoinPath("nodeinfo").String(),
				Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.1",
			},
		},
	}

	return c.JSON(http.StatusOK, info)
}

func ServerInfoGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		nodeInfo := activitypub.GetServerInfo(app)

		return c.JSON(http.StatusOK, nodeInfo)
	}
}

func WebfingerGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostBase := c.Get(middleware.ContextHostBase).(string)
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

		webfinger := activitypub.PublicationAuthorToWebfinger(hostBase, author)

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
		hostBase := c.Get(middleware.ContextHostBase).(string)

		author := c.Get(middleware.ContextActivityPubAuthor).(*models.Record)

		actor := activitypub.AuthorToActor(hostBase, publication, author)

		return SendActivityPubJson(app, c, actor)
	}
}

func FollowersGet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		hostBase := c.Get(middleware.ContextHostBase).(string)
		author := c.Get(middleware.ContextActivityPubAuthor).(*models.Record)

		// TODO load followers from DB

		domain, _ := url.Parse(hostBase)

		collection := ap.CollectionNew(ap.IRI(domain.JoinPath("activitypub", "followers", author.Username()).String()))
		collection.TotalItems = 0

		return SendActivityPubJson(app, c, collection)
	}
}
