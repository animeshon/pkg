package iap

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type Credentials struct {
	AuthenticationIdentity string

	UserId    string
	UserEmail string

	Anonymous bool
}

func (creds *Credentials) Authenticated() bool {
	return len(creds.AuthenticationIdentity) != 0
}

func (creds *Credentials) GetPrincipal() error {
	return nil
}

func FromIncomingContext(ctx context.Context) (*Credentials, bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}

	creds := &Credentials{}

	xGoogIapPrincipal := headers.Get("x-goog-iap-principal")
	if len(xGoogIapPrincipal) != 0 {
		creds.AuthenticationIdentity = xGoogIapPrincipal[0]
	}

	xGoogIapAnonymous := headers.Get("x-goog-iap-anonymous")
	if len(xGoogIapAnonymous) != 0 {
		creds.Anonymous, _ = strconv.ParseBool(xGoogIapAnonymous[0])
	}

	xGoogIapUserId := headers.Get("x-goog-authenticated-user-id")
	if len(xGoogIapUserId) != 0 {
		creds.UserId = xGoogIapUserId[0]
	}

	xGoogUserEmail := headers.Get("x-goog-authenticated-user-email")
	if len(xGoogUserEmail) != 0 {
		creds.UserEmail = xGoogUserEmail[0]
	}

	return creds, true
}
