package publications

import (
	"fmt"
	"mime"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/cometpub/comet/feeds"
	"github.com/cometpub/comet/lib"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

type EntryType = string

const (
	ArticleType  EntryType = "article"
	BookmarkType EntryType = "bookmark"
	NoteType     EntryType = "note"
	PhotoType    EntryType = "photo"
)

var (
	entryTypeMap = map[string]EntryType{
		"articles":  ArticleType,
		"bookmarks": BookmarkType,
		"notes":     NoteType,
		"photos":    PhotoType,
	}
)

func ParseEntryType(str string) (EntryType, bool) {
	t, ok := entryTypeMap[strings.ToLower(str)]
	return t, ok
}

func FindEntryAndSlug(app core.App, includePrivateEntries bool, publication string, slug string) (*models.Record, error) {
	filters := []string{"publication={:publication}", "slug={:slug}"}
	params := dbx.Params{"publication": publication, "slug": slug}

	if !includePrivateEntries {
		filters = append(filters, "status={:status}")
		params["status"] = "public"
	}

	record, err := app.Dao().FindFirstRecordByFilter(
		"entries",
		strings.Join(filters, " && "),
		params,
	)

	if err != nil {
		return nil, err
	}

	// expand the "authors" relations
	if errs := app.Dao().ExpandRecord(record, []string{"authors", "categories", "photos"}, nil); len(errs) > 0 {
		app.Logger().Error("failed to expand entry authors", "error", errs, "entry", record.GetString("slug"))
		return nil, fmt.Errorf("failed to expand entry authors %v", errs)
	}

	return record, nil
}

func FindEntriesForPublication(app core.App, publication string, category string, entryType EntryType, count int, offset int) ([]*models.Record, error) {
	dao := app.Dao()
	filters := []string{"publication={:publication}"}
	params := dbx.Params{"publication": publication}

	if category != "" {
		filters = append(filters, "categories.slug?={:category}")
		params["category"] = category
	}

	if entryType != "" {
		filters = append(filters, "type={:type}")
		params["type"] = entryType
	}

	records, err := dao.FindRecordsByFilter(
		"entries",
		strings.Join(filters, " && "),
		"-published",
		count,
		offset,
		params,
	)

	if err != nil {
		return nil, err
	}

	// expand the "authors" and "categories" relations
	if errs := dao.ExpandRecords(records, []string{"authors", "categories", "photos"}, nil); len(errs) > 0 {
		return nil, fmt.Errorf("failed to expand: %v", errs)
	}

	return records, nil
}

func EntryToFeedItem(hostBase string, record *models.Record) *feeds.Item {
	slug := record.GetString("slug")
	link, err := url.JoinPath(hostBase, "posts", slug)
	if err != nil {
		link = ""
	}

	entry := &feeds.Item{
		Link:      &feeds.Link{Href: link, Rel: "self"},
		Id:        slug,
		Updated:   record.Updated.Time(),
		Published: record.GetDateTime("published").Time(),
	}

	featuredImage := RecordPropToImageSrc(hostBase, record, "featuredImage")
	if featuredImage != "" {
		mimetype := mime.TypeByExtension(filepath.Ext(featuredImage))
		entry.Enclosures = append(
			entry.Enclosures,
			&feeds.Enclosure{Url: featuredImage, Type: mimetype})
	}

	for _, photo := range record.ExpandedAll("photos") {
		url := RecordPropToImageSrc(hostBase, photo, "src")
		mimetype := mime.TypeByExtension(filepath.Ext(url))
		entry.Enclosures = append(
			entry.Enclosures,
			&feeds.Enclosure{
				Url:    url,
				Type:   mimetype,
				Width:  photo.GetInt("width"),
				Height: photo.GetInt("height"),
			},
		)
	}

	for _, author := range record.ExpandedAll("authors") {
		icon := RecordPropToImageSrcThumbnail(hostBase, author, "avatar", "100x100")
		entry.Authors = append(entry.Authors, &feeds.Person{
			Name:     author.GetString("name"),
			Icon:     icon,
			Username: author.GetString("username"),
		})
	}

	for _, category := range record.ExpandedAll("categories") {
		entry.Categories = append(entry.Categories, category.GetString("slug"))
	}

	entryType := record.GetString("type")

	switch entryType {
	case ArticleType:
		entry.Title = record.GetString("name")
		entry.Description = lib.MarkdownToHTML(record.GetString("summary"))
		entry.Content = lib.MarkdownToHTML(record.GetString("content"))
	case NoteType:
		entry.Description = lib.MarkdownToHTML(record.GetString("summary"))
	case BookmarkType:
		entry.Description = lib.MarkdownToHTML(record.GetString("summary"))
	case PhotoType:
		entry.Description = lib.MarkdownToHTML(record.GetString("content"))
	}

	return entry
}
