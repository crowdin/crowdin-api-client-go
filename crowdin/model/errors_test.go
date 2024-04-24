package model

import (
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
					Key    string `json:"key"`
					Errors []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					} `json:"errors"`
				}{
					Key: "name",
					Errors: []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					}{
						{
							Code:    "isEmpty",
							Message: "Value is required and can't be empty",
						},
					},
				},
			},
			{
				Error: struct {
					Key    string `json:"key"`
					Errors []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					} `json:"errors"`
				}{
					Key: "sourceLanguage",
					Errors: []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					}{
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
					Key    string `json:"key"`
					Errors []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					} `json:"errors"`
				}{
					Key: "name",
					Errors: []struct {
						Code    string `json:"code"`
						Message string `json:"message"`
					}{
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
