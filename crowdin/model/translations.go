package model

import (
	"errors"
	"fmt"
	"net/url"
	"time"
	"unicode/utf8"
)

type (
	// PreTranslation represents a pre-translation status.
	PreTranslation struct {
		Identifier string                    `json:"identifier"`
		Status     string                    `json:"status"`
		Progress   int                       `json:"progress"`
		Attributes *PreTranslationAttributes `json:"attributes"`
		CreatedAt  string                    `json:"createdAt"`
		UpdatedAt  string                    `json:"updatedAt"`
		StartedAt  string                    `json:"startedAt,omitempty"`
		FinishedAt string                    `json:"finishedAt,omitempty"`
	}

	PreTranslationAttributes struct {
		LanguageIDs              []string `json:"languageIds"`
		BranchIDs                []int    `json:"branchIds,omitempty"`
		FileIDs                  []int    `json:"fileIds,omitempty"`
		DirectoryIDs             []int    `json:"directoryIds,omitempty"`
		TaskID                   *int     `json:"taskId,omitempty"`
		Method                   *string  `json:"method,omitempty"`
		AutoApproveOption        *string  `json:"autoApproveOption,omitempty"`
		DuplicateTranslations    *bool    `json:"duplicateTranslations,omitempty"`
		SkipApprovedTranslations *bool    `json:"skipApprovedTranslations,omitempty"`
		// Deprecated: use Scope instead.
		TranslateUntranslatedOnly     *bool   `json:"translateUntranslatedOnly,omitempty"`
		Scope                         *string `json:"scope,omitempty"`
		TranslationModifiedBefore     *string `json:"translationModifiedBefore,omitempty"`
		TranslationModifiedAfter      *string `json:"translationModifiedAfter,omitempty"`
		ReplaceTranslationsOption     *string `json:"replaceTranslationsOption,omitempty"`
		ResetApprovalStatus           *bool   `json:"resetApprovalStatus,omitempty"`
		TranslateWithPerfectMatchOnly *bool   `json:"translateWithPerfectMatchOnly,omitempty"`
		MinimumMatchRatio             *int    `json:"minimumMatchRatio,omitempty"`
		Priority                      *string `json:"priority,omitempty"`
		NotifyOnCompletion            *bool   `json:"notifyOnCompletion,omitempty"`
	}

	PreTranslationReport struct {
		Languages        []*LanguageReport `json:"languages"`
		PreTranslateType string            `json:"preTranslateType"`
	}

	LanguageReport struct {
		ID                       string                                  `json:"id"`
		Files                    []*LanguageReportFile                   `json:"files"`
		Skipped                  *LanguageReportSkipped                  `json:"skipped,omitempty"`
		SkippedQaCheckCategories *LanguageReportSkippedQaCheckCategories `json:"skippedQaCheckCategories,omitempty"`
	}

	LanguageReportFile struct {
		ID         string                    `json:"id"`
		Statistics *LanguageReportStatistics `json:"statistics"`
	}

	LanguageReportStatistics struct {
		Phrases int `json:"phrases"`
		Words   int `json:"words"`
	}

	LanguageReportSkipped struct {
		TranslationEQSource int `json:"translation_eq_source"`
		QACheck             int `json:"qa_check"`
		HiddenStrings       int `json:"hidden_strings"`
		AIError             int `json:"ai_error"`
	}

	LanguageReportSkippedQaCheckCategories struct {
		Duplicate  int `json:"duplicate"`
		Spellcheck int `json:"spellcheck"`
	}
)

// PreTranslationsResponse defines the structure of a response when
// getting a pre-translation status.
type PreTranslationsResponse struct {
	Data *PreTranslation `json:"data"`
}

// PreTranslationsListResponse defines the structure of a response when
// getting a list of pre-translations.
type PreTranslationsListResponse struct {
	Data []*PreTranslationsResponse `json:"data"`
}

// PreTranslationReportResponse defines the structure of a response when
// getting a pre-translation report.
type PreTranslationReportResponse struct {
	Data *PreTranslationReport `json:"data"`
}

