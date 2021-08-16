package resource

import (
	"strconv"
	"strings"
)

func CacheName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	parentId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "caches" {
		return nil, false
	}

	return &Name{
		Collection: tokens[0],
		Id:         parentId,
	}, true
}
