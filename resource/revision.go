package resource

import "strings"

func ParseRevision(name string) (string, string) {
	tokens := strings.Split(name, "@")
	if len(tokens) == 1 {
		return tokens[0], ""
	}

	if len(tokens) == 2 {
		return tokens[0], tokens[1]
	}

	return tokens[0], strings.Join(tokens[1:], "@")
}
