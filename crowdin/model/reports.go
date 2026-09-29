package model

import (
	"errors"
	"fmt"
	"net/url"
)

// ReportFormat represents the format of a report
// file to export.
type ReportFormat string

const (
	ReportFormatXLSX ReportFormat = "xlsx"
	ReportFormatCSV  ReportFormat = "csv"
	ReportFormatJSON ReportFormat = "json"
)

// ReportScopeType represents the scope type of a report.
type ReportScopeType string

const (
	ReportScopeTypeProject      ReportScopeType = "project"
	ReportScopeTypeOrganization ReportScopeType = "organization"
	ReportScopeTypeGroup        ReportScopeType = "group"
)

// ReportUnit represents a unit of a report.
type ReportUnit string

const (
	ReportUnitStrings         ReportUnit = "strings"
	ReportUnitWords           ReportUnit = "words"
	ReportUnitChars           ReportUnit = "chars"
	ReportUnitCharsWithSpaces ReportUnit = "chars_with_spaces"
	// ReportUnitHours is used by hourly report settings templates.
	ReportUnitHours ReportUnit = "hours"
)

// ReportMode represents the mode of a report.
type ReportMode string

const (
	ReportModeTranslations ReportMode = "translations"
	ReportModeApprovals    ReportMode = "approvals"
	ReportModeVotes        ReportMode = "votes"
)

// ReportName represents the name of a report.
type ReportName string

const (
	ReportCostsEstimationPostEditing  ReportName = "costs-estimation-pe"
	ReportTransactionCostsPostEditing ReportName = "translation-costs-pe"
	ReportContributionRawData         ReportName = "contribution-raw-data"
	ReportTopMembers                  ReportName = "top-members"
	ReportTranslatorAccuracy          ReportName = "translator-accuracy"
	ReportPreTranslateAccuracy        ReportName = "pre-translate-accuracy"
	ReportSourceContentUpdates        ReportName = "source-content-updates"
	ReportProjectMembers              ReportName = "project-members"
	ReportEditorIssues                ReportName = "editor-issues"
	ReportQACheckIssues               ReportName = "qa-check-issues"
	ReportSavingActivity              ReportName = "saving-activity"
	ReportTranslationActivity         ReportName = "translation-activity"
	ReportTimeSpent                   ReportName = "time-spent"
	// ReportTaskUsage is available for the Enterprise client only.
	ReportTaskUsage ReportName = "task-usage"

	// Deprecated: Use ReportPreTranslateAccuracy instead.
	ReportPreTranslateEfficiency ReportName = "pre-translate-efficiency"

	// Organization reports.
	ReportGroupTranslationCostsPostEditing ReportName = "group-translation-costs-pe"
	ReportGroupTopMembers                  ReportName = "group-top-members"
	ReportGroupTaskUsage                   ReportName = "group-task-usage"
	ReportGroupQACheckIssues               ReportName = "group-qa-check-issues"
	ReportGroupTranslationActivity         ReportName = "group-translation-activity"
	ReportGroupSourceContentUpdates        ReportName = "group-source-content-updates"
	ReportGroupTimeSpent                   ReportName = "group-time-spent"
	ReportGroupPreTranslateAccuracy        ReportName = "group-pre-translate-accuracy"
	ReportGroupTranslatorAccuracy          ReportName = "group-translator-accuracy"
	ReportGroupSavingActivity              ReportName = "group-saving-activity"
)

// ReportArchive represents a report archive.
type ReportArchive struct {
	ID        int    `json:"id"`
	ScopeType string `json:"scopeType"`
	ScopeID   int    `json:"scopeId"`
	UserID    int    `json:"userId"`
	Name      string `json:"name"`
	WebURL    string `json:"webUrl"`
	Scheme    any    `json:"scheme"`
	Status    string `json:"status,omitempty"`
	Progress  int    `json:"progress,omitempty"`
	CreatedAt string `json:"createdAt"`
}

// ReportArchiveResponse defines the structure of a response
// when getting a report archive.
type ReportArchiveResponse struct {
	Data *ReportArchive `json:"data"`
}

// ReportArchiveListResponse defines the structure of a response
// when getting a list of report archives.
type ReportArchiveListResponse struct {
	Data []*ReportArchiveResponse `json:"data"`
}

