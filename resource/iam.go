package resource

import (
	"strconv"
	"strings"
)

var IAMAPI = "//iam.animeapis.com/"

func ServiceAccountParentName(name string) (*Name, bool) {
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

func RoleName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "roles" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         tokens[1],
	}, true
}

func RoleFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, IAMAPI) {
		return nil, false
	}

	return RoleName(strings.TrimPrefix(name, IAMAPI))
}

func PermissionName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 2 {
		return nil, false
	}

	if tokens[0] != "permissions" {
		return nil, false
	}

	return &Name{
		collection: tokens[0],
		id:         tokens[1],
	}, true
}

func PermissionFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, IAMAPI) {
		return nil, false
	}

	return PermissionName(strings.TrimPrefix(name, IAMAPI))
}

func ServiceAccountName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := ServiceAccountParentName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "serviceAccounts" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         tokens[3],
	}, true
}

func ServiceAccountFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, IAMAPI) {
		return nil, false
	}

	return ServiceAccountName(strings.TrimPrefix(name, IAMAPI))
}
