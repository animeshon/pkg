package resource

import (
	"strconv"
	"strings"
)

var WebPageAPI = "//webpage.animeapis.com/"

func SiteName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "sites" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         tokens[1],
	}, true
}

func PageName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := SiteName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	pageId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "pages" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         pageId,
	}, true
}

func SiteFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, WebPageAPI) {
		return nil, false
	}

	return SiteName(strings.TrimPrefix(name, WebPageAPI))
}

func PageFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, WebPageAPI) {
		return nil, false
	}

	return PageName(strings.TrimPrefix(name, WebPageAPI))
}
