package model

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse_Error(t *testing.T) {
	err := &ErrorResponse{
		Response: &http.Response{},
		Err: Error{
			Code:    404,
			Message: "Resource Not Found",
		},
	}

	assert.Equal(t, "404 Resource Not Found", err.Error())
	assert.NotNil(t, err.Response)
}

func TestValidationErrorResponse_Error(t *testing.T) {
	err := &ValidationErrorResponse{
		Response: &http.Response{},
		Errors: []ValidationError{
			{
				Error: struct {
					Key    string  `json:"key"`
					Errors []Error `json:"errors"`
				}{
					Key: "name",
					Errors: []Error{
						{
							Code:    "isEmpty",
							Message: "Value is required and can't be empty",
						},
					},
				},
			},
			{
				Error: struct {
					Key    string  `json:"key"`
					Errors []Error `json:"errors"`
				}{
					Key: "sourceLanguage",
					Errors: []Error{
						{
							Code:    "required",
							Message: "Field is required",
						},
						{
							Code:    "notFound",
							Message: "Field not found",
						},
					},
				},
			},
		},
		Status: http.StatusBadRequest,
	}

	expected := "name: Value is required and can't be empty (isEmpty); sourceLanguage: Field is required (required), Field not found (notFound)"
	result := err.Error()

	assert.Equal(t, expected, result)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.NotNil(t, err.Response)
}

func TestValidationErrorResponse_Error_SingleError(t *testing.T) {
	err := &ValidationErrorResponse{
		Response: &http.Response{},
		Errors: []ValidationError{
			{
				Error: struct {
					Key    string  `json:"key"`
					Errors []Error `json:"errors"`
				}{
					Key: "name",
					Errors: []Error{
						{
							Code:    "required",
							Message: "name is required",
						},
					},
				},
			},
		},
		Status: http.StatusBadRequest,
	}

	expected := "name: name is required (required)"
	result := err.Error()

	assert.Equal(t, expected, result)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.NotNil(t, err.Response)
}

func TestValidationErrorResponse_Error_EmptyErrors(t *testing.T) {
	response := &ValidationErrorResponse{
		Errors: []ValidationError{},
	}

	assert.Equal(t, "", response.Error())
}

func TestParseErrorResponse(t *testing.T) {
	cases := []struct {
		resp *http.Response
		body []byte
		code int
		err  string
	}{
		{
			resp: &http.Response{StatusCode: http.StatusNotFound},
			body: []byte(`{
				"error": {
					"message": "Resource Not Found",
					"code": 404
				}
			}`),
			code: 404,
			err:  "404 Resource Not Found",
		},
		{
			resp: &http.Response{StatusCode: http.StatusForbidden},
			body: []byte(`{
				"error": {
					"message": "Forbidden",
					"code": 403
				}
			}`),
			code: 403,
			err:  "403 Forbidden",
		},
		{
			resp: &http.Response{StatusCode: http.StatusUnauthorized},
			body: []byte(`{
				"error": {
					"message": "Unauthorized",
					"code": 401
				}
			}`),
			code: 401,
			err:  "401 Unauthorized",
		},
	}

	for _, tt := range cases {
		t.Run(tt.err, func(t *testing.T) {
			err := handleErrorResponse(tt.resp, tt.body)

			verr, ok := err.(*ErrorResponse)
			assert.True(t, ok)
			assert.NotNil(t, verr)
			assert.Equal(t, tt.code, verr.Response.StatusCode)
			assert.Equal(t, tt.code, verr.Err.Code)
			assert.Equal(t, tt.err, verr.Error())
		})
	}
}

func TestParseValidationErrorResponse(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusBadRequest,
	}

	cases := []struct {
		name string
		body []byte
		err  string
	}{
		{
			name: "single error",
			body: []byte(`{
				"errors": [
					{
						"error": {
							"key": "credentials",
							"errors": [
								{
									"code": 0,
									"message": "The server returned the following message: Translator API Authorization Failed."
								}
							]
						}
					}
				]
			}`),
			err: "credentials: The server returned the following message: Translator API Authorization Failed. (0)",
		},
		{
			name: "multiple errors",
			body: []byte(`{
				"errors": [
					{
						"error": {
							"key": "name",
							"errors": [
								{
									"code": "isEmpty",
									"message": "Value is required and can't be empty"
								}
							]
						}
					},
					{
						"error": {
							"key": "type",
							"errors": [
								{
									"code": "notInArray",
									"message": "The input was not found in the haystack"
								}
							]
						}
					}
				]
			}`),
			err: "name: Value is required and can't be empty (isEmpty); type: The input was not found in the haystack (notInArray)",
		},
		{
			name: "unknown error code",
			body: []byte(`{
				"errors": [
					{
						"error": {
							"key": "msg",
							"errors": [
								{
									"code": true,
									"message": "Test error"
								}
							]
						}
					}
				]
			}`),
			err: "msg: Test error (n/a)",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := handleErrorResponse(response, tt.body)

			verr, ok := err.(*ValidationErrorResponse)
			assert.True(t, ok)
			assert.NotNil(t, verr)
			assert.Equal(t, http.StatusBadRequest, verr.Status)
			assert.NotNil(t, verr.Response)
			assert.Equal(t, tt.err, verr.Error())
		})
	}
}

func handleErrorResponse(r *http.Response, body []byte) error {
	var errorResponse error
	if r.StatusCode == http.StatusBadRequest {
		errorResponse = &ValidationErrorResponse{Response: r, Status: r.StatusCode}
	} else {
		errorResponse = &ErrorResponse{Response: r}
	}

	if err := json.Unmarshal(body, errorResponse); err != nil {
		return err
	}

	return errorResponse
}
