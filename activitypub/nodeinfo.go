package activitypub

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
)

type ServerInfoSoftware struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

type ServerInfoUsage struct {
	Users      ServerInfoUserUsage `json:"users"`
	LocalPosts int                 `json:"localPosts"`
}

type ServerInfoUserUsage struct {
	Total int `json:"total"`
}

type ServerInfo struct {
	Version   string             `json:"version"`
	Software  ServerInfoSoftware `json:"software"`
	Usage     ServerInfoUsage    `json:"usage"`
	Protocols []string           `json:"protocols"`
	Metadata  map[string]any     `json:"metadata"`
}

func GetServerInfo(app core.App) *ServerInfo {
	dao := app.Dao()

	info := &struct {
		EntriesCount      int `db:"entriesCount"`
		PublicationsCount int `db:"publicationsCount"`
	}{}

	dao.RunInTransaction(func(txDao *daos.Dao) error {
		txDao.DB().Select("COUNT(1) as entriesCount").From("entries").One(info)
		txDao.DB().Select("COUNT(1) as publicationsCount").From("publications").One(info)

		return nil
	})

	result := &ServerInfo{
		Version: "0.1.0",
		Software: ServerInfoSoftware{
			Name:       "comet.pub",
			Repository: "https://comet.pub/app",
		},
		Usage: ServerInfoUsage{
			Users: ServerInfoUserUsage{
				Total: info.PublicationsCount,
			},
			LocalPosts: info.EntriesCount,
		},
		Protocols: []string{
			"activitypub",
			"micropub",
		},
		Metadata: map[string]any{},
	}

	return result
}
