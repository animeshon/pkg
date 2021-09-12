package resource

import (
	"strconv"
	"strings"
)

var LibraryAPI = "//library.animeapis.com/"

func PlaylistParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "audiences" && tokens[0] != "users" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func PlaylistName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := PlaylistParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	playlistId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "playlists" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         playlistId,
	}, true
}

func PlaylistItemName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 {
		return nil, false
	}

	parent, ok := PlaylistName(strings.Join(tokens[:4], "/"))
	if !ok {
		return nil, false
	}

	itemId, err := strconv.ParseInt(tokens[5], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[4] != "items" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[4],
		id:         itemId,
	}, true
}

func PlaylistFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, LibraryAPI) {
		return nil, false
	}

	return PlaylistName(strings.TrimPrefix(name, LibraryAPI))
}

func PlaylisyItemFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, LibraryAPI) {
		return nil, false
	}

	return PlaylistItemName(strings.TrimPrefix(name, LibraryAPI))
}
