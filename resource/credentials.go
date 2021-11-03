package resource

import (
	"strings"
)

var CredentialsAPI = "//credentials.animeapis.com/"

func CredentialsName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := UserName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "credentials" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         tokens[3],
	}, true
}

func FlowName(name string) (*Name, bool) {
	tokens := strings.Split(name, "/")
	if len(tokens) != 4 {
		return nil, false
	}

	parent, ok := UserName(strings.Join(tokens[:2], "/"))
	if !ok {
		return nil, false
	}

	if tokens[2] != "flows" {
		return nil, false
	}

	return &Name{
		Parent: parent,

		collection: tokens[2],
		id:         tokens[3],
	}, true
}

func CredentialsFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, CredentialsAPI) {
		return nil, false
	}

	return CredentialsName(strings.TrimPrefix(name, CredentialsAPI))
}

func FlowFullName(name string) (*Name, bool) {
	if !strings.HasPrefix(name, CredentialsAPI) {
		return nil, false
	}

	return FlowName(strings.TrimPrefix(name, CredentialsAPI))
}
