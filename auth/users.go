package auth

import (
	"github.com/cometpub/comet/publications"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/types"
)

type User struct {
	Id      string         `db:"id" json:"-" xml:"-"`
	Created types.DateTime `db:"created" json:"created" xml:"created"`
	Updated types.DateTime `db:"updated" json:"updated" xml:"updated"`

	Username        string `db:"username" json:"username" xml:"username"`
	Email           string `db:"email" json:"email" xml:"email"`
	EmailVisibility bool   `db:"emailVisibility" json:"-" xml:"-"`
	Verified        bool   `db:"verified" json:"verified" xml:"verified"`

	Name   string `db:"name" json:"name" xml:"name"`
	Avatar string `db:"avatar" json:"avatar,omitempty" xml:"avatar,omitempty"`
}

func ParseUser(hostBase string, record *models.Record) *User {
	if record == nil || record.Collection().Name != "users" {
		return nil
	}

	return &User{
		Id:              record.Id,
		Created:         record.Created,
		Updated:         record.Updated,
		Username:        record.GetString("username"),
		Email:           record.GetString("email"),
		EmailVisibility: record.GetBool("emailVisibility"),
		Verified:        record.GetBool("verified"),
		Name:            record.GetString("name"),
		Avatar:          publications.RecordPropToImageSrc(hostBase, record, "avatar"),
	}
}
