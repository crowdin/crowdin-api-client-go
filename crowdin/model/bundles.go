package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Bundle represents a Crowdin bundle.
type Bundle struct {
	ID                           int      `json:"id"`
	Name                         string   `json:"name"`
	Format                       string   `json:"format"`
	SourcePatterns               []string `json:"sourcePatterns"`
	IgnorePatterns               []string `json:"ignorePatterns"`
	ExportPattern                string   `json:"exportPattern"`
	IsMultilingual               bool     `json:"isMultilingual"`
	IncludeProjectSourceLanguage bool     `json:"includeProjectSourceLanguage"`
	LabelIDs                     []int    `json:"labelIds"`
	ExcludeLabelIDs              []int    `json:"excludeLabelIds"`
	WebURL                       string   `json:"webUrl"`
	CreatedAt                    string   `json:"createdAt"`
	UpdatedAt                    string   `json:"updatedAt"`

	IncludeInContextPseudoLanguage bool     `json:"includeInContextPseudoLanguage"`
	SourceLanguageExportPattern    *string  `json:"sourceLanguageExportPattern,omitempty"`
	LabelMatchRule                 *string  `json:"labelMatchRule,omitempty"`
	ExcludeLabelMatchRule          *string  `json:"excludeLabelMatchRule,omitempty"`
	LanguageIDs                    []string `json:"languageIds,omitempty"`
}

// BundleResponse defines the structure of a response
// when getting a single bundle.
type BundleResponse struct {
	Data *Bundle `json:"data"`
}

// BundlesListResponse defines the structure of a response
// when getting a list of bundles.
type BundlesListResponse struct {
	Data []*BundleResponse `json:"data"`
}

// BundleAddRequest defines the structure of a request
// to add a new bundle.
type BundleAddRequest struct {
	// Defines name.
	Name string `json:"name"`
	// Defines export file format. If not provided, files will be exported
	// in their original format.
	// Note: Required for string-based projects.
	Format string `json:"format,omitempty"`
	// Source patterns.
	SourcePatterns []string `json:"sourcePatterns"`
	// Ignore patterns.
	IgnorePatterns []string `json:"ignorePatterns"`
	// Bundle export pattern. Defines bundle name in resulting
	// translations bundle. Required if `format` is specified.
	// Note: Can't contain : * ? " < > | symbols.
	ExportPattern string `json:"exportPattern,omitempty"`
	// Export translations in multilingual file.
	// Default: false.
	IsMultilingual *bool `json:"isMultilingual"`
	// Add project source language to bundle.
	// Default: false.
	IncludeProjectSourceLanguage *bool `json:"includeProjectSourceLanguage"`
	// Label Identifiers.
	LabelIDs []int `json:"labelIds"`
	// Label Identifiers.
	ExcludeLabelIDs []int `json:"excludeLabelIds"`
	// Source language export pattern.
	// Note: Works only when `includeProjectSourceLanguage` is enabled.
	// Can't contain : * ? " < > | symbols.
	SourceLanguageExportPattern string `json:"sourceLanguageExportPattern,omitempty"`
	// Add In-Context pseudo-language to bundle. Default: true.
	IncludeInContextPseudoLanguage *bool `json:"includeInContextPseudoLanguage,omitempty"`
	// Match rule for labels. Enum: all, any. Default: all.
	// Note: Can only be used when `labelIds` parameter is provided.
	LabelMatchRule string `json:"labelMatchRule,omitempty"`
	// Match rule for excluded labels. Enum: all, any. Default: all.
	// Note: Can only be used when `excludeLabelIds` parameter is provided.
	ExcludeLabelMatchRule string `json:"excludeLabelMatchRule,omitempty"`
	// Language Identifiers. If provided, bundle will only export specified languages.
	// If not provided, bundle will export all project target languages.
	LanguageIDs []string `json:"languageIds,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (b *BundleAddRequest) Validate() error {
	if b == nil {
		return ErrNilRequest
	}
	if b.Name == "" {
		return errors.New("name is required")
	}
	if b.Format == "" {
		return errors.New("format is required")
	}
	if len(b.SourcePatterns) == 0 {
		return errors.New("sourcePatterns is required")
	}
	if b.Format != "" && b.ExportPattern == "" {
		return errors.New("exportPattern is required")
	}
	if b.LabelMatchRule != "" {
		if len(b.LabelIDs) == 0 {
			return errors.New("labelMatchRule can only be used when labelIds is provided")
		}
		if err := validateBundleLabelMatchRule(b.LabelMatchRule); err != nil {
			return err
		}
	}
	if b.ExcludeLabelMatchRule != "" {
		if len(b.ExcludeLabelIDs) == 0 {
			return errors.New("excludeLabelMatchRule can only be used when excludeLabelIds is provided")
		}
		if err := validateBundleLabelMatchRule(b.ExcludeLabelMatchRule); err != nil {
			return err
		}
	}

	return nil
}

// validateBundleLabelMatchRule checks if the label match rule is one of: all, any.
func validateBundleLabelMatchRule(rule string) error {
	if rule != "all" && rule != "any" {
		return fmt.Errorf("invalid label match rule: %q", rule)
	}

	return nil
}

// BundleExport represents a bundle export progress.
type BundleExport struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
	// Attributes holds only the bundle identifier and is kept for backward
	// compatibility. Use ExportAttributes to access all export attributes.
	Attributes struct {
		BundleID int `json:"bundleId"`
	} `json:"attributes"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	// StartedAt is nullable in the API; a null value is decoded as an empty string.
	StartedAt string `json:"startedAt"`
	// FinishedAt is nullable in the API; a null value is decoded as an empty string.
	FinishedAt string `json:"finishedAt"`

	// ExportAttributes holds all attributes of the bundle export
	// (decoded from the `attributes` object of the response).
	ExportAttributes *BundleExportAttributes `json:"-"`
}

