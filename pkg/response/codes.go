package response

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Code sets the HTTP status code for the response
// The code must be a valid HTTP status code
func (b *jsonBuilder) Code(code int) *jsonBuilder {
	if b.isCodeValid(code) {
		b.code = code
	}
	return b
}

// isItemValid checks if a http status code is valid
// It updates the dirty error field if validation fails
func (b *jsonBuilder) isCodeValid(code int) bool {
	if http.StatusText(code) == "" {
		b.howl = fmt.Errorf("undefined http status code=%d", code)
		return false
	}
	return true
}

// CodeFromProto sets the HTTP status code
// for the response to match the grpc status code
func (b *jsonBuilder) CodeFromProto(code codes.Code) *jsonBuilder {
	switch code {
	case codes.OK:
		b.code = http.StatusOK // 200 OK
	case codes.InvalidArgument:
		b.code = http.StatusBadRequest // 400 Bad Request
	case codes.NotFound:
		b.code = http.StatusNotFound // 404 Not Found
	case codes.Internal:
		b.code = http.StatusInternalServerError // 500 Internal Server Error
	case codes.AlreadyExists:
		b.code = http.StatusConflict // 409 Conflict
	case codes.Unimplemented:
		b.code = http.StatusNotImplemented // 501 Not Implemented
	case codes.Unavailable:
		b.code = http.StatusServiceUnavailable // 503 Service Unavailable
	case codes.Canceled:
		b.code = 499 // 499 Client Closed Request
	case codes.Unknown:
		b.code = http.StatusInternalServerError // 500 Internal Server Error
	case codes.DeadlineExceeded:
		b.code = http.StatusGatewayTimeout // 504 Gateway Timeout
	case codes.PermissionDenied:
		b.code = http.StatusForbidden // 403 Forbidden
	case codes.Unauthenticated:
		b.code = http.StatusUnauthorized // 401 Unauthorized
	case codes.ResourceExhausted:
		b.code = http.StatusTooManyRequests // 429 Too Many Requests
	case codes.FailedPrecondition:
		b.code = http.StatusPreconditionFailed // 412 Precondition Failed
	case codes.Aborted:
		b.code = http.StatusConflict // 409 Conflict
	case codes.OutOfRange:
		b.code = http.StatusRequestedRangeNotSatisfiable // 416 Requested Range Not Satisfiable
	default:
		b.howl = fmt.Errorf("unknown grpc status code=%v", code)
	}
	return b
}

// OK sets the HTTP status code to 200 OK.
func (b *jsonBuilder) OK() *jsonBuilder {
	b.code = http.StatusOK
	return b
}

// BadRequest sets the HTTP status code to 400 Bad Request.
func (b *jsonBuilder) BadRequest() *jsonBuilder {
	b.code = http.StatusBadRequest
	return b
}

// InternalError sets the HTTP status code to 500 Internal Server Error.
func (b *jsonBuilder) InternalError() *jsonBuilder {
	b.code = http.StatusInternalServerError
	return b
}

// Created sets the HTTP status code to 201 Created.
func (b *jsonBuilder) Created() *jsonBuilder {
	b.code = http.StatusCreated
	return b
}

// NotFound sets the HTTP status code to 404 Not Found.
func (b *jsonBuilder) NotFound() *jsonBuilder {
	b.code = http.StatusNotFound
	return b
}
