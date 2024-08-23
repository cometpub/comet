package feeds

import (
	"fmt"
	"net/url"
	"regexp"
)

const PAGE_SIZE = 20

type PaginationData struct {
	Page       int
	PerPage    int
	TotalItems int
	TotalPages int
}

type FeedPagination struct {
	*PaginationData
	Self     string
	First    string
	Last     string
	Next     string
	Previous string
}

func (data *PaginationData) FeedPagination(self string) *FeedPagination {
	pagination := &FeedPagination{
		PaginationData: data,
		Self:           self,
	}

	baseUrl := self

	if data.Page != 1 {
		// remove the page URL fragment from the path, accounting for an optional trailing slash
		re := regexp.MustCompile(`\/[^\/]+(\/?)$`)
		baseUrl = re.ReplaceAllString(baseUrl, "")
	}

	if data.TotalPages == 1 {
		return pagination
	}

	if data.Page == 2 {
		pagination.Previous = baseUrl
	} else if data.Page > 2 {
		pagination.Previous = safeJoinUrl(baseUrl, fmt.Sprint(data.Page-1))
		pagination.First = baseUrl
	}

	if data.Page < data.TotalPages {
		pagination.Next = safeJoinUrl(baseUrl, fmt.Sprint(data.Page+1))
	}

	if data.Page < data.TotalPages-1 {
		pagination.Last = safeJoinUrl(baseUrl, fmt.Sprint(data.TotalPages))
	}

	return pagination
}

func safeJoinUrl(base string, elem ...string) string {
	url, err := url.JoinPath(base, elem...)

	if err != nil {
		return ""
	}

	return url
}
