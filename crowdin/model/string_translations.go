package model

import (
	"errors"
	"fmt"
	"net/url"
)

// Approval represents a Crowdin translation approval.
type Approval struct {
	ID            int        `json:"id"`
	User          *ShortUser `json:"user"`
	TranslationID int        `json:"translationId"`
	StringID      int        `json:"stringId"`
	LanguageID    string     `json:"languageId"`
	CreatedAt     string     `json:"createdAt"`
	// File Identifier. Present for asset approvals.
	FileID int `json:"fileId,omitempty"`
	// Correction Identifier. Present for correction approvals (Crowdin Enterprise only).
	CorrectionID int `json:"correctionId,omitempty"`
	// Workflow Step Identifier (Crowdin Enterprise only).
	WorkflowStepID int `json:"workflowStepId,omitempty"`
}

// ApprovalsGetResponse defines the structure of the response when
// getting a single translation approval.
type ApprovalsGetResponse struct {
	Data *Approval `json:"data"`
}

// ApprovalsListResponse defines the structure of the response when
// getting a list of translation approvals.
type ApprovalsListResponse struct {
	Data []*ApprovalsGetResponse `json:"data"`
}

// ApprovalsListOptions specifies the optional parameters to the
// StringTranslationsService.ListApprovals method.
type ApprovalsListOptions struct {
	// Sort a list of approvals.
	// Enum: id, createdAt. Default: id.
	// Example: orderBy=createdAt desc,id.
	OrderBy string `json:"orderBy,omitempty"`
	// File Identifier.
	// Note: Must be used together with `languageId`.
	FileID int `json:"fileId,omitempty"`
	// Label Identifiers.
	// Example: labelIds=1,2,3,4,5
	LabelIDs []int `json:"labelIds,omitempty"`
	// Exclude Label Identifiers.
	ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
	// String Identifier.
	// Note: Must be used together with `languageId`.
	StringID int `json:"stringId,omitempty"`
	// Language Identifier.
	// Note: Must be used together with `stringId` or `fileId`.
	LanguageID string `json:"languageId,omitempty"`
	// Translation Identifier.
	// Note: If specified, `fileId`, `stringId` and `languageId` are ignored.
	TranslationID int `json:"translationId,omitempty"`
	// Correction Identifier.
	// Note: Can't be used with `translationId`, `languageId`, `stringId` or `fileId`.
	// Available only for Crowdin Enterprise projects with advanced workflow.
	CorrectionID int `json:"correctionId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the ApprovalsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ApprovalsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.FileID > 0 {
		v.Add("fileId", fmt.Sprintf("%d", o.FileID))
	}
	if len(o.LabelIDs) > 0 {
		v.Add("labelIds", JoinSlice(o.LabelIDs))
	}
	if len(o.ExcludeLabelIDs) > 0 {
		v.Add("excludeLabelIds", JoinSlice(o.ExcludeLabelIDs))
	}
	if o.StringID > 0 {
		v.Add("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.LanguageID != "" {
		v.Add("languageId", o.LanguageID)
	}
	if o.TranslationID > 0 {
		v.Add("translationId", fmt.Sprintf("%d", o.TranslationID))
	}
	if o.CorrectionID > 0 {
		v.Add("correctionId", fmt.Sprintf("%d", o.CorrectionID))
	}

	return v, len(v) > 0
}

// TranslationAlignment represents a translation alignment.
type TranslationAlignment struct {
	Words []*WordAlignment `json:"words"`
}

// WordAlignment represents a word alignments.
type WordAlignment struct {
	Text       string       `json:"text"`
	Alignments []*Alignment `json:"alignments"`
}

// Alignment represents a word alignment.
type Alignment struct {
	SourceWord  string `json:"sourceWord"`
	SourceLemma string `json:"sourceLemma"`
	TargetWord  string `json:"targetWord"`
	TargetLemma string `json:"targetLemma"`
	Match       int    `json:"match"`
	Probability int    `json:"probability"`
}

// TranslationAlignmentResponse defines the structure of the response when
// aligning translations.
type TranslationAlignmentResponse struct {
	Data *TranslationAlignment `json:"data"`
}

// TranslationAlignmentRequest defines the structure of the request
// to align translations.
type TranslationAlignmentRequest struct {
	// Source Language Identifier.
	SourceLanguageID string `json:"sourceLanguageId"`
	// Target Language Identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// Text for alignment.
	Text string `json:"text"`
}

// Validate checks if the TranslationAlignmentRequest is valid.
// It implements the crowdin.RequestValidator interface.
func (r *TranslationAlignmentRequest) Validate() error {
	if r == nil {
		return errors.New("request cannot be nil")
	}
	if r.SourceLanguageID == "" {
		return errors.New("source language ID is required")
	}
	if r.TargetLanguageID == "" {
		return errors.New("target language ID is required")
	}
	if r.Text == "" {
		return errors.New("text is required")
	}
	return nil
}

// LanguageTranslation represents a language translation.
// Contains the plain, plural, or ICU translation.
type LanguageTranslation struct {
	StringID        int        `json:"stringId"`
	ContentType     string     `json:"contentType"`
	TranslationID   *int       `json:"translationId,omitempty"`
	Text            *string    `json:"text,omitempty"`
	User            *ShortUser `json:"user,omitempty"`
	CreatedAt       *string    `json:"createdAt,omitempty"`
	Provider        *string    `json:"provider,omitempty"`
	ProviderID      *int       `json:"providerId,omitempty"`
	IsPreTranslated *bool      `json:"isPreTranslated,omitempty"`
	MatchRate       *int       `json:"matchRate,omitempty"`
	MatchType       *string    `json:"matchType,omitempty"`
	// QA Issues Status. Enum: inProgress, failed, passed.
	QAIssuesStatus *string `json:"qaIssuesStatus,omitempty"`
	// Asset URL. Present for asset translations.
	URL *string `json:"url,omitempty"`

	Plurals []*LanguageTranslationPlural `json:"plurals,omitempty"`
}

// LanguageTranslationPlural represents a plural language translation
// and is part of the LanguageTranslation.
type LanguageTranslationPlural struct {
	TranslationID   int        `json:"translationId"`
	Text            string     `json:"text"`
	PluralForm      string     `json:"pluralForm"`
	User            *ShortUser `json:"user"`
	CreatedAt       string     `json:"createdAt"`
	Provider        *string    `json:"provider,omitempty"`
	ProviderID      *int       `json:"providerId,omitempty"`
	IsPreTranslated *bool      `json:"isPreTranslated,omitempty"`
	MatchRate       *int       `json:"matchRate,omitempty"`
	MatchType       *string    `json:"matchType,omitempty"`
}

// LanguageTranslationsGetResponse defines the structure of the response when
// retrieving a list of language translations.
type LanguageTranslationsListResponse struct {
	Data []struct {
		Data *LanguageTranslation `json:"data"`
	} `json:"data"`
}

// LanguageTranslationsListOptions specifies the optional parameters to the
// StringTranslationsService.ListLanguageTranslations method.
type LanguageTranslationsListOptions struct {
	// Sort a list of translations.
	// Enum: text, stringId, translationId, createdAt. Default: stringId.
	// Example: orderBy=createdAt desc,text
	OrderBy string `json:"orderBy,omitempty"`
	// String Identifiers. Filter translations by `stringIds`.
	// Example: stringIds=1,2,3,4,5
	StringIDs []int `json:"stringIds,omitempty"`
	// Label Identifiers. Filter translations by `labelIds`.
	// Example: labelIds=1,2,3,4,5
	LabelIDs []int `json:"labelIds,omitempty"`
	// File Identifier. Filter translations by `fileId`.
	// Note: Can't be used with `branchId` or `directoryId` in the same request.
	FileID int `json:"fileId,omitempty"`
	// Branch Identifier. Filter translations by `branchId`.
	// Note: Can't be used with `fileId` or `directoryId` in the same request.
	BranchID int `json:"branchId,omitempty"`
	// Directory Identifier. Filter translations by `directoryId`.
	// Note: Can't be used with `fileId` or `branchId` in the same request.
	DirectoryID int `json:"directoryId,omitempty"`
	// Filter translations by CroQL.
	// Note: Can't be used with `stringIds`, `labelIds` or `fileId`
	// in the same request.
	CroQL string `json:"croql,omitempty"`
	// Enable denormalize placeholders.
	// Enum: 0, 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`
	// Only approved translations. Enum: 0, 1.
	// Note: Can't be used with `croql` in the same request. Available only for Crowdin.
	ApprovedOnly *int `json:"approvedOnly,omitempty"`
	// Only translations that passed workflow. Enum: 0, 1.
	// Note: Can't be used with `minApprovalCount` in the same request.
	// Available only for Crowdin Enterprise.
	PassedWorkflow *int `json:"passedWorkflow,omitempty"`
	// Minimum approval count.
	// Note: Can't be used with `passedWorkflow` in the same request.
	// Available only for Crowdin Enterprise.
	MinApprovalCount int `json:"minApprovalCount,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the LanguageTranslationsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *LanguageTranslationsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if len(o.StringIDs) > 0 {
		v.Add("stringIds", JoinSlice(o.StringIDs))
	}
	if len(o.LabelIDs) > 0 {
		v.Add("labelIds", JoinSlice(o.LabelIDs))
	}
	if o.FileID > 0 {
		v.Add("fileId", fmt.Sprintf("%d", o.FileID))
	}
	if o.BranchID > 0 {
		v.Add("branchId", fmt.Sprintf("%d", o.BranchID))
	}
	if o.DirectoryID > 0 {
		v.Add("directoryId", fmt.Sprintf("%d", o.DirectoryID))
	}
	if o.CroQL != "" {
		v.Add("croql", o.CroQL)
	}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		v.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	if o.ApprovedOnly != nil && (*o.ApprovedOnly == 0 || *o.ApprovedOnly == 1) {
		v.Add("approvedOnly", fmt.Sprintf("%d", *o.ApprovedOnly))
	}
	if o.PassedWorkflow != nil && (*o.PassedWorkflow == 0 || *o.PassedWorkflow == 1) {
		v.Add("passedWorkflow", fmt.Sprintf("%d", *o.PassedWorkflow))
	}
	if o.MinApprovalCount > 0 {
		v.Add("minApprovalCount", fmt.Sprintf("%d", o.MinApprovalCount))
	}

	return v, len(v) > 0
}

// Translation represents a Crowdin translation.
type Translation struct {
	ID                 int        `json:"id"`
	Text               string     `json:"text"`
	PluralCategoryName string     `json:"pluralCategoryName"`
	User               *ShortUser `json:"user"`
	Rating             int        `json:"rating"`
	Provider           *string    `json:"provider,omitempty"`
	IsPreTranslated    bool       `json:"isPreTranslated"`
	CreatedAt          string     `json:"createdAt"`
	// Provider specific identifier: a TM for `tm`, an AI Prompt for `ai`,
	// an MT engine for machine translation providers.
	ProviderID *int `json:"providerId,omitempty"`
	// Translation Memory match rate in percent (40-100).
	MatchRate *int `json:"matchRate,omitempty"`
	// Translation Memory match type. Enum: perfect, exact, fuzzy.
	MatchType *string `json:"matchType,omitempty"`
	// Asset URL. Present for asset translations.
	URL string `json:"url,omitempty"`
	// Workflow Step Identifier (Crowdin Enterprise only).
	WorkflowStepID int `json:"workflowStepId,omitempty"`
}

// SearchTranslation represents a translation found by the search.
type SearchTranslation struct {
	Translation

	// Project Identifier.
	ProjectID int `json:"projectId"`
	// Source String Identifier.
	StringID int `json:"stringId"`
	// Target Language Identifier.
	LanguageID string `json:"languageId"`
}

// SearchTranslationResponse defines the structure of the response when
// getting a single found translation.
type SearchTranslationResponse struct {
	Data *SearchTranslation `json:"data"`
}

// SearchTranslationsListResponse defines the structure of the response when
// searching translations.
type SearchTranslationsListResponse struct {
	Data []*SearchTranslationResponse `json:"data"`
}

// SearchTranslationsListOptions specifies the parameters to the
// StringTranslationsService.SearchTranslations method.
type SearchTranslationsListOptions struct {
	// Search translations by text. Required.
	Filter string `json:"filter"`
	// Project identifiers to search across (max 50).
	// Omit to search all accessible projects.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// Owner (user) whose projects to search when `projectIds` is omitted.
	// Note: Available only for Crowdin.
	UserID int `json:"userId,omitempty"`
	// Filter by target language identifiers.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// Enable denormalize placeholders.
	// Enum: 0, 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the SearchTranslationsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *SearchTranslationsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if len(o.ProjectIDs) > 0 {
		v.Add("projectIds", JoinSlice(o.ProjectIDs))
	}
	if o.UserID > 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}
	if len(o.LanguageIDs) > 0 {
		v.Add("languageIds", JoinSlice(o.LanguageIDs))
	}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		v.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}

	return v, len(v) > 0
}

// Validate checks if the SearchTranslationsListOptions are valid.
func (o *SearchTranslationsListOptions) Validate() error {
	if o == nil || o.Filter == "" {
		return errors.New("filter is required")
	}
	if len(o.ProjectIDs) > 50 {
		return errors.New("projectIds must not contain more than 50 items")
	}

	return nil
}

// TranslationGetResponse defines the structure of the response when
// getting a single translation.
type TranslationGetResponse struct {
	Data *Translation `json:"data"`
}

// TranslationsListResponse defines the structure of the response when
// getting a list of translations.
type TranslationsListResponse struct {
	Data []*TranslationGetResponse `json:"data"`
}

// TranslationGetOptions specifies the optional parameters to the
// StringTranslationsService.GetTranslation method.
type TranslationGetOptions struct {
	// Enable denormalize placeholders.
	// Enum: 0, 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`
}

