package validation

import (
	"fmt"

	"github.com/animeshon/pkg/protoerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type Error struct {
	Field    string
	Error    string
	BadValue interface{}
}

type ErrorList []*Error

func (list ErrorList) ToProtoError() *protoerrors.Error {
	validations := []*errdetails.BadRequest_FieldViolation{}
	for _, i := range list {
		validations = append(validations, protoerrors.FieldViolation(i.Field, i.Error))
	}

	if len(list) == 0 {
		return protoerrors.InvalidArgument("<nil>").BadRequest(validations...)
	}

	return protoerrors.InvalidArgument(list[0].Error).BadRequest(validations...)
}

func Required(field *field.Path) *Error {
	return &Error{Field: field.String(), Error: "The field is required."}
}

func TooLong(field *field.Path, value interface{}, maxLength int) *Error {
	return &Error{Field: field.String(), Error: "The value is too long.", BadValue: value}
}

func TooShort(field *field.Path, value interface{}, minLength int) *Error {
	return &Error{Field: field.String(), Error: "The value is too short.", BadValue: value}
}

func Invalid(field *field.Path, value interface{}, details string) *Error {
	return &Error{Field: field.String(), Error: fmt.Sprintf("The value is invalid: %s.", details), BadValue: value}
}