// PreTranslationsListOptions specifies the optional parameters to the
// TranslationsService.ListPreTranslationsWithOptions method.
type PreTranslationsListOptions struct {
	// Sort a list of pre-translations.
	// Example: orderBy=createdAt desc.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of PreTranslationsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *PreTranslationsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// PreTranslationScope defines which strings pre-translation is applied to.
type PreTranslationScope string

const (
	// PreTranslationScopeUntranslated applies pre-translation to strings without
	// an existing translation (default).
	PreTranslationScopeUntranslated PreTranslationScope = "untranslated"
	// PreTranslationScopeTranslated applies pre-translation to strings that
	// already have a translation (re-translation).
	PreTranslationScopeTranslated PreTranslationScope = "translated"
	// PreTranslationScopeAll applies pre-translation to both untranslated
	// and translated strings.
	PreTranslationScopeAll PreTranslationScope = "all"
)

// ReplaceTranslationsOption defines whether pre-translation replaces
// existing translations.
type ReplaceTranslationsOption string

const (
	// ReplaceTranslationsOptionNone adds the new translation alongside
	// existing ones (default).
	ReplaceTranslationsOptionNone ReplaceTranslationsOption = "none"
	// ReplaceTranslationsOptionAutoTranslated replaces auto-generated translations
	// (TM, MT, AI). Human translations are kept.
	ReplaceTranslationsOptionAutoTranslated ReplaceTranslationsOption = "autoTranslated"
	// ReplaceTranslationsOptionAll replaces all existing translations.
	ReplaceTranslationsOptionAll ReplaceTranslationsOption = "all"
)

// PreTranslationRequest defines the structure of a request to apply pre-translation.
type PreTranslationRequest struct {
	// Set of languages to which pre-translation should be applied.
	// Note: Required unless `taskId` is set.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// Files array that should be translated.
	// Note: Required unless `directoryIds` or `branchIds` is set (file-based projects).
	FileIDs []int `json:"fileIds,omitempty"`
	// Directories array that should be translated. Includes all nested files.
	// Note: Available only for file-based projects.
	DirectoryIDs []int `json:"directoryIds,omitempty"`
	// Branches array that should be translated.
	BranchIDs []int `json:"branchIds,omitempty"`
	// Task Identifier. Applies pre-translation to the strings of the specified task.
	// The target language is taken from the task.
	TaskID int `json:"taskId,omitempty"`
	// Style Guide Identifiers used during pre-translation.
	StyleGuideIDs []int `json:"styleGuideIds,omitempty"`
	// Defines pre-translation method. Enum: "tm", "mt", "ai". Default: "tm".
	//  - tm – pre-translation via Translation Memory.
	//  - mt – pre-translation via Machine Translation. "mt" should be used with `engineId` parameter.
	//  - ai – pre-translation via AI. "ai" should be used with `aiPromptId` parameter.
	Method string `json:"method,omitempty"`
	// Machine Translation engine Identifier. Required if `method` is set to "mt".
	EngineID int `json:"engineId,omitempty"`
	// AI Prompt Identifier. Required if `method` is set to "ai".
	AIPromptID int `json:"aiPromptId,omitempty"`
	// Defines which translations added by TM pre-translation should be auto-approved. Default: "none".
	// Enum: "all", "exceptAutoSubstituted", "perfectMatchApprovedOnly", "perfectMatchOnly", "none"
	//  - all – all
	//  - perfectMatchOnly – with perfect TM match
	//  - exceptAutoSubstituted – all (skip auto-substituted suggestions)
	//  - perfectMatchApprovedOnly - with perfect TM match (approved previously)
	//  - none – no auto-approve
	AutoApproveOption string `json:"autoApproveOption,omitempty"`
	// Adds translations even if the same translation already exists. Default is false.
	// Note: Works only with TM pre-translation method.
	DuplicateTranslations *bool `json:"duplicateTranslations,omitempty"`
	// Skip approved translations. Default is false.
	// Note: Works only with TM pre-translation method.
	SkipApprovedTranslations *bool `json:"skipApprovedTranslations,omitempty"`
	// Applies pre-translation for untranslated strings only. Default is true.
	// Note: Cannot be used together with `scope` in the same request.
	//
	// Deprecated: use Scope instead. PreTranslationScopeUntranslated is equivalent
	// to `true` and PreTranslationScopeAll is equivalent to `false`.
	TranslateUntranslatedOnly *bool `json:"translateUntranslatedOnly,omitempty"`
	// Which strings to apply pre-translation to.
	// Enum: "untranslated", "translated", "all". Default: "untranslated".
	// Note: Cannot be used together with the deprecated `translateUntranslatedOnly`.
	Scope PreTranslationScope `json:"scope,omitempty"`
	// Re-translates only if a string's current translation was modified prior to
	// this date (ISO 8601 date-time).
	// Note: Requires `scope` to be "translated" or "all".
	TranslationModifiedBefore string `json:"translationModifiedBefore,omitempty"`
	// Re-translates only if a string's current translation was modified after
	// this date (ISO 8601 date-time).
	// Note: Cannot be later than `translationModifiedBefore`.
	// Requires `scope` to be "translated" or "all".
	TranslationModifiedAfter string `json:"translationModifiedAfter,omitempty"`
	// Defines whether to replace existing translations with new pre-translations.
	// Enum: "none", "autoTranslated", "all". Default: "none".
	// Note: Values other than "none" require `scope` to be "translated" or "all".
	ReplaceTranslationsOption ReplaceTranslationsOption `json:"replaceTranslationsOption,omitempty"`
	// Removes approval on existing translations when applying pre-translations.
	// Default is false.
	// Note: Requires `scope` to be "translated" or "all". Cannot be used together with
	// `skipApprovedTranslations`, `autoApproveOption` (other than "none"), or when
	// `replaceTranslationsOption` is "all".
	ResetApprovalStatus *bool `json:"resetApprovalStatus,omitempty"`
	// If true, the user who started the pre-translation is notified when it
	// finishes, fails or is canceled. Default is false.
	NotifyOnCompletion *bool `json:"notifyOnCompletion,omitempty"`
	// Applies pre-translation only for the strings with perfect match
	// (source text and contextual information are identical). Default is false.
	// Note: Works only with TM pre-translation method.
	TranslateWithPerfectMatchOnly *bool `json:"translateWithPerfectMatchOnly,omitempty"`
	// Add translation when TM match is greater or equal to minimum match ratio.
	// Acceptable values: 40-100. Default: 100.
	// Note: Works only with TM pre-translation method. Available only for Crowdin Enterprise.
	MinimumMatchRatio int `json:"minimumMatchRatio,omitempty"`
	// Pre-translation priority. Enum: "low", "normal", "high". Default: "normal".
	Priority string `json:"priority,omitempty"`
	// Defines fallback languages mapping. The passed value should contain a map of
	// languageID as a key and an array of fallback language IDs as a value.
	//  - languageID – Crowdin ID for the specified language.
	//  - []string – an array containing fallback language IDs.
	// Note: Available only for TM Pre-Translation.
	FallbackLanguages map[string][]string `json:"fallbackLanguages,omitempty"`
	// Source language identifier.
	// Note: Available only for Crowdin.
	SourceLanguageID string `json:"sourceLanguageId,omitempty"`
	// Custom instruction for AI pre-translation. This instruction will be appended
	// to the AI prompt text. Max length: 10000 characters.
	// Note: Available only for AI pre-translation method.
	CustomInstruction string `json:"customInstruction,omitempty"`
	// Label Identifiers.
	LabelIDs []int `json:"labelIds,omitempty"`
	// Exclude Label Identifiers.
	ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PreTranslationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if len(r.LanguageIDs) == 0 && r.TaskID == 0 {
		return errors.New("languageIds is required")
	}
	if r.Method == "ai" && r.AIPromptID == 0 {
		return errors.New("aiPromptId is required")
	}
	if r.Method == "mt" && r.EngineID == 0 {
		return errors.New("engineId is required")
	}
	if r.MinimumMatchRatio != 0 && (r.MinimumMatchRatio < 40 || r.MinimumMatchRatio > 100) {
		return errors.New("minimumMatchRatio must be from 40 to 100")
	}
	if utf8.RuneCountInString(r.CustomInstruction) > 10000 {
		return errors.New("customInstruction must not exceed 10000 characters")
	}

	return r.validateScope()
}

// validateScope checks the scope-related rules of the pre-translation request.
func (r *PreTranslationRequest) validateScope() error {
	retranslate := false
	switch r.Scope {
	case PreTranslationScopeTranslated, PreTranslationScopeAll:
		retranslate = true
	case "":
		// The deprecated `translateUntranslatedOnly: false` is equivalent to the "all" scope.
		retranslate = r.TranslateUntranslatedOnly != nil && !*r.TranslateUntranslatedOnly
	case PreTranslationScopeUntranslated:
		// valid
	default:
		return fmt.Errorf("invalid scope: %q", r.Scope)
	}

	switch r.ReplaceTranslationsOption {
	case ReplaceTranslationsOptionNone, "":
		// valid
	case ReplaceTranslationsOptionAutoTranslated, ReplaceTranslationsOptionAll:
		if !retranslate {
			return errors.New("replaceTranslationsOption requires scope to be translated or all")
		}
	default:
		return fmt.Errorf("invalid replaceTranslationsOption: %q", r.ReplaceTranslationsOption)
	}

	if r.Scope != "" && r.TranslateUntranslatedOnly != nil {
		return errors.New("scope and translateUntranslatedOnly cannot be used together")
	}
	if err := r.validateTranslationModifiedDates(retranslate); err != nil {
		return err
	}
	if r.DuplicateTranslations != nil && *r.DuplicateTranslations &&
		r.ReplaceTranslationsOption == ReplaceTranslationsOptionAll {
		return errors.New("duplicateTranslations cannot be used when replaceTranslationsOption is all")
	}

	return r.validateResetApprovalStatus(retranslate)
}

// validateTranslationModifiedDates checks the rules of the `translationModifiedBefore`
// and `translationModifiedAfter` options.
func (r *PreTranslationRequest) validateTranslationModifiedDates(retranslate bool) error {
	if r.TranslationModifiedBefore == "" && r.TranslationModifiedAfter == "" {
		return nil
	}
	if !retranslate {
		return errors.New("translationModifiedBefore and translationModifiedAfter require scope to be translated or all")
	}
	if r.TranslationModifiedBefore == "" || r.TranslationModifiedAfter == "" {
		return nil
	}

	before, okBefore := parsePreTranslationDateTime(r.TranslationModifiedBefore)
	after, okAfter := parsePreTranslationDateTime(r.TranslationModifiedAfter)
	if okBefore && okAfter && after.After(before) {
		return errors.New("translationModifiedAfter cannot be later than translationModifiedBefore")
	}

	return nil
}

// parsePreTranslationDateTime parses an RFC 3339 date-time string and reports whether it succeeded.
func parsePreTranslationDateTime(value string) (time.Time, bool) {
	t, err := time.Parse(time.RFC3339, value)

	return t, err == nil
}

// validateResetApprovalStatus checks the rules of the `resetApprovalStatus` option.
func (r *PreTranslationRequest) validateResetApprovalStatus(retranslate bool) error {
	if r.ResetApprovalStatus == nil || !*r.ResetApprovalStatus {
		return nil
	}

	if !retranslate {
		return errors.New("resetApprovalStatus requires scope to be translated or all")
	}
	if r.SkipApprovedTranslations != nil && *r.SkipApprovedTranslations {
		return errors.New("resetApprovalStatus cannot be used together with skipApprovedTranslations")
	}
	if r.AutoApproveOption != "" && r.AutoApproveOption != "none" {
		return errors.New("resetApprovalStatus cannot be used together with autoApproveOption")
	}
	if r.ReplaceTranslationsOption == ReplaceTranslationsOptionAll {
		return errors.New("resetApprovalStatus cannot be used when replaceTranslationsOption is all")
	}

	return nil
}

// BuildProjectDirectoryTranslationRequest defines the structure of a request
// to build project directory translation.
type BuildProjectDirectoryTranslationRequest struct {
	// Specify target languages for build.
	// Leave this field empty to build all target languages.
	TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
	// Defines whether to export only translated strings. Default: false.
	// Note: true value can't be used with `skipUntranslatedFiles=true` in same request .
	SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated file. Default: false.
	// Note: true value can't be used with `skipUntranslatedStrings=true` in same request.
	SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default: false.
	ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`
	// Preserve folder hierarchy. Default: false.
	PreserveFolderHierarchy *bool `json:"preserveFolderHierarchy,omitempty"`

	// Defines whether to export only approved strings.
	// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
	// in same request.
	ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
	// Defines whether to export only strings that passed workflow.
	// Note: true value can't be used with `exportWithMinApprovalsCount>0` in same request
	// or in projects without an assigned workflow.
	ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
}

// Validate checks if the build project directory translation request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectDirectoryTranslationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	if (r.SkipUntranslatedStrings != nil && r.SkipUntranslatedFiles != nil) &&
		(*r.SkipUntranslatedStrings && *r.SkipUntranslatedFiles) {
		return errors.New("skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request")
	}
	if (r.ExportWithMinApprovalsCount != nil && r.ExportStringsThatPassedWorkflow != nil) &&
		(*r.ExportWithMinApprovalsCount > 0 && *r.ExportStringsThatPassedWorkflow) {
		return fmt.Errorf("exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request")
	}

	return nil
}

// BuildProjectDirectoryTranslation represents a project directory build.
type BuildProjectDirectoryTranslation struct {
	ID         int    `json:"id"`
	ProjectID  int    `json:"projectId"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	FinishedAt string `json:"finishedAt,omitempty"`
}

// BuildProjectFileTranslationRequest defines the structure of a request
// to build project file translation.
type BuildProjectFileTranslationRequest struct {
	// Target Language Identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// Defines whether to export only translated strings. Default: false.
	// Note: true value can't be used with `skipUntranslatedFiles=true`` in same request.
	SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated file. Default: false.
	// Note: true value can't be used with `skipUntranslatedStrings=true` in same request.
	SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default: false.
	ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`

	// Defines whether to export only approved strings.
	// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
	// in same request.
	ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
	// Defines whether to export only strings that passed workflow.
	// Note: true value can't be used with `exportWithMinApprovalsCount>0` in same request
	// or in projects without an assigned workflow.
	ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
}

// Validate checks if the build project file translation request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectFileTranslationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if len(r.TargetLanguageID) == 0 {
		return errors.New("targetLanguageId is required")
	}

	if (r.SkipUntranslatedStrings != nil && r.SkipUntranslatedFiles != nil) &&
		(*r.SkipUntranslatedStrings && *r.SkipUntranslatedFiles) {
		return errors.New("skipUntranslatedStrings and skipUntranslatedFiles must not be true at the same request")
	}
	if (r.ExportWithMinApprovalsCount != nil && r.ExportStringsThatPassedWorkflow != nil) &&
		(*r.ExportWithMinApprovalsCount > 0 && *r.ExportStringsThatPassedWorkflow) {
		return fmt.Errorf("exportWithMinApprovalsCount and exportStringsThatPassedWorkflow must not be true at the same request")
	}

	return nil
}

// TranslationsBuildsListOptions specifies the optional parameters to the
// TranslationsService.ListProjectBuilds method.
type TranslationsBuildsListOptions struct {
	ListOptions

	// Branch Identifier. Filter builds by branchId.
	BranchID int `json:"branchId,omitempty"`
}

// Values returns the url.Values representation of the query options.
func (o *TranslationsBuildsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}

	return v, len(v) > 0
}

// TranslationsProjectBuild represents a project build.
type TranslationsProjectBuild struct {
	ID         int              `json:"id"`
	ProjectID  int              `json:"projectId"`
	Status     string           `json:"status"`
	Progress   int              `json:"progress"`
	CreatedAt  string           `json:"createdAt"`
	UpdatedAt  string           `json:"updatedAt"`
	FinishedAt string           `json:"finishedAt,omitempty"`
	Attributes *BuildAttributes `json:"attributes,omitempty"`
	// Error is present only when the build status is "failed".
	Error *BuildError `json:"error,omitempty"`
}

// BuildError represents the error details of a failed build.
type BuildError struct {
	Message string `json:"message"`
}

// BuildAttributes represents the attributes of a project build.
type BuildAttributes struct {
	BranchID                        *int     `json:"branchId,omitempty"`
	DirectoryID                     *int     `json:"directoryId,omitempty"`
	TargetLanguageIDs               []string `json:"targetLanguageIds,omitempty"`
	SkipUntranslatedStrings         *bool    `json:"skipUntranslatedStrings,omitempty"`
	SkipUntranslatedFiles           *bool    `json:"skipUntranslatedFiles,omitempty"`
	ExportApprovedOnly              *bool    `json:"exportApprovedOnly,omitempty"`
	ExportWithMinApprovalsCount     *int     `json:"exportWithMinApprovalsCount,omitempty"`
	ExportStringsThatPassedWorkflow *bool    `json:"exportStringsThatPassedWorkflow,omitempty"`

	Pseudo               *bool   `json:"pseudo,omitempty"`
	Prefix               *string `json:"prefix,omitempty"`
	Suffix               *string `json:"suffix,omitempty"`
	LengthTransformation *int    `json:"lengthTransformation,omitempty"`
	CharTransformation   *string `json:"charTransformation,omitempty"`
}

// TranslationsProjectBuildResponse defines the structure of a response when
// getting a project build.
type TranslationsProjectBuildResponse struct {
	Data *TranslationsProjectBuild `json:"data"`
}

// TranslationsProjectBuildsListResponse defines the structure of a response when
// getting a list of project builds.
type TranslationsProjectBuildsListResponse struct {
	Data []*TranslationsProjectBuildResponse `json:"data"`
}

type (
	// BuildProjectTranslationRequester interface that allows accepting
	// BuildProjectRequest and PseudoBuildProjectRequest types.
	BuildProjectTranslationRequester interface {
		ValidateBuildRequest() error
	}

	// BuildProjectRequest defines the structure of a request to build a project.
	BuildProjectRequest struct {
		// Branch Identifier.
		BranchID int `json:"branchId,omitempty"`
		// Specify target languages for build.
		// Leave this field empty to build all target languages
		TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
		// Defines whether to export only translated strings.
		// Note: true value can't be used with `skipUntranslatedFiles=true` in same request.
		SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
		// Defines whether to export only translated files.
		// Note: true value can't be used with `skipUntranslatedStrings=true` in same request.
		SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
		// Defines whether to export only approved strings.
		ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`

		// Defines whether to export only approved strings.
		// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
		// in same request.
		ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
		// Defines whether to export only strings that passed workflow.
		// Note: true value can't be used with `exportWithMinApprovalsCount>0` in same request
		// or in projects without an assigned workflow.
		ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
	}

	// PsuedoBuildProjectRequest defines the structure of a request to build a project
	// with pseudo translations.
	PseudoBuildProjectRequest struct {
		// Flag for detecting pseudo translation. Default: false.
		Pseudo *bool `json:"pseudo"`
		// Branch Identifier.
		BranchID int `json:"branchId,omitempty"`
		// Add special characters at the beginning of each string to show
		// where messages have been concatenated together.
		Prefix string `json:"prefix,omitempty"`
		// Add special characters at the end of each string to show where
		// messages have been concatenated together.
		Suffix string `json:"suffix,omitempty"`
		// Make string larger or shorter.
		// Acceptable values must be from -50 to 100. Default is 0.
		LengthTransformation *int `json:"lengthTransformation,omitempty"`
		// Transforms characters to other languages.
		// Enum: "asian", "cyrillic", "european", "arabic".
		CharTransformation string `json:"charTransformation,omitempty"`
	}
)

