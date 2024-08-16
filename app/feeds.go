package app

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/middleware"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/models"
)

func XMLWithXSLT(xml string, xslt string) string {
	re := regexp.MustCompile(`^<\?xml ([^)]+)\?>`)
	str := re.ReplaceAllString(xml, fmt.Sprintf(`<?xml-stylesheet href="%s" type="text/xsl"?>`, xslt))
	return str
}

func AtomFeedGet(c echo.Context) error {
	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)

	feed := publications.PublicationToFeed(publication, entries, pagination)

	atomFeed, _ := feed.ToAtom()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, XMLWithXSLT(atomFeed, "/static/feed.xsl"))
}

func JsonFeedGet(c echo.Context) error {
	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)

	feed := publications.PublicationToFeed(publication, entries, pagination)

	jsonFeed, _ := feed.ToJSON()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, jsonFeed)
}

func RSSFeedGet(c echo.Context) error {
	publication := c.Get(middleware.ContextPublication).(*models.Record)
	entries := c.Get(middleware.ContextEntries).([]*models.Record)
	pagination := c.Get(middleware.ContextPagination).(*feeds.FeedPagination)

	feed := publications.PublicationToFeed(publication, entries, pagination)

	rssFeed, _ := feed.ToRss()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, rssFeed)
}
