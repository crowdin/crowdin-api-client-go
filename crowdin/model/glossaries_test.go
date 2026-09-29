package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConceptsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *ConceptsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &ConceptsListOptions{},
		},
		{
			name: "with options",
			opts: &ConceptsListOptions{
				OrderBy: "name",
				ListOptions: ListOptions{
					Limit:  10,
					Offset: 5,
				},
			},
			out: "limit=10&offset=5&orderBy=name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, ok := tt.opts.Values()
			if tt.opts == nil || len(tt.out) == 0 {
				assert.False(t, ok)
			} else {
				assert.True(t, ok)
				assert.Equal(t, tt.out, opts.Encode())
			}
		})
	}
}

func TestGlossariesListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *GlossariesListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &GlossariesListOptions{},
		},
		{
			name: "with groupId = 0",
			opts: &GlossariesListOptions{GroupID: toPtr(0)},
			out:  "groupId=0",
		},
		{
			name: "with userId",
			opts: &GlossariesListOptions{UserID: 1},
			out:  "userId=1",
		},
		{
			name: "with filter",
			opts: &GlossariesListOptions{Filter: "Glossary"},
			out:  "filter=Glossary",
		},
		{
			name: "with all options",
			opts: &GlossariesListOptions{OrderBy: "name", GroupID: toPtr(1),
				ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "groupId=1&limit=10&offset=5&orderBy=name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, opts.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, opts)
			}
		})
	}
}
func TestTermsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *TermsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &TermsListOptions{},
		},
		{
			name: "with all options",
			opts: &TermsListOptions{
				OrderBy:     "name",
				UserID:      1,
				LanguageID:  "en",
				ConceptID:   2,
				CroQL:       "croql",
				ListOptions: ListOptions{Limit: 10, Offset: 5},
			},
			out: "conceptId=2&croql=croql&languageId=en&limit=10&offset=5&orderBy=name&userId=1",
		},
		{
			name: "with translationOfTermId",
			opts: &TermsListOptions{TranslationOfTermID: 3},
			out:  "translationOfTermId=3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, ok := tt.opts.Values()
			if tt.opts == nil || len(tt.out) == 0 {
				assert.False(t, ok)
			} else {
				assert.True(t, ok)
				assert.Equal(t, tt.out, opts.Encode())
			}
		})
	}
}

func TestClearGlossaryOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *ClearGlossaryOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &ClearGlossaryOptions{},
		},
		{
			name: "with all options",
			opts: &ClearGlossaryOptions{
				LanguageID:          "en",
				ConceptID:           2,
				TranslationOfTermID: 3,
			},
			out: "conceptId=2&languageId=en&translationOfTermId=3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, opts.Encode())
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestConceptUpdateRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ConceptUpdateRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name:  "valid empty request",
			req:   &ConceptUpdateRequest{},
			valid: true,
		},
		{
			name:  "valid request",
			req:   &ConceptUpdateRequest{Subject: "subject"},
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

func TestGlossaryAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *GlossaryAddRequest
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
			req:  &GlossaryAddRequest{},
			err:  "name is required",
		},
		{
			name: "required fields missing",
			req:  &GlossaryAddRequest{Name: "glossary"},
			err:  "languageId is required",
		},
		{
			name:  "valid request",
			req:   &GlossaryAddRequest{Name: "glossary", LanguageID: "en"},
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

func TestGlossaryExportRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *GlossaryExportRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name:  "valid empty request",
			req:   &GlossaryExportRequest{},
			valid: true,
		},
		{
			name:  "valid request",
			req:   &GlossaryExportRequest{Format: "xlsx"},
			valid: true,
		},
		{
			name: "valid request with filters",
			req: &GlossaryExportRequest{
				Format:          "csv",
				ExportType:      "terms",
				Text:            "value",
				CaseSensitive:   toPtr(true),
				SearchStrict:    toPtr(false),
				SearchFullMatch: toPtr(true),
				Statuses:        []string{"PREFERRED", "ADMITTED"},
				PartsOfSpeech:   []string{"NOUN", "VERB"},
				Types:           []string{"ACRONYM"},
				Genders:         []string{"MASCULINE"},
				AuthorIDs:       []int{12, 34},
				LanguageIDs:     []string{"uk", "de"},
				DateFrom:        "2024-01-23T07:00:14+00:00",
				DateTo:          "2024-09-27T07:00:14+00:00",
			},
			valid: true,
		},
		{
			name:  "valid exportType concepts",
			req:   &GlossaryExportRequest{ExportType: "concepts"},
			valid: true,
		},
		{
			name: "invalid exportType",
			req:  &GlossaryExportRequest{ExportType: "all"},
			err:  `invalid exportType: "all", must be one of concepts, terms`,
		},
		{
			name: "caseSensitive without text",
			req:  &GlossaryExportRequest{CaseSensitive: toPtr(true)},
			err:  "caseSensitive, searchStrict and searchFullMatch must be used together with text",
		},
		{
			name: "searchStrict without text",
			req:  &GlossaryExportRequest{SearchStrict: toPtr(true)},
			err:  "caseSensitive, searchStrict and searchFullMatch must be used together with text",
		},
		{
			name: "searchFullMatch without text",
			req:  &GlossaryExportRequest{SearchFullMatch: toPtr(false)},
			err:  "caseSensitive, searchStrict and searchFullMatch must be used together with text",
		},
		{
			name: "too many authorIds",
			req:  &GlossaryExportRequest{AuthorIDs: make([]int, 51)},
			err:  "authorIds cannot contain more than 50 values",
		},
		{
			name: "too many languageIds",
			req:  &GlossaryExportRequest{LanguageIDs: make([]string, 51)},
			err:  "languageIds cannot contain more than 50 values",
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

func TestGlossaryImportRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *GlossaryImportRequest
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
			req:  &GlossaryImportRequest{},
			err:  "storageId is required",
		},
		{
			name: "required fields missing",
			req:  &GlossaryImportRequest{StorageID: -1},
			err:  "storageId is required",
		},
		{
			name:  "valid request",
			req:   &GlossaryImportRequest{StorageID: 1, Scheme: map[string]int{"term_en": 1}},
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

func TestGlossaryConcordanceSearchRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *GlossaryConcordanceSearchRequest
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
			req:  &GlossaryConcordanceSearchRequest{},
			err:  "sourceLanguageId is required",
		},
		{
			name: "required fields missing (targetLanguageId)",
			req:  &GlossaryConcordanceSearchRequest{SourceLanguageID: "en"},
			err:  "targetLanguageId is required",
		},
		{
			name: "required fields missing (expressions)",
			req:  &GlossaryConcordanceSearchRequest{SourceLanguageID: "en", TargetLanguageID: "de"},
			err:  "expressions cannot be empty",
		},
		{
			req: &GlossaryConcordanceSearchRequest{
				SourceLanguageID: "en",
				TargetLanguageID: "de",
				Expressions:      []string{},
			},
			err: "expressions cannot be empty",
		},
		{
			req: &GlossaryConcordanceSearchRequest{
				SourceLanguageID: "en",
				TargetLanguageID: "de",
				Expressions:      []string{"term"},
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

func TestTermAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *TermAddRequest
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
			req:  &TermAddRequest{},
			err:  "languageId is required",
		},
		{
			name: "required fields missing",
			req:  &TermAddRequest{LanguageID: "fr"},
			err:  "text is required",
		},
		{
			name:  "valid request",
			req:   &TermAddRequest{LanguageID: "fr", Text: "term"},
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

func TestGlossaryConcordanceSearchAllRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *GlossaryConcordanceSearchAllRequest
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
			req:  &GlossaryConcordanceSearchAllRequest{},
			err:  "sourceLanguageId is required",
		},
		{
			name: "required fields missing (targetLanguageId)",
			req: &GlossaryConcordanceSearchAllRequest{
				GlossaryConcordanceSearchRequest: GlossaryConcordanceSearchRequest{SourceLanguageID: "en"},
			},
			err: "targetLanguageId is required",
		},
		{
			name: "required fields missing (expressions)",
			req: &GlossaryConcordanceSearchAllRequest{
				GlossaryConcordanceSearchRequest: GlossaryConcordanceSearchRequest{
					SourceLanguageID: "en",
					TargetLanguageID: "de",
				},
				UserID: 1,
			},
			err: "expressions cannot be empty",
		},
		{
			name: "valid request",
			req: &GlossaryConcordanceSearchAllRequest{
				GlossaryConcordanceSearchRequest: GlossaryConcordanceSearchRequest{
					SourceLanguageID: "en",
					TargetLanguageID: "de",
					Expressions:      []string{"term"},
				},
			},
			valid: true,
		},
		{
			name: "valid request with userId",
			req: &GlossaryConcordanceSearchAllRequest{
				GlossaryConcordanceSearchRequest: GlossaryConcordanceSearchRequest{
					SourceLanguageID: "en",
					TargetLanguageID: "de",
					Expressions:      []string{"term"},
				},
				UserID: 1,
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
