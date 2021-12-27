package resource

import (
	"strconv"
	"strings"
)

var TrackerAPI = "//tracker.animeapis.com/"

func TrackerParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "users" && tokens[0] != "audiences" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func TrackerName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := TrackerParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "trackers" {
		return nil, false
	}

	playlistId, err := strconv.ParseInt(tokens[3], 10, 64)
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
		id:         playlistId,
	}, true
}

func TrackerFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, TrackerAPI) {
		return nil, false
	}

	return TrackerName(strings.TrimPrefix(name, TrackerAPI))
}

func ActivityName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 {
		return nil, false
	}

	parent, ok := TrackerName(strings.Join(tokens[:4], "/"))
	if !ok {
		return nil, false
	}

	if tokens[4] != "activities" {
		return nil, false
	}

	playlistId, err := strconv.ParseInt(tokens[5], 10, 64)
	if err != nil {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[4],
		id:         playlistId,
	}, true
}

func ActivityFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, TrackerAPI) {
		return nil, false
	}

	return ActivityName(strings.TrimPrefix(name, TrackerAPI))
}
