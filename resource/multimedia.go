package resource

import (
	"strconv"
	"strings"
)

var MultimediaAPI = "//multimedia.animeapis.com/"

func ChapterParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "graphicNovels" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func ChapterName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := ChapterParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "chapters" {
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

func ChapterFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return ChapterName(strings.TrimPrefix(name, MultimediaAPI))
}

func EpisodeParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "animes" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         parentId,
	}, true
}

func EpisodeName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := EpisodeParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "episodes" {
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

func EpisodeFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return EpisodeName(strings.TrimPrefix(name, MultimediaAPI))
}
