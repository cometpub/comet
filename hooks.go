package comet

import (
	"embed"

	"github.com/cometpub/comet/app"
	"github.com/cometpub/comet/auth"
	"github.com/cometpub/comet/middleware"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed all:static
var staticDir embed.FS

// StaticDirFS contains the embedded dist directory files (without the "static" prefix)
var StaticDirFS = echo.MustSubFS(staticDir, "static")

func bindAppHooks(pb core.App) {
	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/static/*", apis.StaticDirectoryHandler(StaticDirFS, false))

		authGroup := e.Router.Group("/auth", middleware.LoadAuthContextFromCookie(pb))
		auth.RegisterLoginRoutes(pb, *authGroup)
		auth.RegisterRegisterRoutes(pb, *authGroup)

		app.InitAppRoutes(e, pb)

		return nil
	})
}