// Values returns the url.Values representation of the TranslationGetOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *TranslationGetOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v := url.Values{}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		v.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}
	return v, len(v) > 0
}

// StringTranslationsListOptions specifies the optional parameters to the
// StringTranslationsService.ListTranslations method.
type StringTranslationsListOptions struct {
	// Sort a list of translations.
	// Enum: id, text, rating, createdAt. Default: id.
	// Example: orderBy=createdAt desc,name,priority
	OrderBy string `json:"orderBy,omitempty"`
	// String Identifier.
	// Note: Must be used together with `languageId`.
	StringID int `json:"stringId,omitempty"`
	// Language Identifier.
	// Note: Must be used together with `stringId`.
	LanguageID string `json:"languageId,omitempty"`
	// Denormalize Placeholders.
	// Enum: 0, 1. Default: 0.
	DenormalizePlaceholders *int `json:"denormalizePlaceholders,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the StringTranslationsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *StringTranslationsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.StringID > 0 {
		v.Add("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.LanguageID != "" {
		v.Add("languageId", o.LanguageID)
	}
	if o.DenormalizePlaceholders != nil &&
		(*o.DenormalizePlaceholders == 0 || *o.DenormalizePlaceholders == 1) {
		v.Add("denormalizePlaceholders", fmt.Sprintf("%d", *o.DenormalizePlaceholders))
	}

	return v, len(v) > 0
}

// TranslationAddRequest defines the structure of the request
// to add a translation.
type TranslationAddRequest struct {
	// String Identifier.
	// Note: Must be used together with `languageId`.
	StringID int `json:"stringId,omitempty"`
	// Language Identifier.
	// Note: Must be used together with `stringId`.
	LanguageID string `json:"languageId"`
	// Translation text.
	Text string `json:"text,omitempty"`
	// Plural form. Enum: zero, one, two, few, many, and other.
	// Note: Will be saved only if the source string has plurals and `pluralCategoryName`
	// is equal to the one available for the language you add translations to.
	PluralCategoryName string `json:"pluralCategoryName,omitempty"`
	// Defines whether to add translation to TM. Default: true.
	AddToTM *bool `json:"addToTm,omitempty"`
	// Translation provider type. Required when `providerId` or `isPreTranslated` is specified.
	// Enum: tm, global_tm, google, microsoft, crowdin, deepl, amazon, google_automl,
	// modernmt, custom_mt, ai.
	// Note: `global_tm` value can't be used with `providerId` in the same request.
	Provider string `json:"provider,omitempty"`
	// Provider specific identifier (TM, AI Prompt or MT engine identifier).
	// Note: Can't be used if `provider` value is `global_tm`.
	ProviderID int `json:"providerId,omitempty"`
	// Defines whether this is an auto-translated translation. Default: false.
	IsPreTranslated *bool `json:"isPreTranslated,omitempty"`

	// File Identifier. Used to add an asset translation together with `storageId`.
	FileID int `json:"fileId,omitempty"`
	// Storage Identifier. Used to add an asset translation together with `fileId`.
	StorageID int `json:"storageId,omitempty"`
}

// Validate checks if the TranslationAddRequest is valid.
// It implements the crowdin.RequestValidator interface.
func (r *TranslationAddRequest) Validate() error {
	if r == nil {
		return errors.New("request cannot be nil")
	}

	// Asset translation.
	if r.StorageID > 0 || r.FileID > 0 {
		if r.FileID == 0 {
			return errors.New("file ID is required")
		}
		if r.LanguageID == "" {
			return errors.New("language ID is required")
		}
		if r.StorageID == 0 {
			return errors.New("storage ID is required")
		}
		return nil
	}

	if r.StringID == 0 {
		return errors.New("string ID is required")
	}
	if r.LanguageID == "" {
		return errors.New("language ID is required")
	}
	if r.Text == "" {
		return errors.New("text is required")
	}
	if r.Provider == "" && (r.ProviderID > 0 || r.IsPreTranslated != nil) {
		return errors.New("provider is required when providerId or isPreTranslated is specified")
	}
	if r.Provider == "global_tm" && r.ProviderID > 0 {
		return errors.New("providerId can't be used with the global_tm provider")
	}
	return nil
}

// Vote represents a Crowdin translation vote.
type Vote struct {
	ID            int        `json:"id"`
	User          *ShortUser `json:"user"`
	TranslationID int        `json:"translationId"`
	VotedAt       string     `json:"votedAt"`
	Mark          string     `json:"mark"`
}

// VoteGetResponse defines the structure of the response when
// getting a single translation vote.
type VoteGetResponse struct {
	Data *Vote `json:"data"`
}

// VotesListResponse defines the structure of the response when
// getting a list of translation votes.
type VotesListResponse struct {
	Data []*VoteGetResponse `json:"data"`
}

// VotesListOptions specifies the optional parameters to the
// StringTranslationsService.ListVotes method.
type VotesListOptions struct {
	// String Identifier.
	// Note: Must be used together with `languageId`.
	StringID int `json:"stringId,omitempty"`
	// Language Identifier.
	// Note: Must be used together with `stringId`.
	LanguageID string `json:"languageId,omitempty"`
	// Translation Identifier.
	// Note: If specified, `stringId` and `languageId` are ignored.
	TranslationID int `json:"translationId,omitempty"`
	// File Identifier.
	// Note: Must be used together with `languageId`.
	FileID int `json:"fileId,omitempty"`
	// Label Identifiers.
	// Example: labelIds=1,2,3,4,5
	LabelIDs []int `json:"labelIds,omitempty"`
	// Exclude Label Identifiers.
	ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the VotesListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *VotesListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.StringID > 0 {
		v.Add("stringId", fmt.Sprintf("%d", o.StringID))
	}
	if o.LanguageID != "" {
		v.Add("languageId", o.LanguageID)
	}
	if o.TranslationID > 0 {
		v.Add("translationId", fmt.Sprintf("%d", o.TranslationID))
	}
	if o.FileID > 0 {
		v.Add("fileId", fmt.Sprintf("%d", o.FileID))
	}
	if len(o.LabelIDs) > 0 {
		v.Add("labelIds", JoinSlice(o.LabelIDs))
	}
	if len(o.ExcludeLabelIDs) > 0 {
		v.Add("excludeLabelIds", JoinSlice(o.ExcludeLabelIDs))
	}

	return v, len(v) > 0
}

// VoteType represents a translation vote type.
type VoteType string

const (
	// VoteTypeUp is an upvote translation.
	VoteTypeUp VoteType = "up"
	// VoteTypeDown is a downvote translation.
	VoteTypeDown VoteType = "down"
)

// VoteAddRequest defines the structure of the request
// to add a translation vote.
type VoteAddRequest struct {
	// Enum: up, down.
	Mark VoteType `json:"mark"`
	// Translation Identifier.
	TranslationID int `json:"translationId"`
}

// Validate checks if the VotesAddRequest is valid.
// It implements the crowdin.RequestValidator interface.
func (r *VoteAddRequest) Validate() error {
	if r == nil {
		return errors.New("request cannot be nil")
	}
	if r.Mark != VoteTypeUp && r.Mark != VoteTypeDown {
		return fmt.Errorf("invalid vote type: %q", r.Mark)
	}
	if r.TranslationID == 0 {
		return errors.New("translation ID is required")
	}
	return nil
}