// UnmarshalJSON decodes a bundle export, populating both
// Attributes and ExportAttributes from the `attributes` object.
func (b *BundleExport) UnmarshalJSON(data []byte) error {
	type bundleExport BundleExport
	if err := json.Unmarshal(data, (*bundleExport)(b)); err != nil {
		return err
	}

	var aux struct {
		Attributes *BundleExportAttributes `json:"attributes"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	b.ExportAttributes = aux.Attributes

	return nil
}

// BundleExportAttributes represents the attributes of a bundle export.
type BundleExportAttributes struct {
	BundleID                        int      `json:"bundleId"`
	TargetLanguageIDs               []string `json:"targetLanguageIds,omitempty"`
	SkipUntranslatedStrings         *bool    `json:"skipUntranslatedStrings,omitempty"`
	SkipUntranslatedFiles           *bool    `json:"skipUntranslatedFiles,omitempty"`
	ExportApprovedOnly              *bool    `json:"exportApprovedOnly,omitempty"`
	ExportWithMinApprovalsCount     *int     `json:"exportWithMinApprovalsCount,omitempty"`
	ExportStringsThatPassedWorkflow *bool    `json:"exportStringsThatPassedWorkflow,omitempty"`
}

// BundleExportRequest defines the structure of a request to export a bundle.
type BundleExportRequest struct {
	// Specify target languages for export.
	// Leave this field empty to export all target languages.
	TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
	// Defines whether to export only translated strings.
	// Note: true value can't be used with `skipUntranslatedFiles=true` in the same request.
	SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated files.
	// Note: true value can't be used with `skipUntranslatedStrings=true` in the same request.
	// Available only for file-based projects.
	SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default: false.
	// Note: Available only for Crowdin.
	ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`
	// Defines whether to export only strings with a minimum number of approvals.
	// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
	// in the same request. Available only for Crowdin Enterprise.
	ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
	// Defines whether to export only strings that passed workflow.
	// Note: true value can't be used with `exportWithMinApprovalsCount>0` in the same
	// request. Available only for Crowdin Enterprise.
	ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BundleExportRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.SkipUntranslatedStrings != nil && *r.SkipUntranslatedStrings &&
		r.SkipUntranslatedFiles != nil && *r.SkipUntranslatedFiles {
		return errors.New("skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request")
	}
	if r.ExportWithMinApprovalsCount != nil && *r.ExportWithMinApprovalsCount > 0 &&
		r.ExportStringsThatPassedWorkflow != nil && *r.ExportStringsThatPassedWorkflow {
		return errors.New("exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request")
	}

	return nil
}

// BundleExportResponse defines the structure of a response
// when exporting a bundle.
type BundleExportResponse struct {
	Data *BundleExport `json:"data"`
}