// Validate checks if the build project request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *BuildProjectRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateBuildRequest()
}

// ValidateBuildRequest implements the BuildProjectTranslationRequest interface.
func (r *BuildProjectRequest) ValidateBuildRequest() error {
	if (r.SkipUntranslatedStrings != nil && r.SkipUntranslatedFiles != nil) &&
		(*r.SkipUntranslatedStrings && *r.SkipUntranslatedFiles) {
		return errors.New("`skipUntranslatedStrings` and `skipUntranslatedFiles` must not be true at the same request")
	}
	if (r.ExportWithMinApprovalsCount != nil && r.ExportStringsThatPassedWorkflow != nil) &&
		(*r.ExportWithMinApprovalsCount > 0 && *r.ExportStringsThatPassedWorkflow) {
		return fmt.Errorf("`exportWithMinApprovalsCount` and `exportStringsThatPassedWorkflow` must not be true at the same request")
	}

	return nil
}

// Validate checks if the build project request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PseudoBuildProjectRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateBuildRequest()
}

// PseudoBuildProjectRequest implements the BuildProjectTranslationRequest interface.
func (r *PseudoBuildProjectRequest) ValidateBuildRequest() error {
	if r.LengthTransformation != nil &&
		(*r.LengthTransformation < -50 || *r.LengthTransformation > 100) {
		return errors.New("lengthTransformation must be from -50 to 100")
	}

	return nil
}

