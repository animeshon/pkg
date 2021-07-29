package iam

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Name    string
	Service string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s [%s]: %s", e.Service, e.Name, e.Err)
}
func (e *Error) Unwrap() error {
	return e.Err
}

func IsErrorNotFound(err error) bool {
	status, ok := status.FromError(err)
	if !ok {
		return false
	}

	return status.Code() == codes.NotFound
}
