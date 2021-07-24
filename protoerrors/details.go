package protoerrors

import (
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retries have been reached or a maximum retry delay cap has been
// reached.
func (e *Error) RetryInfo(retry time.Duration) *Error {
	e.details = append(e.details, &errdetails.RetryInfo{
		RetryDelay: &durationpb.Duration{
			Seconds: int64(retry.Seconds()),
			Nanos:   int32(retry.Nanoseconds() % 10e9),
		},
	})

	return e
}

// Describes additional debugging info.
func (e *Error) DebugInfo(stackEntries []string, detail string) *Error {
	e.details = append(e.details, &errdetails.DebugInfo{
		StackEntries: stackEntries,
		Detail:       detail,
	})

	return e
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
func QuotaViolation(subject, desc string) *errdetails.QuotaFailure_Violation {
	return &errdetails.QuotaFailure_Violation{
		Subject:     subject,
		Description: desc,
	}
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryInfo and Help types for other details about handling a
// quota failure.
func (e *Error) QuotaFailure(violations ...*errdetails.QuotaFailure_Violation) *Error {
	e.details = append(e.details, &errdetails.QuotaFailure{
		Violations: violations,
	})

	return e
}

// Describes the cause of the error with structured details.
//
// Example of an error when contacting the "pubsub.googleapis.com" API when it
// is not enabled:
//
//     { "reason": "API_DISABLED"
//       "domain": "googleapis.com"
//       "metadata": {
//         "resource": "projects/123",
//         "service": "pubsub.googleapis.com"
//       }
//     }
//
// This response indicates that the pubsub.googleapis.com API is not enabled.
//
// Example of an error that is returned when attempting to create a Spanner
// instance in a region that is out of stock:
//
//     { "reason": "STOCKOUT"
//       "domain": "spanner.googleapis.com",
//       "metadata": {
//         "availableRegions": "us-central1,us-east2"
//       }
//     }
func (e *Error) ErrorInfo(reason, domain string, metadata map[string]string) *Error {
	e.details = append(e.details, &errdetails.ErrorInfo{
		Reason:   reason,
		Domain:   domain,
		Metadata: metadata,
	})

	return e
}

// A message type used to describe a single precondition failure.
func PreconditionViolation(violationType, subject, desc string) *errdetails.PreconditionFailure_Violation {
	return &errdetails.PreconditionFailure_Violation{
		Type:        violationType,
		Subject:     subject,
		Description: desc,
	}
}

// Describes what preconditions have failed.
//
// For example, if an RPC failed because it required the Terms of Service to be
// acknowledged, it could list the terms of service violation in the
// PreconditionFailure message.
func (e *Error) PreconditionFailure(violations ...*errdetails.PreconditionFailure_Violation) *Error {
	e.details = append(e.details, &errdetails.PreconditionFailure{
		Violations: violations,
	})

	return e
}

// A message type used to describe a single bad request field.
func FieldViolation(field, description string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
func (e *Error) BadRequest(violations ...*errdetails.BadRequest_FieldViolation) *Error {
	e.details = append(e.details, &errdetails.BadRequest{
		FieldViolations: violations,
	})

	return e
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
func (e *Error) RequestInfo(requestId, servingData string) *Error {
	e.details = append(e.details, &errdetails.RequestInfo{
		RequestId:   requestId,
		ServingData: servingData,
	})

	return e
}

// Describes the resource that is being accessed.
func (e *Error) ResourceInfo(resourceType, resourceName, owner, description string) *Error {
	e.details = append(e.details, &errdetails.ResourceInfo{
		ResourceType: resourceType,
		ResourceName: resourceName,
		Owner:        owner,
		Description:  description,
	})

	return e
}

// Describes a URL link.
func HelpLink(url, description string) *errdetails.Help_Link {
	return &errdetails.Help_Link{
		Description: description,
		Url:         url,
	}
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
func (e *Error) Help(links ...*errdetails.Help_Link) *Error {
	e.details = append(e.details, &errdetails.Help{
		Links: links,
	})

	return e
}

// Provides a localized error message that is safe to return to the user
// which can be attached to an RPC error.
func (e *Error) LocalizedMessage(locale, message string) *Error {
	e.details = append(e.details, &errdetails.LocalizedMessage{
		Locale:  locale,
		Message: message,
	})

	return e
}
