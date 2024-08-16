package feeds

import (
	"encoding/xml"
	"time"
)

const ns = "http://www.w3.org/2005/Atom"
const nsComet = "https://comet.pub/Atom"

type AtomPerson struct {
	Name     string `xml:"name,omitempty"`
	Uri      string `xml:"uri,omitempty"`
	Email    string `xml:"email,omitempty"`
	Icon     string `xml:"comet:avatar,omitempty"`
	Username string `xml:"comet:username,omitempty"`
}

type AtomSummary struct {
	XMLName xml.Name `xml:"summary"`
	Content string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
}

type AtomCategory struct {
	Term   string `xml:"term,attr"` // required
	Scheme string `xml:"schema,attr,omitempty"`
	Label  string `xml:"label,attr,omitempty"`
}

type AtomContent struct {
	XMLName xml.Name `xml:"content"`
	Content string   `xml:",chardata"`
	Type    string   `xml:"type,attr"`
}

type AtomAuthor struct {
	XMLName xml.Name `xml:"author"`
	AtomPerson
}

type AtomContributor struct {
	XMLName xml.Name `xml:"contributor"`
	AtomPerson
}

type AtomEntry struct {
	XMLName     xml.Name `xml:"entry"`
	Xmlns       string   `xml:"xmlns,attr,omitempty"`
	Title       string   `xml:"title"`   // required
	Updated     string   `xml:"updated"` // required
	Id          string   `xml:"id"`      // required
	Categories  []string `xml:"category,omitempty"`
	Content     *AtomContent
	Rights      string `xml:"rights,omitempty"`
	Source      string `xml:"source,omitempty"`
	Published   string `xml:"published,omitempty"`
	Contributor *AtomContributor
	Links       []AtomLink    // required if no child 'content' elements
	Summary     *AtomSummary  // required if content has src or content is base64
	Authors     []*AtomAuthor // required if feed lacks an author
}

// Multiple links with different rel can coexist
type AtomLink struct {
	//Atom 1.0 <link rel="enclosure" type="audio/mpeg" title="MP3" href="http://www.example.org/myaudiofile.mp3" length="1234" />
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
	Length  string   `xml:"length,attr,omitempty"`
}

type AtomGenerator struct {
	Uri     string `xml:"uri,attr,omitempty"`
	Version string `xml:"version,attr,omitempty"`
	Content string `xml:",chardata"`
}

type AtomFeed struct {
	XMLName     xml.Name        `xml:"feed"`
	Xmlns       string          `xml:"xmlns,attr"`
	XmlnsComet  string          `xml:"xmlns:comet,attr"`
	Title       string          `xml:"title"`   // required
	Id          string          `xml:"id"`      // required
	Updated     string          `xml:"updated"` // required
	Categories  []*AtomCategory `xml:"category,omitempty"`
	Icon        string          `xml:"icon,omitempty"`
	Logo        string          `xml:"logo,omitempty"`
	Rights      string          `xml:"rights,omitempty"` // copyright used
	Subtitle    string          `xml:"subtitle,omitempty"`
	Links       []*AtomLink
	Authors     []*AtomAuthor `xml:"author,omitempty"`
	Contributor *AtomContributor
	Generator   *AtomGenerator `xml:"generator,omitempty"`
	Entries     []*AtomEntry   `xml:"entry"`
}

type Atom struct {
	*Feed
}

func newAtomEntry(i *Item) *AtomEntry {
	id := i.Id
	link := i.Link
	if link == nil {
		link = &Link{}
	}

	authors := []*AtomAuthor{}
	for _, author := range i.Authors {
		authors = append(authors, &AtomAuthor{AtomPerson: AtomPerson{Name: author.Name, Email: author.Email, Icon: author.Icon, Username: author.Username}})
	}

	link_rel := link.Rel
	if link_rel == "" {
		link_rel = "alternate"
	}

	x := &AtomEntry{
		Title:     i.Title,
		Links:     []AtomLink{{Href: link.Href, Rel: link_rel, Type: link.Type}},
		Id:        id,
		Authors:   authors,
		Updated:   i.Updated.Format(time.RFC3339),
		Published: i.Published.Format(time.RFC3339),
	}

	// if there's a description, assume it's html
	if len(i.Description) > 0 {
		x.Summary = &AtomSummary{Content: i.Description, Type: "html"}
	}

	// if there's a content, assume it's html
	if len(i.Content) > 0 {
		x.Content = &AtomContent{Content: i.Content, Type: "html"}
	}

	for _, enclosure := range i.Enclosures {
		x.Links = append(x.Links, AtomLink{Href: enclosure.Url, Rel: "enclosure", Type: enclosure.Type, Length: enclosure.Length})
	}

	return x
}

// create a new AtomFeed with a generic Feed struct's data
func (a *Atom) AtomFeed() *AtomFeed {
	updated := a.Updated.Format(time.RFC3339)
	links := []*AtomLink{}
	for _, link := range a.Links {
		links = append(links, &AtomLink{Href: link.Href, Rel: link.Rel, Type: link.Type, Length: link.Length})
	}
	feed := &AtomFeed{
		Xmlns:      ns,
		XmlnsComet: nsComet,
		Title:      a.Title,
		Links:      links,
		Subtitle:   a.Description,
		Id:         a.Id,
		Updated:    updated,
		Rights:     a.Copyright,
	}
	for _, author := range a.Authors {
		feed.Authors = append(feed.Authors, &AtomAuthor{AtomPerson: AtomPerson{Name: author.Name, Email: author.Email, Icon: author.Icon, Username: author.Username}})
	}
	for _, e := range a.Items {
		feed.Entries = append(feed.Entries, newAtomEntry(e))
	}
	return feed
}

// FeedXml returns an XML-Ready object for an Atom object
func (a *Atom) FeedXml() interface{} {
	return a.AtomFeed()
}

// FeedXml returns an XML-ready object for an AtomFeed object
func (a *AtomFeed) FeedXml() interface{} {
	return a
}
