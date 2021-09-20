package interceptor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestParamsWalk(t *testing.T) {
	req := struct {
		Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
		Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
		Nested struct {
			Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
			Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
		} `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
		unexported interface{}
	}{
		Name:   "name/123",
		Parent: "parent/123",
		Nested: struct {
			Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
			Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
		}{
			Name:   "parent/123/nested/123",
			Parent: "parent/123",
		},
		unexported: "-",
	}

	for _, i := range []struct {
		params map[string][]string
		count  int
	}{
		{
			params: map[string][]string{},
			count:  0,
		},
		{
			params: map[string][]string{"name": {"name/123"}},
			count:  0,
		},
		{
			params: map[string][]string{"parent": {"parent/123"}},
			count:  0,
		},
		{
			params: map[string][]string{"nested.name": {"parent/123/nested/123"}},
			count:  0,
		},
		{
			params: map[string][]string{"nested.parent": {"parent/123"}},
			count:  0,
		},
		{
			params: map[string][]string{
				"name":   {"name/123"},
				"parent": {"parent/123"},
			},
			count: 0,
		},
		{
			params: map[string][]string{
				"nested.name":   {"parent/123/nested/123"},
				"nested.parent": {"parent/123"},
			},
			count: 0,
		},
		{
			params: map[string][]string{
				"name":          {"name/123"},
				"parent":        {"parent/123"},
				"nested.name":   {"parent/123/nested/123"},
				"nested.parent": {"parent/123"},
			},
			count: 0,
		},
		{
			params: map[string][]string{"other": {"-"}},
			count:  1,
		},
		{
			params: map[string][]string{"name": {"name/123"}, "other": {"-"}},
			count:  1,
		},
		{
			params: map[string][]string{"nested.name": {"parent/123/nested/123"}, "other": {"-"}},
			count:  1,
		},
		{
			params: map[string][]string{"other": {"-"}, "another": {"-"}},
			count:  2,
		},
	} {
		remainder := walk(nil, req, i.params)
		assert.Len(t, remainder, i.count, "%v", remainder)
	}
}
