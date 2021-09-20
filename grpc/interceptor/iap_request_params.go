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

// IdentityAwareProxyRequestParams returns a new unary client interceptor for x-goog-iap-request-params.
func IdentityAwareProxyRequestParams() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return protoerrors.InvalidArgument("The request is missing incoming metadata.").Err()
		}

		// Bypass Identity-Aware Proxy checks for request coming from the intranet.
		if isInternal(md) {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		header := md.Get("x-goog-iap-request-params")
		if len(header) == 0 {
			return protoerrors.InvalidArgument("The request is missing required IAP headers.").Err()
		}

		params, err := url.ParseQuery(header[0])
		if err != nil {
			return protoerrors.InvalidArgument("The request has invalid IAP headers.").Err()
		}

		body, _ := json.Marshal(req)
		logrus.Infof("[debug] request-params = %s", string(body))

		for key, value := range params {
			logrus.Infof("[debug] request-params = %s: %v", key, value)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
