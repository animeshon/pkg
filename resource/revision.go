package resource

import "strings"

func ParseRevision(name string) (string, string) {
	tokens := strings.Split(name, "@")
	if len(tokens) == 1 {
		return tokens[0], ""
	}

	return tokens[0], strings.Join(tokens[1:], "@")
}
