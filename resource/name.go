package resource

import (
	"strconv"
	"strings"
)

type Name struct {
	Parent *Name

	collection string
	id         interface{}
}

func NewName(collection string, id string) *Name {
	return &Name{
		collection: collection,
		id:         id,
	}
}

func NewNameInt64(collection string, id int64) *Name {
	return &Name{
		collection: collection,
		id:         id,
	}
}

func (resource *Name) String() string {
	var name []string

	if resource.Parent != nil {
		name = append(name, resource.Parent.String())
	}

	if len(resource.collection) != 0 {
		name = append(name, resource.collection)
	}

	if resource.id != nil {
		switch resource.id.(type) {
		case string:
			name = append(name, resource.id.(string))
		case int64:
			name = append(name, strconv.FormatInt(resource.id.(int64), 10))
		default:
			panic("resource id must be a string or an int64")
		}
	}

	return strings.Join(name, "/")
}

func (resource *Name) Child(collection string, id string) *Name {
	return &Name{
		Parent: resource,

		collection: collection,
		id:         id,
	}
}

func (resource *Name) ChildInt64(collection string, id int64) *Name {
	return &Name{
		Parent: resource,

		collection: collection,
		id:         id,
	}
}

func (resource *Name) ID() interface{} {
	return resource.id
}

func (resource *Name) Collection() string {
	return resource.collection
}
