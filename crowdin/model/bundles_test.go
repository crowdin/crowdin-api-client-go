package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBundleAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *BundleAddRequest
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
			req:  &BundleAddRequest{},
			err:  "name is required",
		},
		{
			name: "missing sourcePatterns",
			req:  &BundleAddRequest{Name: "Resx bundle"},
			err:  "sourcePatterns is required",
		},
		{
			name: "missing exportPattern",
			req:  &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"}},
			err:  "exportPattern is required",
		},
		{
			name:  "valid request",
			req:   &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"}, ExportPattern: "translations"},
			valid: true,
		},
		{
			name:  "valid request without format",
			req:   &BundleAddRequest{Name: "Original bundle", SourcePatterns: []string{"/master"}},
			valid: true,
		},
		{
			name: "labelMatchRule without labelIds",
			req: &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"},
				ExportPattern: "translations", LabelMatchRule: "any"},
			err: "labelMatchRule can only be used when labelIds is provided",
		},
		{
			name: "invalid labelMatchRule",
			req: &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"},
				ExportPattern: "translations", LabelIDs: []int{1}, LabelMatchRule: "some"},
			err: `invalid label match rule: "some"`,
		},
		{
			name: "excludeLabelMatchRule without excludeLabelIds",
			req: &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"},
				ExportPattern: "translations", ExcludeLabelMatchRule: "all"},
			err: "excludeLabelMatchRule can only be used when excludeLabelIds is provided",
		},
		{
			name: "invalid excludeLabelMatchRule",
			req: &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"},
				ExportPattern: "translations", ExcludeLabelIDs: []int{1}, ExcludeLabelMatchRule: "none"},
			err: `invalid label match rule: "none"`,
		},
		{
			name: "valid request with label match rules",
			req: &BundleAddRequest{Name: "Resx bundle", Format: "crowdin-resx", SourcePatterns: []string{"/master"},
				ExportPattern: "translations", LabelIDs: []int{1}, LabelMatchRule: "any",
				ExcludeLabelIDs: []int{2}, ExcludeLabelMatchRule: "all", LanguageIDs: []string{"uk"},
				IncludeInContextPseudoLanguage: toPtr(false), SourceLanguageExportPattern: "source.resx"},
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

func TestBundleExportRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *BundleExportRequest
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
			req:   &BundleExportRequest{},
			valid: true,
		},
		{
			name: "must not skip both",
			req:  &BundleExportRequest{SkipUntranslatedStrings: toPtr(true), SkipUntranslatedFiles: toPtr(true)},
			err:  "skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request",
		},
		{
			name: "must not use min approvals with passed workflow",
			req:  &BundleExportRequest{ExportWithMinApprovalsCount: toPtr(1), ExportStringsThatPassedWorkflow: toPtr(true)},
			err:  "exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request",
		},
		{
			name: "valid request",
			req: &BundleExportRequest{TargetLanguageIDs: []string{"uk"}, SkipUntranslatedStrings: toPtr(true),
				SkipUntranslatedFiles: toPtr(false), ExportApprovedOnly: toPtr(true),
				ExportWithMinApprovalsCount: toPtr(0), ExportStringsThatPassedWorkflow: toPtr(true)},
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
