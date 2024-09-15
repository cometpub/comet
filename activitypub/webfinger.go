package activitypub

import (
	"fmt"
	"net/url"

	"github.com/pocketbase/pocketbase/models"
)

type WebfingerLink struct {
	Rel      string `json:"rel"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

type Webfinger struct {
	Account string          `json:"subject"`
	Aliases []string        `json:"aliases"`
	Links   []WebfingerLink `json:"links"`
}

func PublicationAuthorToWebfinger(hostBase string, author *models.Record) *Webfinger {
	domain, _ := url.Parse(hostBase)

	return &Webfinger{
		Account: fmt.Sprintf("acct:%s@%s", author.Username(), domain.Host),
		Aliases: []string{
			domain.JoinPath("activitypub", "authors", author.Username()).String(),
		},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: domain.JoinPath("feed").String(),
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: domain.JoinPath("activitypub", "authors", author.Username()).String(),
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: domain.JoinPath("activitypub", "authorize_interaction").String() + "?uri={uri}",
			},
		},
	}
}