// UploadTranslationsRequest defines the structure of a request to upload translations.
type UploadTranslationsRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// File Identifier for import.
	// Note: Required for content in all formats except XLIFF.
	FileID int `json:"fileId,omitempty"`
	// Branch Identifier for import.
	// Note: Required for string based API.
	BranchID int `json:"branchId,omitempty"`
	// Defines whether to add translation if it's the same as the source string.
	// Default: false.
	ImportEqSuggestions *bool `json:"importEqSuggestions,omitempty"`
	// Mark uploaded translations as approved. Default: false.
	AutoApproveImported *bool `json:"autoApproveImported,omitempty"`
	// Allow translations upload to hidden source strings. Default: false.
	TranslateHidden *bool `json:"translateHidden,omitempty"`
	// Defines whether to add translation to TM. Default: true.
	AddToTM *bool `json:"addToTm,omitempty"`
}

// Validate checks if the upload translations request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *UploadTranslationsRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if r.FileID > 0 && r.BranchID > 0 {
		return errors.New("fileId and branchId can not be used at the same request")
	}
	return nil
}

type (
	// UploadTranslations represents the uploaded translations.
	UploadTranslations struct {
		ProjectID  int    `json:"projectId"`
		StorageID  int    `json:"storageId"`
		LanguageID string `json:"languageId"`
		FileID     int    `json:"fileId"`
	}

	// UploadTranslationsResponse defines the structure of a response when
	// uploading translations.
	UploadTranslationsResponse struct {
		Data *UploadTranslations `json:"data"`
	}
)

