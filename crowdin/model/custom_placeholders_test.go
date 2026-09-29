package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomPlaceholderAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *CustomPlaceholderAddRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &CustomPlaceholderAddRequest{},
			err:  "definition is required",
		},
		{
			name: "missing definition",
			req:  &CustomPlaceholderAddRequest{Description: "URL validation"},
			err:  "definition is required",
		},
		{
			name:  "required fields",
			req:   &CustomPlaceholderAddRequest{Definition: `start, then "http", end`},
			valid: true,
		},
		{
			name: "all fields",
			req: &CustomPlaceholderAddRequest{
				Definition:        `start, then "http", end`,
				Description:       "URL validation",
				ArgumentDelimiter: `"`,
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.valid {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
