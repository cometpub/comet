package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cometpub/comet/middleware"
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func RegisterPublishRoutes(app core.App, group *echo.Group) {
	group.Use(
		middleware.RequireUserAuth,
		middleware.LoadPublicationFromRequest(app),
		middleware.RequirePublication(app),
		middleware.RequirePublicationAuthor(app),
	)

	group.POST("/note", PublishNotePost(app))
}

func PublishNotePost(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		authUser := c.Get(apis.ContextAuthRecordKey).(*models.Record)

		coll, err := app.Dao().FindCollectionByNameOrId("entries")

		if err != nil {
			return err
		}

		publication := c.Get(middleware.ContextPublication).(*models.Record)

		record := models.NewRecord(coll)
		form := forms.NewRecordUpsert(app, record)

		newSlug := ""
		name := form.Data()["name"].(string)

		if name != "" {
			newSlug = slug.Make(name)
		} else {
			newSlug = strconv.FormatInt(time.Now().Unix(), 10)
		}

		form.LoadData(map[string]any{
			"slug":        newSlug,
			"authors":     []string{authUser.Id},
			"publication": publication.Id,
			"published":   time.Now(),
			"type":        "note",
		})

		form.LoadRequest(c.Request(), "")

		if err := form.Submit(); err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}