// TranslationImportRequest defines the structure of a request to import translations.
type TranslationImportRequest struct {
	// Storage Identifier.
	StorageID int `json:"storageId"`
	// Language Identifiers.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// File Identifier for import.
	// Note: Required for content in all formats except XLIFF (file-based projects).
	FileID int `json:"fileId,omitempty"`
	// Branch Identifier for import.
	// Note: Available only for string-based projects.
	BranchID int `json:"branchId,omitempty"`
	// Defines whether to add translation if it's the same as the source string.
	// Default: false.
	ImportEqSuggestions *bool `json:"importEqSuggestions,omitempty"`
	// Mark uploaded translations as approved. Default: false.
	AutoApproveImported *bool `json:"autoApproveImported,omitempty"`
	// Allow translations upload to hidden source strings. Default: false.
	TranslateHidden *bool `json:"translateHidden,omitempty"`
	// Defines whether to add translation to TM. Default: true.
	AddToTM *bool `json:"addToTm,omitempty"`
	// Spreadsheet file import options.
	// Note: Available only for string-based projects.
	ImportOptions *TranslationImportOptions `json:"importOptions,omitempty"`
}

// TranslationImportOptions defines the spreadsheet file import options
// for string-based projects.
type TranslationImportOptions struct {
	// Defines data columns mapping. The column numbering starts at 0.
	// Keys: "none", "identifier", "sourceOrTranslation", "translation", or a
	// language identifier (e.g. "en", "de").
	Scheme map[string]int `json:"scheme,omitempty"`
}

