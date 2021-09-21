package protomerge

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func MergeWithFieldMask(dst, src proto.Message, mask *fieldmaskpb.FieldMask) {
	for _, path := range mask.GetPaths() {
		field := src.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(path))
		value := src.ProtoReflect().Get(field)

		fmt.Println(value.String())
	}
}
