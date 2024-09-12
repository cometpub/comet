package activitypub

import (
	"fmt"
	"net/url"

	"github.com/pocketbase/pocketbase/models"
)

type WebfingerLink struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type Webfinger struct {
	Account string          `json:"subject"`
	Aliases []string        `json:"aliases"`
	Links   []WebfingerLink `json:"links"`
}

func PublicationAuthorToWebfinger(publication *models.Record, author *models.Record) *Webfinger {
	domain, _ := url.Parse(publication.GetString("domain"))

	return &Webfinger{
		Account: fmt.Sprintf("acct:%s@%s", author.Username(), domain.Host),
		Aliases: []string{
			fmt.Sprintf("https://%s.comet.pub/activitypub/users/%s", publication.GetString("slug"), author.Username()),
		},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: domain.String(),
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: domain.String(),
			},
		},
	}
}