// ReportArchivesListOptions specifies the optional parameters to
// the ReportsService.ListArchives method.
type ReportArchivesListOptions struct {
	// Filter only project report archives.
	// Enum: project, organization, group.
	ScopeType ReportScopeType `json:"scopeType,omitempty"`
	// Filter archives by specific scope id.
	// [Enterprise client] Use only if scopeType set to group or project.
	ScopeID int `json:"scopeId,omitempty"`
	// Filter archives by user identifier.
	UserID int `json:"userId,omitempty"`
	// Filter archives by task identifier.
	TaskID int `json:"taskId,omitempty"`
	// Filter archives by name.
	Name string `json:"name,omitempty"`
	// Archive date from in UTC, ISO 8601.
	DateFrom string `json:"dateFrom,omitempty"`
	// Archive date to in UTC, ISO 8601.
	DateTo string `json:"dateTo,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the options.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ReportArchivesListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.ScopeType != "" {
		v.Add("scopeType", string(o.ScopeType))
	}
	if o.ScopeID != 0 {
		v.Add("scopeId", fmt.Sprintf("%d", o.ScopeID))
	}
	if o.UserID != 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}
	if o.TaskID != 0 {
		v.Add("taskId", fmt.Sprintf("%d", o.TaskID))
	}
	if o.Name != "" {
		v.Add("name", o.Name)
	}
	if o.DateFrom != "" {
		v.Add("dateFrom", o.DateFrom)
	}
	if o.DateTo != "" {
		v.Add("dateTo", o.DateTo)
	}

	return v, len(v) > 0
}

type (
	// ReportStatus represents the status of a generated
	// or archived report.
	ReportStatus struct {
		Identifier string                 `json:"identifier"`
		Status     string                 `json:"status"`
		Progress   int                    `json:"progress"`
		Attributes ReportStatusAttributes `json:"attributes"`
		CreatedAt  string                 `json:"createdAt"`
		UpdatedAt  string                 `json:"updatedAt"`
		StartedAt  string                 `json:"startedAt"`
		FinishedAt string                 `json:"finishedAt"`
	}

	// ReportStatusAttributes represents the attributes of
	// a report status.
	ReportStatusAttributes struct {
		Format     string `json:"format"`
		ReportName string `json:"reportName"`
		Schema     any    `json:"schema"`
	}
)

// ReportStatusResponse defines the structure of a response
// when getting a report status.
type ReportStatusResponse struct {
	Data *ReportStatus `json:"data"`
}

// ExportReportArchiveRequest defines the structure of a request to
// export a report archive.
type ExportReportArchiveRequest struct {
	// Export file format.
	// Enum: xlsx, csv, json. Default: xlsx.
	Format ReportFormat `json:"format,omitempty"`
}

type (
	// ReportBaseRates defines the base rates for a report.
	ReportBaseRates struct {
		// Applies to all languages by default.
		FullTranslation float64 `json:"fullTranslation,omitempty"`
		Proofread       float64 `json:"proofread,omitempty"`
		// Hourly rate. Used by the time spent reports and
		// hourly report settings templates.
		Hourly float64 `json:"hourly,omitempty"`
	}

	// ReportIndividualRates defines the individual rates for a report.
	// Custom rates for certain languages or users.
	ReportIndividualRates struct {
		LanguageIDs     []string `json:"languageIds,omitempty"`
		UserIDs         []int    `json:"userIds,omitempty"`
		FullTranslation float64  `json:"fullTranslation,omitempty"`
		Proofread       float64  `json:"proofread,omitempty"`
		// Hourly rate. Used by the time spent reports and
		// hourly report settings templates.
		Hourly float64 `json:"hourly,omitempty"`
	}

	// ReportNetRateSchemes defines the net rate schemes for a report.
	// Percentage paid of full translation rate.
	ReportNetRateSchemes struct {
		// Match type enum: "perfect", "100", "99-82", "81-60".
		TMMatch []ReportNetRateSchemeMatch `json:"tmMatch,omitempty"`
		// Match type enum: "100", "99-82", "81-60".
		MTMatch []ReportNetRateSchemeMatch `json:"mtMatch,omitempty"`
		// AI match. Match type: "100" or a percentage range (e.g., "99-82").
		// Note: If this field is not filled in, the schema will use the MT match values.
		AIMatch []ReportNetRateSchemeMatch `json:"aiMatch,omitempty"`
		// Match type enum: "100", "99-82".
		SuggestionMatch []ReportNetRateSchemeMatch `json:"suggestionMatch,omitempty"`
	}

	// ReportNetRateSchemeMatch defines the match type and price
	// for a net rate scheme.
	ReportNetRateSchemeMatch struct {
		// Match type, %. Enum: perfect, 100, 99-82, 81-60.
		MatchType string `json:"matchType,omitempty"`
		// Price, %.
		Price float64 `json:"price,omitempty"`
	}
)

// ReportGenerateRequest defines the structure of a request to
// generate a report.
type ReportGenerateRequest struct {
	// Report name.
	Name ReportName `json:"name"`
	// Schema for the report generation request.
	// Can be one of the following types:
	//  - CostsEstimationPostEditingSchema
	//  - TransactionCostsPostEditingSchema
	//  - TopMembersSchema
	//  - ContributionRawDataSchema
	//  - TranslatorAccuracySchema
	//  - PreTranslateAccuracySchema
	//  - PreTranslateEfficiencySchema (Deprecated)
	//  - SourceContentUpdatesSchema
	//  - ProjectMembersSchema
	//  - EditorIssuesSchema
	//  - QACheckIssuesSchema
	//  - SavingActivitySchema
	//  - TranslationActivitySchema
	//  - TimeSpentSchema
	//  - TaskUsageSchema (Enterprise only)
	Schema ReportSchema `json:"schema"`
}

// ReportSchema is an interface that defines the schema
// for a report generation request.
type ReportSchema interface {
	ValidateSchema() error
}

type (
	// CostsEstimationPostEditingSchema defines the schema for the costs
	// estimation post-editing report.
	CostsEstimationPostEditingSchema struct {
		// Report unit.
		// Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Report currency.
		// Enum: USD, EUR, JPY, GBP, AUD, CAD, CHF, CNY, SEK, NZD, MXN,
		// SGD, HKD, NOK, KRW, TRY, RUB, INR, BRL, ZAR, GEL, UAH, DDK
		Currency string `json:"currency,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Base rates.
		BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
		// Individual rates (Custom rates for certain languages or users).
		IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
		// Net Rate Schemes (Percentage paid of full translation rate).
		// Note: A new translation will be included in the report at the lowest rate
		// if multiple scheme categories can be applied to the translation.
		NetRateSchemes *ReportNetRateSchemes `json:"netRateSchemes,omitempty"`
		// Calculate internal matches. Default: false.
		CalculateInternalMatches *bool `json:"calculateInternalMatches,omitempty"`
		// Include pre-translated strings. Default: false.
		IncludePreTranslatedStrings *bool `json:"includePreTranslatedStrings,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// List of label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Defines which strings include in report.
		// Enum: strings_with_label, strings_without_label. Default: strings_with_label.
		LabelIncludeType string `json:"labelIncludeType,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Workflow step identifier (Enterprise only).
		WorkflowStepID int `json:"workflowStepId,omitempty"`

		// Task Identifier.
		// Used to generate report by task.
		//
		// Deprecated: use TaskIDs instead.
		TaskID int `json:"taskId,omitempty"`
		// Task identifiers.
		// Used to generate report by tasks.
		TaskIDs []int `json:"taskIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// TransactionCostsPostEditingSchema defines the schema for the transaction
	// costs post-editing report.
	TransactionCostsPostEditingSchema struct {
		// Report unit.
		// Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Report currency.
		// Enum: USD, EUR, JPY, GBP, AUD, CAD, CHF, CNY, SEK, NZD, MXN,
		// SGD, HKD, NOK, KRW, TRY, RUB, INR, BRL, ZAR, GEL, UAH, DDK
		Currency string `json:"currency,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Base rates.
		BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
		// Individual rates (Custom rates for certain languages or users).
		IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
		// Net Rate Schemes (Percentage paid of full translation rate).
		// Note: A new translation will be included in the report at the lowest rate
		// if multiple scheme categories can be applied to the translation.
		NetRateSchemes *ReportNetRateSchemes `json:"netRateSchemes,omitempty"`
		// Exclude approvals when the same user has made translations for the string.
		//
		// Deprecated: the API no longer supports this field.
		ExcludeApprovalsForEditedTranslations *bool `json:"excludeApprovalsForEditedTranslations,omitempty"`
		// Approvals are treated as submitting an identical translation at the
		// 100% match rate of the corresponding category.
		UseCategoryBasedProofreadRates *bool `json:"useCategoryBasedProofreadRates,omitempty"`
		// Calculations are based on the edit distance between the TM match
		// and final translation.
		UseTmEditDistance *bool `json:"useTmEditDistance,omitempty"`
		// Grouping parameter.
		// Enum: user, language. Default: user.
		GroupBy string `json:"groupBy,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// User Identifier for which the report should be generated.
		UserIDs []int `json:"userIds,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// List of label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Defines which strings include in report.
		// Enum: strings_with_label, strings_without_label. Default: strings_with_label.
		LabelIncludeType string `json:"labelIncludeType,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Workflow step identifier (Enterprise only).
		WorkflowStepID int `json:"workflowStepId,omitempty"`

		// Task Identifier.
		// Used to generate report by task.
		//
		// Deprecated: use TaskIDs instead.
		TaskID int `json:"taskId,omitempty"`
		// Task identifiers.
		// Used to generate report by tasks.
		TaskIDs []int `json:"taskIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// TopMembersSchema defines the schema for the top members report.
	TopMembersSchema struct {
		// Defines report unit.
		// Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
	}

	// ContributionRawDataSchema defines the schema for the contribution
	// raw data report.
	ContributionRawDataSchema struct {
		// Report mode. Enum: translations, approvals, votes.
		Mode ReportMode `json:"mode"`
		// Report unit. Enum: strings, words, chars, chars_with_spaces.
		// Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Task Identifier.
		// Used to generate report by task.
		TaskID int `json:"taskId,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// User Identifier for which the report should be generated.
		UserID string `json:"userId,omitempty"`
		// Column names to include in the report.
		// Enum: userId, languageId, stringId, translationId, fileId, filePath,
		// pluralForm, sourceStringTextHash, mtEngine, mtId, tmName, tmId,
		// preTranslated, tmMatch, mtMatch, suggestionMatch, sourceUnits,
		// targetUnits, createdAt, updatedAt, mark.
		Columns []string `json:"columns,omitempty"`
		// List of TM identifiers.
		TMIDs []int `json:"tmIds,omitempty"`
		// List of MT identifiers.
		MTIDs []int `json:"mtIds,omitempty"`
		// List of AI prompt identifiers.
		AIPromptIDs []int `json:"aiPromptIds,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// TranslatorAccuracySchema defines the schema for the translator
	// accuracy report.
	TranslatorAccuracySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces.
		// Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Split into categories by edit distance.
		//
		// Deprecated: use MatchScoreCategories instead.
		PostEditingCategories []string `json:"postEditingCategories,omitempty"`
		// Split into categories by match score. Ranges should be in
		// descending order (e.g., 100-90, 89-80).
		MatchScoreCategories []string `json:"matchScoreCategories,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// List of label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Defines which strings include in report.
		// Enum: strings_with_label, strings_without_label.
		LabelIncludeType string `json:"labelIncludeType,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// PreTranslateAccuracySchema defines the schema for pre translate
	// accuracy report.
	PreTranslateAccuracySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces.
		// Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Split into categories by edit distance.
		//
		// Deprecated: use MatchScoreCategories instead.
		PostEditingCategories []string `json:"postEditingCategories,omitempty"`
		// Split into categories by match score. Ranges should be in
		// descending order (e.g., 100-90, 89-80).
		MatchScoreCategories []string `json:"matchScoreCategories,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Task Identifier for which the report should be generated.
		TaskID int `json:"taskId,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// List of label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Defines which strings include in report.
		// Enum: strings_with_label, strings_without_label.
		LabelIncludeType string `json:"labelIncludeType,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// PreTranslateEfficiencySchema defines the schema for pre translate
	// efficiency report.
	// Deprecated: Use PreTranslateAccuracySchema instead.
	PreTranslateEfficiencySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces.
		// Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Split into categories by edit distance.
		PostEditingCategories []string `json:"postEditingCategories,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// SourceContentUpdatesSchema defines the schema for the source content updates report.
	SourceContentUpdatesSchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// ProjectMembersSchema defines the schema for the project members report.
	ProjectMembersSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// EditorIssuesSchema defines the schema for the editor issues report.
	EditorIssuesSchema struct {
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Issue type filter.
		IssueType string `json:"issueType,omitempty"`
	}

	// QACheckIssuesSchema defines the schema for the QA check issues report.
	QACheckIssuesSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// SavingActivitySchema defines the schema for the saving activity report.
	SavingActivitySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report mode. Enum: currency, relative.
		Mode string `json:"mode,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of file identifiers.
		FileIDs []int `json:"fileIds,omitempty"`
		// List of directory identifiers.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// List of branch identifiers.
		BranchIDs []int `json:"branchIds,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
	}

	// TranslationActivitySchema defines the schema for the translation activity report.
	TranslationActivitySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
	}

	// TimeSpentSchema defines the schema for the time spent report.
	TimeSpentSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Grouping parameter. Enum: user, language, task.
		GroupBy string `json:"groupBy,omitempty"`
		// Base rates. Only the `hourly` rate is used.
		BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
		// Individual rates. Only the `hourly` rate is used.
		IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
		// Task type. Enum: 0 - translate, 1 - proofread,
		// 2 - translate by vendor, 3 - proofread by vendor.
		TypeTasks *int `json:"typeTasks,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Task identifiers for which the report should be generated.
		TaskIDs []int `json:"taskIds,omitempty"`
		// Workflow step identifier (Enterprise only).
		WorkflowStepID int `json:"workflowStepId,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// TaskUsageSchema defines the schema for the project task usage report
	// (Enterprise only).
	TaskUsageSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format"`
		// Report type. Enum: workload, created-vs-resolved, performance, time, cost.
		Type string `json:"type"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Grouping parameter. Enum: user, language, type.
		GroupBy string `json:"groupBy,omitempty"`
		// Task type. Enum: 0 - translate, 1 - proofread,
		// 2 - translate by vendor, 3 - proofread by vendor.
		TypeTasks *int `json:"typeTasks,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Task creator identifier.
		CreatorID int `json:"creatorId,omitempty"`
		// Task assignee identifier.
		AssigneeID int `json:"assigneeId,omitempty"`
		// Words count from (used with type `time`).
		WordsCountFrom int `json:"wordsCountFrom,omitempty"`
		// Words count to (used with type `time`).
		WordsCountTo int `json:"wordsCountTo,omitempty"`
		// Task statuses to filter by (used with type `cost`).
		// Enum: todo, in_progress, done, closed, review.
		Statuses []string `json:"statuses,omitempty"`
	}
)

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ReportGenerateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Schema == nil {
		return errors.New("schema is required")
	}

	return r.Schema.ValidateSchema()
}

