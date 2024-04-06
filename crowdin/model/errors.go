package model

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	// ErrNilRequest is returned when a request for a validation is nil.
	ErrNilRequest = errors.New("request cannot be nil")
)

// Error represents the schema for the error response.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse is the error response structure from the API.
type ErrorResponse struct {
	Response *http.Response `json:"-"`

	Err Error `json:"error"`
}

// Error implements the Error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %s", r.Err.Code, r.Err.Message)
}

// ValidationError represents the schema for the invalid
// request error response.
type ValidationError struct {
	Error struct {
		Key    string `json:"key"`
		Errors []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"errors"`
	} `json:"error"`
}

// ValidationErrorResponse is the validation error response
// structure from the API.
type ValidationErrorResponse struct {
	Response *http.Response `json:"-"`

	Errors []ValidationError `json:"errors"`
	Status int
}

// Error implements the Error interface.
func (r *ValidationErrorResponse) Error() string {
	var sb strings.Builder
	for i, err := range r.Errors {
		if i != 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(fmt.Sprintf("%s: ", err.Error.Key))
		for j, e := range err.Error.Errors {
			if j != 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s (%s)", e.Message, e.Code))
		}
	}
	return sb.String()
}