// Validate checks if the import translations request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *TranslationImportRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if r.FileID > 0 && r.BranchID > 0 {
		return errors.New("fileId and branchId can not be used at the same request")
	}

	return nil
}

type (
	// TranslationImport represents a translation import job.
	TranslationImport struct {
		Identifier string                       `json:"identifier"`
		Status     string                       `json:"status"`
		Progress   int                          `json:"progress"`
		Attributes *TranslationImportAttributes `json:"attributes"`
		CreatedAt  string                       `json:"createdAt"`
		UpdatedAt  string                       `json:"updatedAt"`
		StartedAt  *string                      `json:"startedAt"`
		FinishedAt *string                      `json:"finishedAt"`
	}

	// TranslationImportAttributes represents the attributes of a translation import job.
	TranslationImportAttributes struct {
		StorageID           int      `json:"storageId"`
		FileID              int      `json:"fileId,omitempty"`
		BranchID            int      `json:"branchId,omitempty"`
		LanguageIDs         []string `json:"languageIds"`
		ImportEqSuggestions bool     `json:"importEqSuggestions"`
		AutoApproveImported bool     `json:"autoApproveImported"`
		TranslateHidden     bool     `json:"translateHidden"`
		AddToTM             bool     `json:"addToTm"`
	}

	// TranslationImportResponse defines the structure of a response when
	// importing translations or getting an import status.
	TranslationImportResponse struct {
		Data *TranslationImport `json:"data"`
	}
)

