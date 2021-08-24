package resource

import (
	"strconv"
	"strings"
)

var VisionAPI = "//vision.animeapis.com/"

func ImageAnnotationParentName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "users" {
		return nil, false
	}

	return &Name{
		Collection: tokens[0],
		Id:         parentId,
	}, true
}

func ImageAnnotationName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := ImageAnnotationParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	albumId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "imageAnnotations" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		Collection: tokens[2],
		Id:         albumId,
	}, true
}

func ImageAnnotationFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, VisionAPI) {
		return nil, false
	}

	return ImageAnnotationName(strings.TrimPrefix(name, VisionAPI))
}
