package resource

import (
	"strconv"
	"strings"
)

var ImageAPI = "//image.animeapis.com/"

func AlbumParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "organizations" && tokens[0] != "users" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func AlbumName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := AlbumParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	albumId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "albums" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         albumId,
	}, true
}

func ImageName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 6 {
		return nil, false
	}

	parent, ok := AlbumName(strings.Join(tokens[:4], "/"))
	if !ok {
		return nil, false
	}

	imageId, err := strconv.ParseInt(tokens[5], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[4] != "images" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[4],
		id:         imageId,
	}, true
}

func AlbumFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, ImageAPI) {
		return nil, false
	}

	return AlbumName(strings.TrimPrefix(name, ImageAPI))
}

func ImageFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, ImageAPI) {
		return nil, false
	}

	return ImageName(strings.TrimPrefix(name, ImageAPI))
}
