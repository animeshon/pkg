package resource

import (
	"strconv"
	"strings"
)

var WebCacheAPI = "//webcache.animeapis.com/"

func CacheName(name string) (*Name, string, bool) {
	name, revision := ParseRevision(name)

	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, "", false
	}

	cacheId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, "", false
	}

	if tokens[0] != "caches" {
		return nil, "", false
	}

	return &Name{
		collection: tokens[0],
		id:         cacheId,
	}, revision, true
}

func CacheFullName(name string) (*Name, string, bool) {
	if !strings.HasPrefix(name, WebCacheAPI) {
		return nil, "", false
	}

	return CacheName(strings.TrimPrefix(name, WebCacheAPI))
}
