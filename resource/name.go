package resource

import (
	"strconv"
	"strings"
)

type Name struct {
	Parent *Name

	Collection string
	Id         int64
}

func GenerateName(rawName string) *Name {
	nameParts := strings.Split(rawName, "/")
	nameCollection := nameParts[0]
	nameID, _ := strconv.ParseInt(nameParts[1], 10, 64)
	return NewName(nameCollection, nameID)
}

func NewName(collection string, id int64) *Name {
	return &Name{
		Collection: collection,
		Id:         id,
	}
}

func (resource *Name) String() string {
	name := resource.Collection + "/" + strconv.FormatInt(resource.Id, 10)
	if resource.Parent != nil {
		name = resource.Parent.String() + "/" + name
	}

	return name
}

func (resource *Name) Child(collection string, id int64) *Name {
	return &Name{
		Parent: resource,

		Collection: collection,
		Id:         id,
	}
}