type (
	// TranslationImportReport represents a translation import report.
	TranslationImportReport struct {
		Languages []*TranslationImportReportLanguage `json:"languages"`
	}

	// TranslationImportReportLanguage represents a language entry of
	// the translation import report.
	TranslationImportReportLanguage struct {
		ID string `json:"id"`
		// Files statistics (file-based projects).
		Files []*TranslationImportReportItem `json:"files,omitempty"`
		// Branches statistics (string-based projects).
		Branches                 []*TranslationImportReportItem  `json:"branches,omitempty"`
		Skipped                  *TranslationImportReportSkipped `json:"skipped,omitempty"`
		SkippedQACheckCategories map[string]int                  `json:"skippedQaCheckCategories,omitempty"`
	}

	// TranslationImportReportItem represents a file or branch entry
	// of the translation import report.
	TranslationImportReportItem struct {
		ID         string                    `json:"id"`
		Statistics *LanguageReportStatistics `json:"statistics"`
	}

	// TranslationImportReportSkipped represents skipped translations
	// statistics of the translation import report.
	TranslationImportReportSkipped struct {
		TranslationEqSource int `json:"translationEqSource"`
		HiddenStrings       int `json:"hiddenStrings"`
		QACheck             int `json:"qaCheck"`
	}

	// TranslationImportReportResponse defines the structure of a response when
	// getting a translation import report.
	TranslationImportReportResponse struct {
		Data *TranslationImportReport `json:"data"`
	}
)

