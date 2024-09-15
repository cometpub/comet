package activitypub

import (
	"encoding/pem"
	"mime"
	"net/url"
	"path/filepath"

	"github.com/cometpub/comet/publications"
	"github.com/pocketbase/pocketbase/models"

	ap "github.com/go-ap/activitypub"
)

func AuthorToActor(hostBase string, publication *models.Record, author *models.Record) *ap.Person {
	domain, _ := url.Parse(hostBase)
	slug := author.Username()

	iri := ap.IRI(domain.JoinPath("activitypub", "authors", slug).String())

	actor := ap.PersonNew(iri)
	actor.URL = iri

	actor.Name.Set(ap.DefaultLang, ap.Content(publication.GetString("title")))
	actor.Summary.Set(ap.DefaultLang, ap.Content(publication.GetString("subtitle")))
	actor.PreferredUsername.Set(ap.DefaultLang, ap.Content(author.Username()))

	actor.Inbox = ap.IRI(domain.JoinPath("activitypub", "inbox", slug).String())
	actor.Followers = ap.IRI(domain.JoinPath("activitypub", "followers", slug).String())
	actor.Outbox = ap.IRI(domain.JoinPath("activitypub", "outbox", slug).String())

	LoadActivityPubPrivateKey(publication, author)
	pubKeyBytes := GetPublicKey(publication, author)

	actor.PublicKey.Owner = iri
	actor.PublicKey.ID = ap.IRI(iri + "#main-key")
	actor.PublicKey.PublicKeyPem = string(pem.EncodeToMemory(&pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyBytes,
	}))

	if avatar := publications.RecordPropToImageSrc(hostBase, author, "avatar"); avatar != "" {
		icon := &ap.Image{}
		icon.Type = ap.ImageType
		icon.MediaType = ap.MimeType(mime.TypeByExtension(filepath.Ext(avatar)))
		icon.URL = ap.IRI(avatar)
		actor.Icon = icon
	}

	return actor
}
