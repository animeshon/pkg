package resource

import (
	"strconv"
	"strings"
)

var KnowledgeAPI = "//knowledge.animeapis.com/"

func ContributionParentName(name string) (*Name, bool) {
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

func ContributionName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := ContributionParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	albumId, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[2] != "contributions" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         albumId,
	}, true
}

func ContributionFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, KnowledgeAPI) {
		return nil, false
	}

	return ContributionName(strings.TrimPrefix(name, KnowledgeAPI))
}
