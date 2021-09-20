package interceptor

import (
	"context"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/animeshon/pkg/protoerrors"
	"github.com/golang/protobuf/proto"
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

func walk(tree []string, i interface{}, params map[string][]string) map[string][]string {
	value := reflect.ValueOf(i)

	switch value.Kind() {
	case reflect.Ptr:
		if !value.Elem().IsValid() {
			return params
		}

		return walk(tree, value.Elem().Interface(), params)
	case reflect.Interface:
		return walk(tree, value.Elem().Interface(), params)
	case reflect.Struct:
		for j := 0; j < value.NumField(); j += 1 {
			if !value.Field(j).CanInterface() {
				continue
			}

			tag, ok := reflect.TypeOf(i).Field(j).Tag.Lookup("protobuf")
			if !ok {
				continue
			}

			properties := proto.Properties{}
			properties.Parse(tag)

			walk(append(tree, properties.OrigName), value.Field(j).Interface(), params)
		}
	case reflect.String:
		index := strings.Join(tree, ".")
		param := params[index]
		if len(param) == 0 {
			return params
		}

		if param[0] == value.Interface().(string) {
			delete(params, index)
		}

		return params
	}

	return params
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

		remainder := walk(nil, req, params)
		if len(remainder) > 0 {
			return nil, protoerrors.InvalidArgument("The request body did not match the request parameters specified in the header or path segments.").Err()
		}

		return handler(ctx, req)
	}
}
