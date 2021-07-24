package iap

import (
	"context"
	"strings"

	jwtvalidator "github.com/animeshon/pkg/jwt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JwtInterceptor struct {
	iss string
	aud string
}

func Jwt(iss, aud string) *JwtInterceptor {
	return &JwtInterceptor{iss: iss, aud: aud}
}

var StatusUnauthenticatedInvalid = status.New(codes.Unauthenticated, "The authorization credentials provided for the request are invalid. Check the value of the Authorization HTTP request header.")

func (interceptor *JwtInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, StatusUnauthenticatedInvalid.Err() // TODO: should be missing.
		}

		header := md["authorization"]
		if len(header) == 0 {
			return nil, StatusUnauthenticatedInvalid.Err() // TODO: should be missing.
		}

		validator := jwtvalidator.NewJwtValidator(interceptor.aud, interceptor.iss)

		value := header[0]
		if strings.HasPrefix(strings.ToLower(value), "bearer ") {
			value = value[7:]
		}

		token, err := jwt.Parse(value, validator.ValidationKeyGetter)
		if err != nil {
			logrus.WithError(err).Error("failed to validate first-party jwt")
			return nil, StatusUnauthenticatedInvalid.Err()
		}

		return handler(context.WithValue(ctx, "(token)", token), req)
	}
}
