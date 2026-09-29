package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreTranslationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PreTranslationRequest
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
			req:  &PreTranslationRequest{},
			err:  "languageIds is required",
		},
		{
			name:  "fileIds are not required",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}},
			valid: true,
		},
		{
			name:  "valid request with directoryIds",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}, DirectoryIDs: []int{1}},
			valid: true,
		},
		{
			name:  "valid request with branchIds",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}, BranchIDs: []int{1}},
			valid: true,
		},
		{
			name:  "valid request by task without languageIds",
			req:   &PreTranslationRequest{TaskID: 5},
			valid: true,
		},
		{
			name: "valid request",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, FileIDs: []int{1, 2},
				Method: "tm", EngineID: 1, AutoApproveOption: "all", DuplicateTranslations: toPtr(false)},
			valid: true,
		},
		{
			name: "missing aiPromptId",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, FileIDs: []int{1, 2},
				Method: "ai", AutoApproveOption: "all", DuplicateTranslations: toPtr(false)},
			err: "aiPromptId is required",
		},
		{
			name: "valid request with ai method",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, FileIDs: []int{1, 2},
				Method: "ai", AIPromptID: 1, AutoApproveOption: "all", DuplicateTranslations: toPtr(false)},
			valid: true,
		},
		{
			name: "missing engineId with mt method",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, FileIDs: []int{1, 2},
				Method: "mt", AutoApproveOption: "all", DuplicateTranslations: toPtr(false)},
			err: "engineId is required",
		},
		{
			name: "valid request with mt method",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, FileIDs: []int{1, 2},
				Method: "mt", EngineID: 1, AutoApproveOption: "all", DuplicateTranslations: toPtr(false)},
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

func TestBuildProjectDirectoryTranslationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *BuildProjectDirectoryTranslationRequest
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
			req:   &BuildProjectDirectoryTranslationRequest{},
			valid: true,
		},
		{
			name: "must not skip both",
			req: &BuildProjectDirectoryTranslationRequest{SkipUntranslatedStrings: toPtr(true),
				SkipUntranslatedFiles: toPtr(true)},
			err: "skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request",
		},
		{
			name: "must not export both",
			req: &BuildProjectDirectoryTranslationRequest{
				ExportWithMinApprovalsCount:     toPtr(1),
				ExportStringsThatPassedWorkflow: toPtr(true),
			},
			err: "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request",
		},
		{
			name: "valid request",
			req: &BuildProjectDirectoryTranslationRequest{TargetLanguageIDs: []string{"en"},
				SkipUntranslatedStrings: toPtr(true), ExportApprovedOnly: toPtr(true)},
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

func TestBuildProjectFileTranslationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *BuildProjectFileTranslationRequest
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
			req:  &BuildProjectFileTranslationRequest{},
			err:  "targetLanguageId is required",
		},
		{
			name: "must not skip both",
			req: &BuildProjectFileTranslationRequest{TargetLanguageID: "uk", SkipUntranslatedStrings: toPtr(true),
				SkipUntranslatedFiles: toPtr(true)},
			err: "skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request",
		},
		{
			name: "must not export both",
			req: &BuildProjectFileTranslationRequest{TargetLanguageID: "uk", ExportWithMinApprovalsCount: toPtr(1),
				ExportStringsThatPassedWorkflow: toPtr(true)},
			err: "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request",
		},
		{
			name: "valid request",
			req: &BuildProjectFileTranslationRequest{TargetLanguageID: "uk", SkipUntranslatedStrings: toPtr(true),
				ExportApprovedOnly: toPtr(true)},
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

func TestTranslationsBuildsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *TranslationsBuildsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &TranslationsBuildsListOptions{},
		},
		{
			name: "with all options",
			opts: &TranslationsBuildsListOptions{BranchID: 1, ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out:  "branchId=1&limit=10&offset=5",
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

func TestBuildProjectRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *BuildProjectRequest
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
			req:   &BuildProjectRequest{},
			valid: true,
		},
		{
			name: "must not skip both",
			req:  &BuildProjectRequest{SkipUntranslatedStrings: toPtr(true), SkipUntranslatedFiles: toPtr(true)},
			err:  "`skipUntranslatedStrings` and `skipUntranslatedFiles` must not be true at the same request",
		},
		{
			name: "must not export both",
			req:  &BuildProjectRequest{ExportWithMinApprovalsCount: toPtr(1), ExportStringsThatPassedWorkflow: toPtr(true)},
			err:  "`exportWithMinApprovalsCount` and `exportStringsThatPassedWorkflow` must not be true at the same request",
		},
		{
			name: "valid request",
			req: &BuildProjectRequest{BranchID: 1, TargetLanguageIDs: []string{"en"}, SkipUntranslatedStrings: toPtr(true),
				ExportApprovedOnly: toPtr(true)},
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

func TestPseudoBuildProjectRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PseudoBuildProjectRequest
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
			req:   &PseudoBuildProjectRequest{},
			valid: true,
		},
		{
			name: "invalid lengthTransformation",
			req:  &PseudoBuildProjectRequest{LengthTransformation: toPtr(-100)},
			err:  "lengthTransformation must be from -50 to 100",
		},
		{
			name: "valid request",
			req: &PseudoBuildProjectRequest{Pseudo: toPtr(true), BranchID: 1, Prefix: "prefix", Suffix: "suffix",
				LengthTransformation: toPtr(50), CharTransformation: "cyrillic"},
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

func TestUploadTranslationsRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *UploadTranslationsRequest
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
			req:  &UploadTranslationsRequest{},
			err:  "storageId is required",
		},
		{
			name: "one of fileId or branchId is required",
			req:  &UploadTranslationsRequest{StorageID: 1, FileID: 2, BranchID: 3},
			err:  "fileId and branchId can not be used at the same request",
		},
		{
			name: "valid request",
			req: &UploadTranslationsRequest{StorageID: 1, FileID: 2, BranchID: 0,
				ImportEqSuggestions: toPtr(true), AutoApproveImported: toPtr(true), TranslateHidden: toPtr(true),
				AddToTM: toPtr(false)},
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

func TestExportTranslationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ExportTranslationRequest
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
			req:  &ExportTranslationRequest{},
			err:  "targetLanguageId is required",
		},
		{
			name: "valid request",
			req: &ExportTranslationRequest{TargetLanguageID: "uk", Format: "xliff", FileIDs: []int{1}, LabelIDs: []int{4},
				SkipUntranslatedFiles: toPtr(false), ExportApprovedOnly: toPtr(false)},
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

func TestPreTranslationRequestValidate_Scope(t *testing.T) {
	tests := []struct {
		name  string
		req   *PreTranslationRequest
		err   string
		valid bool
	}{
		{
			name:  "valid scope untranslated",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeUntranslated},
			valid: true,
		},
		{
			name: "invalid scope",
			req:  &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: "some"},
			err:  `invalid scope: "some"`,
		},
		{
			name: "scope with deprecated translateUntranslatedOnly",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeAll,
				TranslateUntranslatedOnly: toPtr(false)},
			err: "scope and translateUntranslatedOnly cannot be used together",
		},
		{
			name: "invalid replaceTranslationsOption",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeAll,
				ReplaceTranslationsOption: "some"},
			err: `invalid replaceTranslationsOption: "some"`,
		},
		{
			name: "replaceTranslationsOption requires re-translation scope",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"},
				ReplaceTranslationsOption: ReplaceTranslationsOptionAutoTranslated},
			err: "replaceTranslationsOption requires scope to be translated or all",
		},
		{
			name: "replaceTranslationsOption none without scope",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"},
				ReplaceTranslationsOption: ReplaceTranslationsOptionNone},
			valid: true,
		},
		{
			name: "replaceTranslationsOption with deprecated translateUntranslatedOnly false",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, TranslateUntranslatedOnly: toPtr(false),
				ReplaceTranslationsOption: ReplaceTranslationsOptionAll},
			valid: true,
		},
		{
			name: "replaceTranslationsOption with deprecated translateUntranslatedOnly true",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, TranslateUntranslatedOnly: toPtr(true),
				ReplaceTranslationsOption: ReplaceTranslationsOptionAll},
			err: "replaceTranslationsOption requires scope to be translated or all",
		},
		{
			name: "translationModifiedBefore requires re-translation scope",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeUntranslated,
				TranslationModifiedBefore: "2026-01-01T00:00:00+00:00"},
			err: "translationModifiedBefore and translationModifiedAfter require scope to be translated or all",
		},
		{
			name: "translationModifiedAfter requires re-translation scope",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"},
				TranslationModifiedAfter: "2026-01-01T00:00:00+00:00"},
			err: "translationModifiedBefore and translationModifiedAfter require scope to be translated or all",
		},
		{
			name: "translationModifiedAfter later than translationModifiedBefore",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeTranslated,
				TranslationModifiedBefore: "2026-01-01T00:00:00+00:00",
				TranslationModifiedAfter:  "2026-02-01T00:00:00+00:00"},
			err: "translationModifiedAfter cannot be later than translationModifiedBefore",
		},
		{
			name: "valid translation modified date range",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeAll,
				TranslationModifiedBefore: "2026-02-01T00:00:00+00:00",
				TranslationModifiedAfter:  "2026-01-01T00:00:00+00:00"},
			valid: true,
		},
		{
			name: "duplicateTranslations with replace all",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeAll,
				ReplaceTranslationsOption: ReplaceTranslationsOptionAll, DuplicateTranslations: toPtr(true)},
			err: "duplicateTranslations cannot be used when replaceTranslationsOption is all",
		},
		{
			name: "resetApprovalStatus requires re-translation scope",
			req:  &PreTranslationRequest{LanguageIDs: []string{"uk"}, ResetApprovalStatus: toPtr(true)},
			err:  "resetApprovalStatus requires scope to be translated or all",
		},
		{
			name: "resetApprovalStatus with skipApprovedTranslations",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeTranslated,
				ResetApprovalStatus: toPtr(true), SkipApprovedTranslations: toPtr(true)},
			err: "resetApprovalStatus cannot be used together with skipApprovedTranslations",
		},
		{
			name: "resetApprovalStatus with autoApproveOption",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeTranslated,
				ResetApprovalStatus: toPtr(true), AutoApproveOption: "all"},
			err: "resetApprovalStatus cannot be used together with autoApproveOption",
		},
		{
			name: "resetApprovalStatus with replace all",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeTranslated,
				ResetApprovalStatus: toPtr(true), ReplaceTranslationsOption: ReplaceTranslationsOptionAll},
			err: "resetApprovalStatus cannot be used when replaceTranslationsOption is all",
		},
		{
			name: "valid resetApprovalStatus",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Scope: PreTranslationScopeTranslated,
				ResetApprovalStatus: toPtr(true), AutoApproveOption: "none",
				ReplaceTranslationsOption: ReplaceTranslationsOptionAutoTranslated, SkipApprovedTranslations: toPtr(false)},
			valid: true,
		},
		{
			name:  "resetApprovalStatus false is ignored",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}, ResetApprovalStatus: toPtr(false)},
			valid: true,
		},
		{
			name: "minimumMatchRatio too low",
			req:  &PreTranslationRequest{LanguageIDs: []string{"uk"}, MinimumMatchRatio: 39},
			err:  "minimumMatchRatio must be from 40 to 100",
		},
		{
			name: "minimumMatchRatio too high",
			req:  &PreTranslationRequest{LanguageIDs: []string{"uk"}, MinimumMatchRatio: 101},
			err:  "minimumMatchRatio must be from 40 to 100",
		},
		{
			name:  "valid minimumMatchRatio",
			req:   &PreTranslationRequest{LanguageIDs: []string{"uk"}, MinimumMatchRatio: 40},
			valid: true,
		},
		{
			name: "customInstruction too long",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Method: "ai", AIPromptID: 1,
				CustomInstruction: strings.Repeat("a", 10001)},
			err: "customInstruction must not exceed 10000 characters",
		},
		{
			name: "valid customInstruction",
			req: &PreTranslationRequest{LanguageIDs: []string{"uk"}, Method: "ai", AIPromptID: 1,
				CustomInstruction: strings.Repeat("a", 10000), NotifyOnCompletion: toPtr(true), SourceLanguageID: "en"},
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

func TestPreTranslationsListOptionsValues(t *testing.T) {
	tests := []struct {
		name     string
		opts     *PreTranslationsListOptions
		expected string
		isSet    bool
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &PreTranslationsListOptions{},
		},
		{
			name:     "with orderBy",
			opts:     &PreTranslationsListOptions{OrderBy: "createdAt desc"},
			expected: "orderBy=createdAt+desc",
			isSet:    true,
		},
		{
			name: "with all options",
			opts: &PreTranslationsListOptions{OrderBy: "createdAt desc",
				ListOptions: ListOptions{Limit: 10, Offset: 5}},
			expected: "limit=10&offset=5&orderBy=createdAt+desc",
			isSet:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opts.Values()
			assert.Equal(t, tt.isSet, ok)
			assert.Equal(t, tt.expected, v.Encode())
		})
	}
}

func TestTranslationImportRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *TranslationImportRequest
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
			req:  &TranslationImportRequest{},
			err:  "storageId is required",
		},
		{
			name: "fileId and branchId at the same request",
			req:  &TranslationImportRequest{StorageID: 1, FileID: 2, BranchID: 3},
			err:  "fileId and branchId can not be used at the same request",
		},
		{
			name:  "valid request with storageId only",
			req:   &TranslationImportRequest{StorageID: 1},
			valid: true,
		},
		{
			name: "valid file-based request",
			req: &TranslationImportRequest{StorageID: 1, FileID: 2, LanguageIDs: []string{"uk"},
				ImportEqSuggestions: toPtr(true), AutoApproveImported: toPtr(true), TranslateHidden: toPtr(false),
				AddToTM: toPtr(false)},
			valid: true,
		},
		{
			name: "valid string-based request",
			req: &TranslationImportRequest{StorageID: 1, BranchID: 3,
				ImportOptions: &TranslationImportOptions{Scheme: map[string]int{"identifier": 0, "uk": 1}}},
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
