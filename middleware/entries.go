package middleware

import (
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
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

			record, err := publications.FindEntryAndSlug(app, publication.Id, slug)

			if err != nil {
				app.Logger().Error("Load Publication Entry", "error", "Not Found", "publication", publication.Id, "slug", slug)
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
			hostBase := c.Get(ContextHostBase).(string)
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
						app.Logger().Error("Load Publication Entries", "error", "Params invalid", "typeParam", typeParam, "currentPage", currentPage)
						return apis.NewNotFoundError("", nil)
					}
				} else {
					entryType = t
				}
			}

			entriesCount, err := publications.FindEntriesCountForDomain(app, publication.GetString("domain"), entryType, category)

			if err != nil {
				app.Logger().Error("Load Publication Entries", "message", "Loading entries", "error", err)
				return c.NoContent(http.StatusInternalServerError)
			}

			if entriesCount == 0 && category != "" {
				app.Logger().Error("Load Publication Entries", "error", "Category not found", "category", category)
				return apis.NewNotFoundError("Category not found", nil)
			}

			if (currentPage-1)*feeds.PAGE_SIZE > entriesCount {
				app.Logger().Error("Load Publication Entries", "error", "Page not found", "page", currentPage, "entriesCount", entriesCount)
				return apis.NewNotFoundError("", nil)
			}

			entries, err := publications.FindEntriesForPublication(app, publication.Id, category, entryType, feeds.PAGE_SIZE, (currentPage-1)*feeds.PAGE_SIZE)

			if err != nil {
				app.Logger().Error("Load Publication Entries", "message", "Failed to load publications", "error", err)
				return apis.NewApiError(http.StatusInternalServerError, "Failed to load publication entries", nil)
			}

			page := feeds.PaginationData{
				Page:       currentPage,
				PerPage:    feeds.PAGE_SIZE,
				TotalItems: entriesCount,
				TotalPages: int(math.Ceil(float64(entriesCount) / feeds.PAGE_SIZE)),
			}

			baseUrl, _ := url.JoinPath(hostBase, c.Request().URL.Path)

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
