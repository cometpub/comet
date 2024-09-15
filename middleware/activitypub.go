package middleware

import (
	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

const ContextActivityPubAuthor = "ap_author"

func LoadActivityPubAuthorForRequest(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			publication := c.Get(ContextPublication).(*models.Record)
			username := c.PathParam("username")

			author := publications.FindPublicationAuthor(app, publication, username)

			if author == nil {
				return apis.NewNotFoundError("resource not found", nil)
			}

			c.Set(ContextActivityPubAuthor, author)

			return next(c)
		}
	}
}
