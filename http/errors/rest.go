package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrInternalServerError = "internal server error"
	ErrNotFoundError       = "request resource not found"
	ErrUnauthorized        = "Unauthorized"
)

// httpError is the struct type of http errors
type httpError struct {
	Code     int   `json:"-"`
	Message  any   `json:"message"`
	Internal error `json:"-"`
}

// NewHTTPError generates new http error
func NewHTTPError(code int, message any, internal error) error {
	return &httpError{
		Code:     code,
		Message:  message,
		Internal: internal,
	}
}

// Error satisfies error interface.
func (he *httpError) Error() string {
	if he.Internal == nil {
		return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
	}
	return fmt.Sprintf("code=%d, message=%v, internal=%v", he.Code, he.Message, he.Internal)
}

// NewBadRequestError generates a new bad request http error
func NewBadRequestError(message any, internal error) *httpError {
	return &httpError{
		Code:     http.StatusBadRequest,
		Message:  message,
		Internal: internal,
	}
}

// NewInternalServerError generates a new internal server error
func NewInternalServerError(internal error) *httpError {
	return &httpError{
		Code:     http.StatusInternalServerError,
		Message:  ErrInternalServerError,
		Internal: internal,
	}
}

// NewNotFoundError generates not found error
func NewNotFoundError(internal error) *httpError {
	return &httpError{
		Code:     http.StatusNotFound,
		Message:  ErrNotFoundError,
		Internal: internal,
	}
}

// NewUnauthorizedError generates a new ErrUnauthorized error
func NewUnauthorizedError(internal error) *httpError {
	return &httpError{
		Code:     http.StatusUnauthorized,
		Message:  ErrUnauthorized,
		Internal: internal,
	}
}

// ParseErrorForResponse parses the http error from error
func ParseErrorForResponse(err error) *httpError {
	var httpErr *httpError
	if errors.As(err, &httpErr) {
		return &httpError{
			Code:     httpErr.Code,
			Message:  httpErr.Message,
			Internal: httpErr.Internal,
		}
	}
	return nil
}
