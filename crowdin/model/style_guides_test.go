package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyleGuidesListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *StyleGuidesListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &StyleGuidesListOptions{},
		},
		{
			name: "with user ID",
			opts: &StyleGuidesListOptions{UserID: 2},
			out:  "userId=2",
		},
		{
			name: "with all options",
			opts: &StyleGuidesListOptions{
				OrderBy: "createdAt desc,name",
				UserID:  2,
				ListOptions: ListOptions{
					Limit:  10,
					Offset: 5,
				},
			},
			out: "limit=10&offset=5&orderBy=createdAt+desc%2Cname&userId=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, val.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, val)
			}
		})
	}
}

func TestStyleGuideCreateRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *StyleGuideCreateRequest
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
			req:  &StyleGuideCreateRequest{},
			err:  "name is required",
		},
		{
			name:  "name only",
			req:   &StyleGuideCreateRequest{Name: "Guide"},
			valid: true,
		},
		{
			name: "all fields",
			req: &StyleGuideCreateRequest{
				Name:           "Guide",
				StorageID:      toPtr(1),
				GroupID:        toPtr(0),
				AIInstructions: "Be concise",
				LanguageIDs:    []string{"uk"},
				ProjectIDs:     []int{1},
				IsShared:       toPtr(true),
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
