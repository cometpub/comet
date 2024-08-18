package middleware

import (
	"fmt"

	"github.com/cometpub/comet/publications"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

const HostHeader = "X-Forwarded-Host"
const ContextPublication = "publication"
const ContextHostBase = "host_base"

func LoadPublicationFromRequest(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			httpHost := req.Host

			// Redirected from the reverse proxy?
			if httpHost == "comet.pub" {
				httpHost = req.Header.Get(HostHeader)
				if httpHost == "" {
					httpHost = req.Host
				}
			}

			hostBase := "https://" + httpHost

			// TEMP
			if httpHost == "127.0.0.1:8090" || httpHost == "localhost:8090" {
				hostBase = "http://" + hostBase
				httpHost = "comet.tonysull.co"
			}

			// publications with custom domain
			publication, err := publications.FindPublicationByDomain(app, fmt.Sprintf("https://%s", httpHost))

			if err != nil {
				app.Logger().Error("Load Publication", "error", err)
				return next(c)
			}

			c.Set(ContextPublication, publication)
			c.Set(ContextHostBase, hostBase)

			return next(c)
		}
	}
}

func RequirePublication(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			record := c.Get(ContextPublication)
			if record == nil {
				app.Logger().Error("Require Publication", "error", "Publication not found")
				return apis.NewUnauthorizedError("Publication not found.", nil)
			}

			return next(c)
		}
	}
}
