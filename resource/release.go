package resource

import (
	"strconv"
	"strings"
)

var ReleaseAPI = "//release.animeapis.com/"

func ReleaseParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "users" && tokens[0] != "organizations" {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func ReleaseName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := ReleaseParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "releases" {
		return nil, false
	}

	releaseId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return &Name{
			Parent: parent,

			collection: tokens[2],
			id:         tokens[3],
		}, true
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         releaseId,
	}, true
}

func ReleaseFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, ReleaseAPI) {
		return nil, false
	}

	return ReleaseName(strings.TrimPrefix(name, ReleaseAPI))
}