// ExportTranslationRequest defines the structure of a request
// to export translations.
type ExportTranslationRequest struct {
	// Specify target language for export.
	TargetLanguageID string `json:"targetLanguageId"`
	// Defines export file format. Use API Type feature specified at the
	// corresponding file format from Crowdin Store.
	// Note: the `format` parameter is required in all cases except when you'd like
	// to export translations for a single file in its original format.
	Format string `json:"format,omitempty"`
	// Label Identifiers.
	LabelIDs []int `json:"labelIds,omitempty"`
	// Branch Identifiers.
	// Note: Can't be used with `directoryIds` or `fileIds` in same request.
	BranchIDs []int `json:"branchIds,omitempty"`
	// Directory Identifiers.
	// Note: Can't be used with `branchIds` or `fileIds` in same request.
	DirectoryIDs []int `json:"directoryIds,omitempty"`
	// File Identifiers.
	// Note: Can't be used with `branchIds` or `directoryIds` in same request.
	FileIDs []int `json:"fileIds,omitempty"`
	// Defines whether to export only translated strings. Default is false.
	// Note: Can't be used with `skipUntranslatedFiles` in same request.
	SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only translated file. Default is false.
	// Note: Can't be used with `skipUntranslatedStrings` in same request.
	SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
	// Defines whether to export only approved strings. Default is false.
	ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`

	// Defines whether to export only approved strings.
	// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
	// in same request.
	ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
	// Defines whether to export only strings that passed workflow.
	// Note: true value can't be used with `exportWithMinApprovalsCount>0` in same request
	// or in projects without an assigned workflow.
	ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ExportTranslationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.TargetLanguageID == "" {
		return errors.New("targetLanguageId is required")
	}

	return nil
}

type (
	// DownloadLink represents a download link.
	DownloadLink struct {
		URL      string `json:"url"`
		ExpireIn string `json:"expireIn"`

		Etag *string `json:"etag,omitempty"`
	}

	// DownloadLinkResponse defines the structure of a response when
	// getting a download URL with its expiration time.
	DownloadLinkResponse struct {
		Data *DownloadLink `json:"data"`
	}
)
