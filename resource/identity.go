package resource

import (
	"strconv"
	"strings"
)

var IdentityAPI = "//identity.animeapis.com/"

func UserName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	userId, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return nil, false
	}

	if tokens[0] != "users" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         userId,
	}, true
}

func UserFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, IdentityAPI) {
		return nil, false
	}

	return UserName(strings.TrimPrefix(name, IdentityAPI))
}
