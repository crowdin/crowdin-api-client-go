package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectProgressListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *ProjectProgressListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &ProjectProgressListOptions{},
		},
		{
			name: "with all options",
			opts: &ProjectProgressListOptions{LanguageIDs: []string{"uk", "fr"},
				ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "languageIds=uk%2Cfr&limit=10&offset=5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v.Encode())
			}
		})
	}
}

func TestQACheckListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *QACheckListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &QACheckListOptions{},
		},
		{
			name: "with all options",
			opts: &QACheckListOptions{Category: []string{"variables", "tags"},
				Validation:  []string{"spellcheck", "escaped_quotes_check", "multiple_spaces_check"},
				LanguageIDs: []string{"uk", "fr"}},
			out: "category=variables%2Ctags&languageIds=uk%2Cfr&validation=spellcheck%2Cescaped_quotes_check%2Cmultiple_spaces_check",
		},
		{
			name: "with taskId and fileId",
			opts: &QACheckListOptions{TaskID: 5, FileID: 7, ListOptions: ListOptions{Limit: 10}},
			out:  "fileId=7&limit=10&taskId=5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v.Encode())
			}
		})
	}
}

func TestQACheckRevalidationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *QACheckRevalidationRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name:  "empty request",
			req:   &QACheckRevalidationRequest{},
			valid: true,
		},
		{
			name: "valid request",
			req: &QACheckRevalidationRequest{QACheckCategories: []string{"ai"}, LanguageIDs: []string{"uk"},
				FailedOnly: toPtr(true), ExternalQACheckIDs: []int{1}},
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

func TestQACheckValidateRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *QACheckValidateRequest
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
			req:  &QACheckValidateRequest{},
			err:  "stringId is required",
		},
		{
			name: "missing languageId",
			req:  &QACheckValidateRequest{StringID: 1},
			err:  "languageId is required",
		},
		{
			name: "missing text",
			req:  &QACheckValidateRequest{StringID: 1, LanguageID: "uk"},
			err:  "text is required",
		},
		{
			name:  "valid request",
			req:   &QACheckValidateRequest{StringID: 1, LanguageID: "uk", Text: "Text", PluralCategoryName: "few"},
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
