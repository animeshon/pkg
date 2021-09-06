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
		collection: tokens[0],
		id:         parentId,
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

	annotationId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "imageAnnotations" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         annotationId,
	}, true
}

func ImageAnalysisName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 8 {
		return nil, false
	}

	parent, ok := ImageName(strings.Join(tokens[:6], "/"))
	if !ok {
		return nil, false
	}

	analysisId := int64(-1)
	if tokens[7] != "latest" {
		var err error
		analysisId, err = strconv.ParseInt(tokens[7], 10, 64)
		if err != nil {
			return nil, false
		}
	}

	if tokens[6] != "analyses" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[6],
		id:         analysisId,
	}, true
}

func ImageAnnotationFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, VisionAPI) {
		return nil, false
	}

	return ImageAnnotationName(strings.TrimPrefix(name, VisionAPI))
}

func ImageAnalysisFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, VisionAPI) {
		return nil, false
	}

	return ImageAnalysisName(strings.TrimPrefix(name, VisionAPI))
}
