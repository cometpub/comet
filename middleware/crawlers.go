package middleware

import (
	"github.com/labstack/echo/v5"
	agents "github.com/monperrus/crawler-user-agents"
)

const ContextIsCrawler = "is_crawler"

func LoadCrawlersForRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().UserAgent()

		isCrawler := agents.IsCrawler(userAgent)
		force := c.QueryParams().Has("force")

		c.Set(ContextIsCrawler, isCrawler && !force)

		return next(c)
	}
}
