package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchString(t *testing.T) {
	for _, i := range []struct {
		expected bool
		all      bool
		pattern  string
		value    string
	}{
		{false, false, "{", ""},
		{false, false, "users/{user}", "organizations/123456"},
		{false, false, "users/{user}", "users"},
		{false, false, "users/{user}", "users/"},
		{false, false, "users/{user}/albums/{album}", "users//albums/123"},
		{false, false, "users/{user}/albums/{album}", "users/123/albums/"},

		{true, false, "users/{user}", "users/123456"},
		{true, false, "users/{user}", "users/123456/albums/12345"},
		{true, false, "users/{user}/albums/{album}", "users/123/albums/123"},

		{false, true, "users/{user}", "users/123456/albums/12345"},
		{false, true, "users/{user}/albums/{album}", "users/123/albums/123/images/123"},

		{true, true, "users/{user}", "users/123456"},
		{true, true, "users/{user.name}", "users/123456"},
		{true, true, "users/{user}/albums/{album}", "users/123/albums/123"},
	} {
		assert.Equal(t, i.expected, MatchString(i.pattern, i.value, i.all))
	}
}

func TestFindString(t *testing.T) {
	for _, i := range []struct {
		expected  bool
		all       bool
		variables map[string]string
		pattern   string
		value     string
	}{
		{false, false, nil, "{", ""},
		{false, false, nil, "users/{user}", "organizations/123456"},
		{false, false, nil, "users/{user}", "users"},
		{false, false, nil, "users/{user}", "users/"},
		{false, false, nil, "users/{user}/albums/{album}", "users//albums/123"},
		{false, false, nil, "users/{user}/albums/{album}", "users/123/albums/"},

		{true, false, map[string]string{"user": "123456"}, "users/{user}", "users/123456"},
		{true, false, map[string]string{"user": "123456"}, "users/{user}", "users/123456/albums/12345"},
		{true, false, map[string]string{"user": "123", "album": "123"}, "users/{user}/albums/{album}", "users/123/albums/123"},

		{false, true, nil, "users/{user}", "users/123456/albums/12345"},
		{false, true, nil, "users/{user}/albums/{album}", "users/123/albums/123/images/123"},

		{true, true, map[string]string{"user": "123456"}, "users/{user}", "users/123456"},
		{true, true, map[string]string{"user.name": "123456"}, "users/{user.name}", "users/123456"},
		{true, true, map[string]string{"user": "123", "album": "123"}, "users/{user}/albums/{album}", "users/123/albums/123"},
	} {
		vars, ok := FindString(i.pattern, i.value, i.all)

		assert.Equal(t, i.expected, ok)
		assert.Equal(t, i.variables, vars)
	}
}

func TestFindStringSubmatch(t *testing.T) {
	for _, i := range []struct {
		expected bool
		pattern  string
		value    string
		submatch string
	}{
		{false, "{", "", ""},
		{false, "users/{user}", "organizations/123456", ""},
		{false, "users/{user}", "users", ""},
		{false, "users/{user}", "users/", ""},
		{false, "users/{user}/albums/{album}", "users//albums/123", ""},
		{false, "users/{user}/albums/{album}", "users/123/albums/", ""},

		{true, "users/{user}", "users/123456", "users/123456"},
		{true, "users/{user.name}", "users/123456", "users/123456"},
		{true, "users/{user}", "users/123456/albums/12345", "users/123456"},
		{true, "users/{user}/albums/{album}", "users/123/albums/123", "users/123/albums/123"},
		{true, "users/{user}/albums/{album}", "users/123/albums/123/images/123", "users/123/albums/123"},
	} {
		submatch, ok := FindStringSubmatch(i.pattern, i.value)

		assert.Equal(t, i.expected, ok)
		assert.Equal(t, i.submatch, submatch)
	}
}

func TestReplaceString(t *testing.T) {
	for _, i := range []struct {
		expected  bool
		variables map[string]string
		pattern   string
		value     string
	}{
		{false, nil, "{", ""},
		{false, nil, "users/{user}", ""},
		{false, map[string]string{"organization": "123"}, "users/{user}", ""},
		{false, map[string]string{"user": "123"}, "users/{user}/albums/{album}", ""},

		{true, map[string]string{"user": "123456"}, "users/{user}", "users/123456"},
		{true, map[string]string{"user.name": "123456"}, "users/{user.name}", "users/123456"},
		{true, map[string]string{"user": "123", "album": "123"}, "users/{user}/albums/{album}", "users/123/albums/123"},
	} {
		value, ok := ReplaceString(i.pattern, i.variables)

		assert.Equal(t, i.expected, ok)
		assert.Equal(t, i.value, value)
	}
}
