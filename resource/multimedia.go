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

	chapterId, err := strconv.ParseInt(tokens[3], 10, 64)
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
		id:         chapterId,
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

	episodeId, err := strconv.ParseInt(tokens[3], 10, 64)
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
		id:         episodeId,
	}, true
}

func EpisodeFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return EpisodeName(strings.TrimPrefix(name, MultimediaAPI))
}

func AnimeName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "animes" {
		return nil, false
	}

	animeId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return &Name{
			collection: tokens[0],
			id:         tokens[1],
		}, true
	}

	return &Name{
		collection: tokens[0],
		id:         animeId,
	}, true
}

func AnimeFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return AnimeName(strings.TrimPrefix(name, MultimediaAPI))
}

func LightNovelName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "lightNovels" {
		return nil, false
	}

	lightNovelId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return &Name{
			collection: tokens[0],
			id:         tokens[1],
		}, true
	}

	return &Name{
		collection: tokens[0],
		id:         lightNovelId,
	}, true
}

func LightNovelFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return LightNovelName(strings.TrimPrefix(name, MultimediaAPI))
}

func GraphicNovelName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "graphicNovels" {
		return nil, false
	}

	graphicNovelId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return &Name{
			collection: tokens[0],
			id:         tokens[1],
		}, true
	}

	return &Name{
		collection: tokens[0],
		id:         graphicNovelId,
	}, true
}

func GraphicNovelFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return GraphicNovelName(strings.TrimPrefix(name, MultimediaAPI))
}

func VisualNovelName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "visualNovels" {
		return nil, false
	}

	visualNovelId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return &Name{
			collection: tokens[0],
			id:         tokens[1],
		}, true
	}

	return &Name{
		collection: tokens[0],
		id:         visualNovelId,
	}, true
}

func VisualNovelFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, MultimediaAPI) {
		return nil, false
	}

	return VisualNovelName(strings.TrimPrefix(name, MultimediaAPI))
}
