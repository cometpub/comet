package comet

import (
	_ "github.com/cometpub/comet/migrations"

	"github.com/pocketbase/pocketbase"
)

type Comet struct {
	*pocketbase.PocketBase
}

func New() Comet {
	comet := Comet{pocketbase.New()}

	bindAppHooks(comet)

	return comet
}
