package middleware

import (
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

const ContextEntry = "entry"
const ContextEntries = "entries"
const ContextPagination = "pagination"

func LoadPublicationEntryFromRequest(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			publication := c.Get(ContextPublication).(*models.Record)

			slug := c.PathParam("slug")

			record, err := app.Dao().FindFirstRecordByFilter(
				"entries", "publication={:publication} && slug={:slug}",
				dbx.Params{
					"publication": publication.Id,
					"slug":        slug,
				},
			)

			if err != nil {
				return apis.NewNotFoundError("", err)
			}

			c.Set(ContextEntry, record)

			return next(c)
		}
	}
}

func LoadPublicationEntriesFromRequest(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			publication := c.Get(ContextPublication).(*models.Record)
			category := c.PathParam("category")
			typeParam := c.PathParam("type")
			currentPage, _ := intOrDefault(c.PathParam("page"), 1)

			var entryType publications.EntryType

			if typeParam != "" {
				t, ok := publications.ParseEntryType(typeParam)

				if !ok {
					// Should the request hit `/:page`?
					parsed, err := strconv.Atoi(typeParam)
					if err == nil {
						currentPage = parsed
						typeParam = ""
						entryType = ""
					} else {
						return apis.NewNotFoundError("", nil)
					}
				} else {
					entryType = t
				}
			}

			entriesCount, err := publications.FindEntriesCountForDomain(app, publication.GetString("domain"), entryType, category)

			if err != nil {
				return c.NoContent(http.StatusInternalServerError)
			}

			if entriesCount == 0 && category != "" {
				return apis.NewNotFoundError("Category not found", nil)
			}

			if (currentPage-1)*feeds.PAGE_SIZE > entriesCount {
				return apis.NewNotFoundError("", nil)
			}

			entries, err := publications.FindEntriesForPublication(app, publication.Id, category, entryType, feeds.PAGE_SIZE, (currentPage-1)*feeds.PAGE_SIZE)

			if err != nil {
				return apis.NewApiError(http.StatusInternalServerError, "Failed to load publication entries", nil)
			}

			page := feeds.PaginationData{
				Page:       currentPage,
				PerPage:    feeds.PAGE_SIZE,
				TotalItems: entriesCount,
				TotalPages: int(math.Ceil(float64(entriesCount) / feeds.PAGE_SIZE)),
			}

			baseUrl, _ := url.JoinPath(publication.GetString("domain"), c.Request().URL.String())

			pagination := page.FeedPagination(baseUrl)

			c.Set(ContextEntries, entries)
			c.Set(ContextPagination, pagination)

			return next(c)
		}
	}
}

func intOrDefault(value string, defaultValue int) (int, error) {
	if value == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue, err
	}

	return parsed, nil
}
