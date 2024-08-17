package publications

import (
	"fmt"
	"net/url"

	"github.com/cometpub/comet/feeds"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func FindPublicationByDomain(app core.App, domain string) (*models.Record, error) {
	record, err := app.Dao().FindFirstRecordByFilter(
		"publications",
		"domain={:domain}",
		dbx.Params{"domain": domain},
	)

	if err != nil {
		return nil, err
	}

	// expand the "authors" relations
	if errs := app.Dao().ExpandRecord(record, []string{"authors"}, nil); len(errs) > 0 {
		app.Logger().Error("failed to expand publication authors", "error", errs, "publication", record.GetString("slug"))
		return nil, fmt.Errorf("failed to expand publication authors %v", errs)
	}

	return record, nil
}

func FindEntriesCountForDomain(app core.App, domain string, entryType EntryType, category string) (int, error) {
	expr := dbx.HashExp{"domain": domain}
	view := "publication_entries_count"

	if entryType != "" {
		expr["type"] = entryType
	}

	if category != "" {
		expr["slug"] = category
		view = "publication_category_entries_count"
	}

	records, err := app.Dao().FindRecordsByExpr(
		view,
		expr,
	)

	if err != nil {
		return 0, err
	}

	count := 0

	for _, r := range records {
		count += r.GetInt("items")
	}

	return count, nil
}

func PublicationToFeed(hostBase string, record *models.Record, entries []*models.Record, pagination *feeds.FeedPagination) *feeds.Feed {
	domain := record.GetString("domain")

	links := []*feeds.Link{{Href: pagination.Self, Rel: "self"}}
	if pagination.First != "" {
		links = append(links, &feeds.Link{Href: pagination.First, Rel: "first"})
	}
	if pagination.Last != "" {
		links = append(links, &feeds.Link{Href: pagination.Last, Rel: "last"})
	}
	if pagination.Previous != "" {
		links = append(links, &feeds.Link{Href: pagination.Previous, Rel: "previous"})
	}
	if pagination.Next != "" {
		links = append(links, &feeds.Link{Href: pagination.Next, Rel: "next"})
	}

	icon := RecordPropToImageSrcThumbnail(hostBase, record, "icon", "100x100")
	logo := RecordPropToImageSrc(hostBase, record, "logo")

	feed := &feeds.Feed{
		Title:     record.GetString("title"),
		Links:     links,
		Subtitle:  record.GetString("subtitle"),
		Id:        pagination.Self,
		Updated:   record.Updated.Time(),
		Published: record.GetDateTime("published").Time(),
		Icon:      &feeds.Image{Url: icon},
		Image:     &feeds.Image{Url: logo},
	}

	for _, author := range record.ExpandedAll("authors") {
		icon, _ := url.JoinPath(domain, author.GetString("avatar"))

		feed.Authors = append(feed.Authors, &feeds.Person{
			Name:     author.GetString("name"),
			Icon:     icon,
			Username: author.GetString("username"),
		})
	}

	for _, entry := range entries {
		feed.Items = append(feed.Items, EntryToFeedItem(hostBase, entry))
	}

	return feed
}
