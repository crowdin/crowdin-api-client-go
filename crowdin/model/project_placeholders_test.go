package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectPlaceholderAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ProjectPlaceholderAddRequest
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
			req:  &ProjectPlaceholderAddRequest{},
			err:  "customPlaceholderId is required",
		},
		{
			name: "invalid type",
			req:  &ProjectPlaceholderAddRequest{CustomPlaceholderID: 1, Type: "medium"},
			err:  `invalid type: "medium", must be one of high, low`,
		},
		{
			name:  "required fields",
			req:   &ProjectPlaceholderAddRequest{CustomPlaceholderID: 1},
			valid: true,
		},
		{
			name:  "type high",
			req:   &ProjectPlaceholderAddRequest{CustomPlaceholderID: 1, Type: ProjectPlaceholderTypeHigh},
			valid: true,
		},
		{
			name: "all fields",
			req: &ProjectPlaceholderAddRequest{
				CustomPlaceholderID: 3,
				Type:                ProjectPlaceholderTypeLow,
				Index:               1,
				IsBlocking:          toPtr(false),
				Formats:             []string{"docbook", "adoc"},
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
