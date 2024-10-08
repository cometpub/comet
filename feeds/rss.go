package feeds

// rss support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"encoding/xml"
	"fmt"
	"time"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type RssFeedXml struct {
	XMLName          xml.Name `xml:"rss"`
	Version          string   `xml:"version,attr"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Channel          *RssFeed
}

type RssContent struct {
	XMLName xml.Name `xml:"content:encoded"`
	Content string   `xml:",cdata"`
}

type RssImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   int      `xml:"width,omitempty"`
	Height  int      `xml:"height,omitempty"`
}

type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"` // Author used
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`       // created or updated
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"` // updated used
	Category       string   `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Ttl            int      `xml:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Image          *RssImage
	TextInput      *RssTextInput
	Items          []*RssItem `xml:"item"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`       // required
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Content     *RssContent
	Authors     []string `xml:"author,omitempty"`
	Categories  []string `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Enclosures  []*RssEnclosure
	Guid        *RssGuid // Id used
	PubDate     string   `xml:"pubDate,omitempty"` // created or updated
	Source      string   `xml:"source,omitempty"`
}

type RssEnclosure struct {
	//RSS 2.0 <enclosure url="http://example.com/file.mp3" length="123456789" type="audio/mpeg" />
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type RssGuid struct {
	//RSS 2.0 <guid isPermaLink="true">http://inessential.com/2002/09/01.php#a2</guid>
	XMLName     xml.Name `xml:"guid"`
	Id          string   `xml:",chardata"`
	IsPermaLink string   `xml:"isPermaLink,attr,omitempty"` // "true", "false", or an empty string
}

type Rss struct {
	*Feed
}

// create a new RssItem with a generic Item struct's data
func newRssItem(i *Item) *RssItem {
	item := &RssItem{
		Title:       i.Title,
		Description: i.Description,
		PubDate:     i.Updated.Format(time.RFC1123Z),
		Categories:  i.Categories,
	}
	if i.Id != "" {
		item.Guid = &RssGuid{Id: i.Id, IsPermaLink: i.IsPermaLink}
	}
	if i.Link != nil {
		item.Link = i.Link.Href
	}
	if len(i.Content) > 0 {
		item.Content = &RssContent{Content: i.Content}
	}
	if i.Source != nil {
		item.Source = i.Source.Href
	}

	for _, enclosure := range i.Enclosures {
		item.Enclosures = append(item.Enclosures, &RssEnclosure{Url: enclosure.Url, Type: enclosure.Type, Length: enclosure.Length})
	}

	for _, author := range i.Authors {
		item.Authors = append(item.Authors, author.Name)
	}

	return item
}

// create a new RssFeed with a generic Feed struct's data
func (r *Rss) RssFeed() *RssFeed {
	pub := r.Updated.Format(time.RFC1123Z)
	build := r.Updated.Format(time.RFC1123Z)
	author := ""
	if len(r.Authors) > 0 {
		author = r.Authors[0].Email
		if len(r.Authors[0].Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Authors[0].Email, r.Authors[0].Name)
		}
	}

	var image *RssImage
	if r.Image != nil {
		image = &RssImage{Url: r.Image.Url, Title: r.Image.Title, Link: r.Image.Link, Width: r.Image.Width, Height: r.Image.Height}
	}

	var href string
	for _, link := range r.Links {
		if link.Rel == "self" {
			href = link.Href
			break
		}
	}
	channel := &RssFeed{
		Title:          r.Title,
		Link:           href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
		Image:          image,
	}
	for _, i := range r.Items {
		channel.Items = append(channel.Items, newRssItem(i))
	}
	return channel
}

// FeedXml returns an XML-Ready object for an Rss object
func (r *Rss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.RssFeed().FeedXml()

}

// FeedXml returns an XML-ready object for an RssFeed object
func (r *RssFeed) FeedXml() interface{} {
	return &RssFeedXml{
		Version:          "2.0",
		Channel:          r,
		ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
	}
}
