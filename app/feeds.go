package app

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/lib"
	"github.com/cometpub/comet/middleware"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	accept "github.com/timewasted/go-accept-headers"
)

func RegisterFeedRoutes(app core.App, group *echo.Group) {
	group.Use(
		middleware.LoadPublicationFromRequest(app),
		middleware.RequirePublication(app),
		middleware.LoadPublicationEntriesFromRequest(app),
		middleware.LoadCrawlersForRequest,
	)

	// main routes for Atom feeds
	group.GET("/feed", FeedGetByAccepts)
	group.GET("/feed/:page", FeedGetByAccepts)

	// handle routes for /articles, /notes, /photos, and /bookmarks
	group.GET("/:type", FeedGetByAccepts)
	group.GET("/:type/:page", FeedGetByAccepts)

	// handle feeds for category tags
	group.GET("/category/:category", FeedGetByAccepts)
	group.GET("/category/:category/:page", FeedGetByAccepts)
}

func XMLWithXSLT(c echo.Context, feed feeds.XmlFeed, xslt string) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	c.Response().Write([]byte(fmt.Sprintf(`<?xml-stylesheet href="%s" type="text/xsl"?>`, xslt)))

	encoder := xml.NewEncoder(c.Response())
	encoder.Indent("", "\t")

	return encoder.Encode(feed.FeedXml())
}

func FeedGetByAccepts(c echo.Context) error {
	header := c.Request().Header.Get(echo.HeaderAccept)

	types := accept.Parse(header)

	for _, acceptType := range types {
		if acceptType.Type == "application" {
			if acceptType.Subtype == "feed+json" {
				return JsonFeedGet(c)
			} else if acceptType.Subtype == "rss+xml" {
				return RSSFeedGet(c)
			}
		}
	}

	return AtomFeedGet(c)
}

func AtomFeedGet(c echo.Context) error {
	isCrawler := c.Get(middleware.ContextIsCrawler).(bool)
	if isCrawler {
		return CrawlerFeedGet(c)
	}

	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)
	hostBase := c.Get(middleware.ContextHostBase).(string)

	feed := publications.PublicationToFeed(hostBase, publication, entries, pagination)

	return XMLWithXSLT(c, &feeds.Atom{feed}, "/static/feed.xsl")
}

func JsonFeedGet(c echo.Context) error {
	isCrawler := c.Get(middleware.ContextIsCrawler).(bool)
	if isCrawler {
		return CrawlerFeedGet(c)
	}

	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)
	hostBase := c.Get(middleware.ContextHostBase).(string)

	feed := publications.PublicationToFeed(hostBase, publication, entries, pagination)

	jsonFeed, _ := feed.ToJSON()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, jsonFeed)
}

func RSSFeedGet(c echo.Context) error {
	isCrawler := c.Get(middleware.ContextIsCrawler).(bool)
	if isCrawler {
		return CrawlerFeedGet(c)
	}

	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)
	hostBase := c.Get(middleware.ContextHostBase).(string)

	feed := publications.PublicationToFeed(hostBase, publication, entries, pagination)

	rssFeed, _ := feed.ToRss()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, rssFeed)
}

func CrawlerFeedGet(c echo.Context) error {
	publication := c.Get(middleware.ContextPublication).(*models.Record)
	hostBase := c.Get(middleware.ContextHostBase).(string)

	seo := &feeds.SEO{
		Title:       publication.GetString("title"),
		Description: publication.GetString("subtitle"),
		Image:       publications.RecordPropToImageSrc(hostBase, publication, "logo"),
		Url:         publication.GetString("domain"),
	}

	forceUrl, _ := url.Parse(hostBase + c.Request().URL.String())
	query := forceUrl.Query()
	query.Add("force", "true")
	forceUrl.RawQuery = query.Encode()

	return lib.Render(c, http.StatusOK, feeds.CrawlerFeed(seo, forceUrl.String()))
}
