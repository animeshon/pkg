package iap

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type Credentials struct {
	Principal string

	UserId    string
	UserEmail string
}

func (creds *Credentials) Authenticated() bool {
	return len(creds.Principal) != 0 && creds.Principal != "anonymous"
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
		creds.Principal = xGoogIapPrincipal[0]
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