// ValidateSchema implements the ReportSchema interface and checks if the
// CostsEstimationPostEditing schema is valid.
func (r *CostsEstimationPostEditingSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TransactionCostsPostEditing schema is valid.
func (r *TransactionCostsPostEditingSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TopMembers schema is valid.
func (r *TopMembersSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// ContributionRawData schema is valid.
func (r *ContributionRawDataSchema) ValidateSchema() error {
	if r.Mode == "" {
		return errors.New("mode is required")
	}

	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TranslatorAccuracySchema schema is valid.
func (r *TranslatorAccuracySchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// PreTranslateAccuracy schema is valid.
func (r *PreTranslateAccuracySchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// PreTranslateEfficiency schema is valid.
func (r *PreTranslateEfficiencySchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// SourceContentUpdates schema is valid.
func (r *SourceContentUpdatesSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// ProjectMembers schema is valid.
func (r *ProjectMembersSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// EditorIssues schema is valid.
func (r *EditorIssuesSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// QACheckIssues schema is valid.
func (r *QACheckIssuesSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// SavingActivity schema is valid.
func (r *SavingActivitySchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TranslationActivity schema is valid.
func (r *TranslationActivitySchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TimeSpent schema is valid.
func (r *TimeSpentSchema) ValidateSchema() error {
	return nil
}

// ValidateSchema implements the ReportSchema interface and checks if the
// TaskUsage schema is valid.
func (r *TaskUsageSchema) ValidateSchema() error {
	if r.Format == "" {
		return errors.New("format is required")
	}
	if r.Type == "" {
		return errors.New("type is required")
	}

	return nil
}

// GroupReportGenerateRequest defines the structure of a request to
// generate a group or organization report.
type GroupReportGenerateRequest struct {
	// Report name.
	Name ReportName `json:"name"`
	// Schema for the group report generation request.
	// One of the following types:
	//  - GroupTransactionCostsPostEditingSchema
	//  - GroupTopMembersSchema
	//  - GroupTaskUsageSchema
	//  - GroupQACheckIssuesSchema
	//  - GroupTranslationActivitySchema
	//  - GroupSourceContentUpdatesSchema
	//  - GroupTimeSpentSchema
	//  - GroupPreTranslateAccuracySchema
	//  - GroupTranslatorAccuracySchema
	//  - GroupSavingActivitySchema
	Schema ReportGroupSchema `json:"schema"`
}

// ReportGroupSchema is an interface that defines the schema
// for a group report generation request.
//
// Schema can be one of the following types:
//   - GroupTransactionCostsPostEditingSchema
//   - GroupTopMembersSchema
//   - GroupTaskUsageSchema
//   - GroupQACheckIssuesSchema
//   - GroupTranslationActivitySchema
//   - GroupSourceContentUpdatesSchema
//   - GroupTimeSpentSchema
//   - GroupPreTranslateAccuracySchema
//   - GroupTranslatorAccuracySchema
//   - GroupSavingActivitySchema
type ReportGroupSchema interface {
	ValidateGroupSchema() error
}

type (
	// GroupTransactionCostsPostEditingSchema defines the schema for the group
	// translation costs post-editing report.
	GroupTransactionCostsPostEditingSchema struct {
		// Project Identifier for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Report unit.
		// Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Report currency.
		// Enum: USD, EUR, JPY, GBP, AUD, CAD, CHF, CNY, SEK, NZD, MXN,
		// SGD, HKD, NOK, KRW, TRY, RUB, INR, BRL, ZAR, GEL, UAH, DDK
		Currency string `json:"currency,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Base rates.
		BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
		// Individual rates (Custom rates for certain languages or users).
		IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
		// Net Rate Schemes (Percentage paid of full translation rate).
		// Note: A new translation will be included in the report at the lowest rate
		// if multiple scheme categories can be applied to the translation.
		NetRateSchemes *ReportNetRateSchemes `json:"netRateSchemes,omitempty"`
		// Exclude approvals when the same user has made translations for the string.
		//
		// Deprecated: the API no longer supports this field.
		ExcludeApprovalsForEditedTranslations *bool `json:"excludeApprovalsForEditedTranslations,omitempty"`
		// Approvals are treated as submitting an identical translation at the
		// 100% match rate of the corresponding category.
		UseCategoryBasedProofreadRates *bool `json:"useCategoryBasedProofreadRates,omitempty"`
		// Calculations are based on the edit distance between the TM match
		// and final translation.
		UseTmEditDistance *bool `json:"useTmEditDistance,omitempty"`
		// Grouping parameter.
		// Enum: user, language, project. Default: user.
		GroupBy string `json:"groupBy,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// User Identifier for which the report should be generated.
		UserIDs []int `json:"userIds,omitempty"`
		// Task identifiers. Used to generate report by tasks.
		TaskIDs []int `json:"taskIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// GroupTopMembersSchema defines the schema for the group top members report.
	GroupTopMembersSchema struct {
		// Project Identifier for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Defines report unit.
		// Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Export file format.
		// Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
	}

	// GroupTaskUsageSchema defines the schema for the group task usage report.
	GroupTaskUsageSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report type filter.
		Type string `json:"type,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Group results by the specified attribute.
		GroupBy string `json:"groupBy,omitempty"`
		// Task type identifier filter.
		TypeTasks int `json:"typeTasks,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Task creator identifier filter.
		CreatorID int `json:"creatorId,omitempty"`
		// Task assignee identifier filter.
		AssigneeID int `json:"assigneeId,omitempty"`
		// Words count from (used with type `time`).
		WordsCountFrom int `json:"wordsCountFrom,omitempty"`
		// Words count to (used with type `time`).
		WordsCountTo int `json:"wordsCountTo,omitempty"`
		// Task statuses to filter by (used with type `cost`).
		// Enum: todo, in_progress, done, closed, review.
		Statuses []string `json:"statuses,omitempty"`
	}

	// GroupQACheckIssuesSchema defines the schema for the group QA check issues report.
	GroupQACheckIssuesSchema struct {
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// GroupTranslationActivitySchema defines the schema for the group translation activity report.
	GroupTranslationActivitySchema struct {
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// List of user identifiers for filtering.
		UserIDs []int `json:"userIds,omitempty"`
	}

	// GroupSourceContentUpdatesSchema defines the schema for the group
	// source content updates report.
	GroupSourceContentUpdatesSchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// GroupTimeSpentSchema defines the schema for the group time spent report.
	GroupTimeSpentSchema struct {
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Grouping parameter. Enum: user, language, task, project.
		GroupBy string `json:"groupBy,omitempty"`
		// Base rates. Only the `hourly` rate is used.
		BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
		// Individual rates. Only the `hourly` rate is used.
		IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
		// Task type. Enum: 0 - translate, 1 - proofread,
		// 2 - translate by vendor, 3 - proofread by vendor.
		TypeTasks *int `json:"typeTasks,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Task identifiers.
		TaskIDs []int `json:"taskIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// GroupPreTranslateAccuracySchema defines the schema for the group
	// pre-translation accuracy report.
	GroupPreTranslateAccuracySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Split into categories by match score. Ranges should be in
		// descending order (e.g., 100-90, 89-80).
		MatchScoreCategories []string `json:"matchScoreCategories,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Task identifiers. Used to generate report by tasks.
		TaskIDs []int `json:"taskIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// GroupTranslatorAccuracySchema defines the schema for the group
	// translator accuracy report.
	GroupTranslatorAccuracySchema struct {
		// Split into categories by match score. Ranges should be in
		// descending order (e.g., 100-90, 89-80).
		MatchScoreCategories []string `json:"matchScoreCategories,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// List of user identifiers.
		UserIDs []int `json:"userIds,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// If true, the report will not be saved to the archive.
		SkipArchiving *bool `json:"skipArchiving,omitempty"`
	}

	// GroupSavingActivitySchema defines the schema for the group saving activity report.
	GroupSavingActivitySchema struct {
		// Report unit. Enum: strings, words, chars, chars_with_spaces. Default: words.
		Unit ReportUnit `json:"unit,omitempty"`
		// Project identifiers for which the report should be generated.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Export file format. Enum: xlsx, csv, json. Default: xlsx.
		Format ReportFormat `json:"format,omitempty"`
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Language Identifier for which the report should be generated.
		LanguageID string `json:"languageId,omitempty"`
		// Report mode. Enum: relative, currency.
		// Note: `currency` mode is available only if the group has
		// the savingsReportSettingsTemplateId set.
		Mode string `json:"mode,omitempty"`
	}
)

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *GroupReportGenerateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Schema == nil {
		return errors.New("schema is required")
	}

	return r.Schema.ValidateGroupSchema()
}

// ValidateGroupSchema checks if the GroupCostsEstimationPostEditing schema is valid.
func (r *GroupTransactionCostsPostEditingSchema) ValidateGroupSchema() error {
	if r.BaseRates == nil {
		return errors.New("baseRates is required")
	}
	if r.IndividualRates == nil {
		return errors.New("individualRates is required")
	}
	if r.NetRateSchemes == nil {
		return errors.New("netRateSchemes is required")
	}

	return nil
}

// ValidateGroupSchema checks if the GroupTopMembers schema is valid.
func (r *GroupTopMembersSchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupTaskUsage schema is valid.
func (r *GroupTaskUsageSchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupQACheckIssues schema is valid.
func (r *GroupQACheckIssuesSchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupTranslationActivity schema is valid.
func (r *GroupTranslationActivitySchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupSourceContentUpdates schema is valid.
func (r *GroupSourceContentUpdatesSchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupTimeSpent schema is valid.
func (r *GroupTimeSpentSchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupPreTranslateAccuracy schema is valid.
func (r *GroupPreTranslateAccuracySchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupTranslatorAccuracy schema is valid.
func (r *GroupTranslatorAccuracySchema) ValidateGroupSchema() error {
	return nil
}

// ValidateGroupSchema checks if the GroupSavingActivity schema is valid.
func (r *GroupSavingActivitySchema) ValidateGroupSchema() error {
	return nil
}

// ReportSettingsTemplate represents a report settings template.
type ReportSettingsTemplate struct {
	ID        int                          `json:"id"`
	Name      string                       `json:"name"`
	Currency  string                       `json:"currency"`
	Unit      string                       `json:"unit"`
	Config    ReportSettingsTemplateConfig `json:"config"`
	CreatedAt string                       `json:"createdAt"`
	UpdatedAt string                       `json:"updatedAt"`
	IsPublic  bool                         `json:"isPublic"`
	IsGlobal  *bool                        `json:"isGlobal,omitempty"`

	ProjectID int `json:"projectId,omitempty"`
	GroupID   int `json:"groupId,omitempty"`
}

// ReportSettingsTemplateResponse defines the structure of a response
// when getting a report settings template.
type ReportSettingsTemplateResponse struct {
	Data *ReportSettingsTemplate `json:"data"`
}

// ReportSettingsTemplateListResponse defines the structure of a response
// when getting a list of report settings templates.
type ReportSettingsTemplateListResponse struct {
	Data []*ReportSettingsTemplateResponse `json:"data"`
}

// ReportSettingsTemplatesListOptions specifies the optional parameters to
// the ReportsService.ListTemplates method.
type ReportSettingsTemplatesListOptions struct {
	// [Enterprise client] Project Identifier.
	ProjectID int `json:"projectId,omitempty"`
	// [Enterprise client] Group Identifier.
	GroupID int `json:"groupId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of the options.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ReportSettingsTemplatesListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.ProjectID != 0 {
		v.Add("projectId", fmt.Sprintf("%d", o.ProjectID))
	}
	if o.GroupID != 0 {
		v.Add("groupId", fmt.Sprintf("%d", o.GroupID))
	}

	return v, len(v) > 0
}

// ReportSettingsTemplateAddRequest defines the structure of a request to
// add a report settings template.
type ReportSettingsTemplateAddRequest struct {
	// Template name.
	Name string `json:"name"`
	// Report currency.
	// Enum: USD, EUR, JPY, GBP, AUD, CAD, CHF, CNY, SEK, NZD, MXN,
	// SGD, HKD, NOK, KRW, TRY, RUB, INR, BRL, ZAR, GEL, UAH, DDK
	Currency string `json:"currency"`
	// Report unit.
	// Enum: strings, words, chars, chars_with_spaces
	Unit ReportUnit `json:"unit"`
	// Report config.
	Config *ReportSettingsTemplateConfig `json:"config"`
	// Report visibility.
	IsPublic *bool `json:"isPublic,omitempty"`
	// Report global visibility.
	IsGlobal *bool `json:"isGlobal,omitempty"`

	// [Enterprise client] Project Identifier.
	ProjectID int `json:"projectId,omitempty"`
	// [Enterprise client] Group Identifier.
	GroupID int `json:"groupId,omitempty"`
}

// ReportSettingsTemplateUpdateRequest defines the structure of a request to
// create a report settings template.
type ReportSettingsTemplateConfig struct {
	// Base rates.
	BaseRates *ReportBaseRates `json:"baseRates,omitempty"`
	// Individual rates (Custom rates for certain languages or users).
	IndividualRates []*ReportIndividualRates `json:"individualRates,omitempty"`
	// Net Rate Schemes (Percentage paid of full translation rate).
	// Note: A new translation will be included in the report at the lowest rate
	// if multiple scheme categories can be applied to the translation.
	NetRateSchemes *ReportNetRateSchemes `json:"netRateSchemes,omitempty"`
	// Calculate internal matches.
	CalculateInternalMatches *bool `json:"calculateInternalMatches,omitempty"`
	// Include pre-translated strings.
	IncludePreTranslatedStrings *bool `json:"includePreTranslatedStrings,omitempty"`
	// Approvals are treated as submitting an identical translation at the
	// 100% match rate of the corresponding category.
	UseCategoryBasedProofreadRates *bool `json:"useCategoryBasedProofreadRates,omitempty"`
	// Calculations are based on the edit distance between the TM match
	// and final translation.
	UseTmEditDistance *bool `json:"useTmEditDistance,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ReportSettingsTemplateAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Currency == "" {
		return errors.New("currency is required")
	}
	if r.Unit == "" {
		return errors.New("unit is required")
	}
	if r.Config == nil {
		return errors.New("config is required")
	}
	if r.Config.BaseRates == nil || len(r.Config.IndividualRates) == 0 {
		return errors.New("config fields are required")
	}
	// Hourly templates don't use net rate schemes.
	if r.Unit != ReportUnitHours && r.Config.NetRateSchemes == nil {
		return errors.New("config fields are required")
	}

	return nil
}
