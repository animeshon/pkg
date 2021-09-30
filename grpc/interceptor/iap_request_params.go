package interceptor

import (
	"context"
	"net/url"
	"reflect"
	"strings"

	"github.com/animeshon/pkg/protoerrors"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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

		// If no 'x-goog-iap-request-params' header is found send the request to the
		// next interceptor.
		//
		// The IAP assertion header combined with other headers should be enough to
		// confirm that the request came from a trusted source.
		//
		// The responsibility of checking IAP headers is not delegated to this
		// interceptor.
		header := md.Get("x-goog-iap-request-params")
		if len(header) == 0 {
			return handler(ctx, req)
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
