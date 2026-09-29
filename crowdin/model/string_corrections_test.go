package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCorrectionsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *StringCorrectionsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &StringCorrectionsListOptions{},
		},
		{
			name: "string ID only",
			opts: &StringCorrectionsListOptions{StringID: 1},
			out:  "stringId=1",
		},
		{
			name: "invalid denormalize placeholders",
			opts: &StringCorrectionsListOptions{StringID: 1, DenormalizePlaceholders: toPtr(2)},
			out:  "stringId=1",
		},
		{
			name: "with all options",
			opts: &StringCorrectionsListOptions{
				StringID:                1,
				OrderBy:                 "createdAt desc,text",
				DenormalizePlaceholders: toPtr(0),
				ListOptions:             ListOptions{Limit: 10, Offset: 5},
			},
			out: "denormalizePlaceholders=0&limit=10&offset=5&orderBy=createdAt+desc%2Ctext&stringId=1",
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

func TestStringCorrectionAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *StringCorrectionAddRequest
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
			req:  &StringCorrectionAddRequest{},
			err:  "stringId is required",
		},
		{
			name: "missing text",
			req:  &StringCorrectionAddRequest{StringID: 1},
			err:  "text is required",
		},
		{
			name:  "valid request",
			req:   &StringCorrectionAddRequest{StringID: 1, Text: "Corrected"},
			valid: true,
		},
		{
			name:  "with plural category",
			req:   &StringCorrectionAddRequest{StringID: 1, Text: "Corrected", PluralCategoryName: "few"},
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
