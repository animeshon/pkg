package interceptor

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/animeshon/pkg/protoerrors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func isInternal(md metadata.MD) bool {
	internal := md.Get("x-envoy-internal")
	if len(internal) == 0 {
		return false
	}

	yes, err := strconv.ParseBool(internal[0])
	if err != nil {
		return false
	}

	return yes
}

func isAnonymous(md metadata.MD) bool {
	anonymous := md.Get("x-goog-iap-anonymous")
	if len(anonymous) == 0 {
		return false
	}

	yes, err := strconv.ParseBool(anonymous[0])
	if err != nil {
		return false
	}

	return yes
}

// IdentityAwareProxyRequestParams returns a new unary client interceptor for x-goog-iap-request-params.
func IdentityAwareProxyRequestParams() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, protoerrors.InvalidArgument("The request is missing incoming metadata.").Err()
		}

		// Bypass Identity-Aware Proxy checks for anonymous requests or requests
		// coming from the intranet.
		if isInternal(md) || isAnonymous(md) {
			return handler(ctx, req)
		}

		header := md.Get("x-goog-iap-request-params")
		if len(header) == 0 {
			return nil, protoerrors.InvalidArgument("The request is missing required IAP headers.").Err()
		}

		params, err := url.ParseQuery(header[0])
		if err != nil {
			return nil, protoerrors.InvalidArgument("The request has invalid IAP headers.").Err()
		}

		body, _ := json.Marshal(req)
		logrus.Infof("[debug] request-params = %s", string(body))

		for key, value := range params {
			logrus.Infof("[debug] request-params = %s: %v", key, value)
		}

		return handler(ctx, req)
	}
}
