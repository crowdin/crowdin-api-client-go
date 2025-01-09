package model

import (
	"encoding/json"
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
	Code    any    `json:"code"`
	Message string `json:"message"`
}

// UnmarshalJSON unmarshals the code field. It allows the code field
// to be either an integer or a string.
func (e *Error) UnmarshalJSON(data []byte) error {
	type Alias Error
	aux := &struct {
		Code json.RawMessage `json:"code"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var intCode int
	if err := json.Unmarshal(aux.Code, &intCode); err == nil {
		e.Code = intCode
		return nil
	}

	var strCode string
	if err := json.Unmarshal(aux.Code, &strCode); err == nil {
		e.Code = strCode
		return nil
	}

	return nil
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
		Key    string  `json:"key"`
		Errors []Error `json:"errors"`
	} `json:"error"`
}

type ValidationBatchError struct {
	Errors []ValidationError `json:"errors"`
	Index  *int64        `json:"index"`
}

// ValidationErrorResponse is the validation error response
// structure from the API.
type ValidationErrorResponse struct {
	Response *http.Response `json:"-"`

	Errors []ValidationError `json:"errors"`
	Status int
}

type ValidationErrorBatchResponse struct {
	Response *http.Response `json:"-"`

	Errors []ValidationBatchError `json:"errors"`
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

			var code string
			switch v := e.Code.(type) {
			case int:
				code = fmt.Sprintf("%d", v)
			case string:
				code = v
			default:
				code = "n/a"
			}

			sb.WriteString(fmt.Sprintf("%s (%s)", e.Message, code))
		}
	}
	return sb.String()
}

func (r *ValidationErrorBatchResponse) Error() string {
	var sb strings.Builder
	for index, errB := range r.Errors {
		sb.WriteString(fmt.Sprintf("\n Index: %d \n", index))
		for i, err := range errB.Errors {
			if i != 0 {
				sb.WriteString("; ")
			}
			sb.WriteString(fmt.Sprintf("%s: ", err.Error.Key))
			for j, e := range err.Error.Errors {
				if j != 0 {
					sb.WriteString(", ")
				}
	
				var code string
				switch v := e.Code.(type) {
				case int:
					code = fmt.Sprintf("%d", v)
				case string:
					code = v
				default:
					code = "n/a"
				}
	
				sb.WriteString(fmt.Sprintf("%s (%s)", e.Message, code))
			}
		}
	}
	return sb.String()
}

// GraphQLError represents a single GraphQL error.
type GraphQLError struct {
	Message    string         `json:"message"`
	Extensions map[string]any `json:"extensions"`
	Locations  []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
}

// Error implements error interface.
func (e GraphQLError) Error() string {
	return fmt.Sprintf("%s, Locations: %+v", e.Message, e.Locations)
}

// GraphQLErrorResponse represents a GraphQL error response.
type GraphQLErrorResponse struct {
	Errors []GraphQLError `json:"errors"`
}

// Error implements the Error interface.
func (r *GraphQLErrorResponse) Error() string {
	sb := strings.Builder{}
	for i, err := range r.Errors {
		if i != 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}
