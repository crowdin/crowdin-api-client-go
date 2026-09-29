package model

import (
	"errors"
	"fmt"
	"net/url"
)

// TranslationProgress defines the structure of a translations status progress.
type TranslationProgress struct {
	// Words statistics. Keys: total, translated, preTranslateAppliedTo, approved.
	Words map[string]int `json:"words"`
	// Phrases statistics. Keys: total, translated, preTranslateAppliedTo, approved.
	Phrases             map[string]int  `json:"phrases"`
	TranslationProgress int             `json:"translationProgress"`
	ApprovalProgress    int             `json:"approvalProgress"`
	QAChecksStatus      *QAChecksStatus `json:"qaChecksStatus,omitempty"`
	LanguageID          *string         `json:"languageId,omitempty"`
	BranchID            *int            `json:"branchId,omitempty"`
	FileID              *int            `json:"fileId,omitempty"`
	Language            *Language       `json:"language,omitempty"`
	Etag                *string         `json:"etag,omitempty"`
}

// QAChecksStatus represents the QA checks status statistics of a translation progress.
type QAChecksStatus struct {
	Total      int `json:"total"`
	InProgress int `json:"inProgress"`
	Passed     int `json:"passed"`
	Failed     int `json:"failed"`
}

// TranslationStatusProgressResponse defines the structure of a response when getting
// a translation status progress (for a branch, directory, file, language or project).
type TranslationProgressResponse struct {
	Data []struct {
		Data *TranslationProgress `json:"data"`
	} `json:"data"`
}

// ProjectProgressListOptions specifies the optional parameters to the
// TranslationStatusService.GetProjectProgress method.
type ProjectProgressListOptions struct {
	// Filter progress by Language Identifier.
	LanguageIDs []string `json:"languageIds,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of ProjectProgressListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ProjectProgressListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if len(o.LanguageIDs) > 0 {
		v.Add("languageIds", JoinSlice(o.LanguageIDs))
	}

	return v, len(v) > 0
}

// QACheck represents a QA check issue.
type QACheck struct {
	StringID              int    `json:"stringId"`
	LanguageID            string `json:"languageId"`
	Category              string `json:"category"`
	CategoryDescription   string `json:"categoryDescription"`
	Validation            string `json:"validation"`
	ValidationDescription string `json:"validationDescription"`
	PluralID              int    `json:"pluralId"`
	Text                  string `json:"text"`
}

// QAChecksResponse defines the structure of a response
// when getting a list of QA check issues.
type QAChecksResponse struct {
	Data []struct {
		Data *QACheck `json:"data"`
	} `json:"data"`
}

// QACheckListOptions specifies the optional parameters to the
// TranslationStatusService.ListQAChecks method.
type QACheckListOptions struct {
	// Defines category of QA check issue. It can be one category or a list of comma-separated ones.
	// Example: category=variables,tags
	// Enum: empty, variables, tags, punctuation, symbol_register, spaces, size, special_symbols,
	//       wrong_translation, spellcheck, icu
	Category []string `json:"category,omitempty"`
	// Defines the QA check issue validation type. It can be one validation type or a list
	// of comma-separated ones. Example: validation=capitalize_check,punctuation_check
	// Enum: empty_string_check, empty_suggestion_check, max_length_check, tags_check,
	//       mismatch_ids_check, cdata_check, specials_symbols_check, leading_newlines_check,
	//       trailing_newlines_check, leading_spaces_check, trailing_spaces_check, multiple_spaces_check,
	//       custom_blocked_variables_check, highest_priority_custom_variables_check,
	//       highest_priority_variables_check, c_variables_check, python_variables_check,
	//       rails_variables_check, java_variables_check, dot_net_variables_check, twig_variables_check,
	//       php_variables_check, freemarker_variables_check, lowest_priority_variable_check,
	//       lowest_priority_custom_variables_check, punctuation_check, spaces_before_punctuation_check,
	//       spaces_after_punctuation_check, non_breaking_spaces_check, capitalize_check,
	//       multiple_uppercase_check, parentheses_check, entities_check, escaped_quotes_check,
	//       wrong_translation_issue_check, spellcheck, icu_check
	Validation []string `json:"validation,omitempty"`
	// Filter progress by Language Identifier.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// Filter the collection by the specified task identifier.
	TaskID int `json:"taskId,omitempty"`
	// Filter the collection by the specified file identifier.
	FileID int `json:"fileId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of QACheckListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *QACheckListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if len(o.Category) > 0 {
		v.Add("category", JoinSlice(o.Category))
	}
	if len(o.Validation) > 0 {
		v.Add("validation", JoinSlice(o.Validation))
	}
	if len(o.LanguageIDs) > 0 {
		v.Add("languageIds", JoinSlice(o.LanguageIDs))
	}
	if o.TaskID > 0 {
		v.Add("taskId", fmt.Sprintf("%d", o.TaskID))
	}
	if o.FileID > 0 {
		v.Add("fileId", fmt.Sprintf("%d", o.FileID))
	}

	return v, len(v) > 0
}

// QACheckRevalidationRequest defines the structure of a request
// to revalidate QA checks.
type QACheckRevalidationRequest struct {
	// QA check categories to revalidate. If not specified,
	// all active categories will be checked. Enum: terms, ai.
	QACheckCategories []string `json:"qaCheckCategories,omitempty"`
	// Language IDs to revalidate. If not specified, all languages will be checked.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// If true, only languages with failed QA checks will be revalidated.
	// Default: false.
	FailedOnly *bool `json:"failedOnly,omitempty"`
	// External QA check IDs to revalidate.
	// Note: Available only for Crowdin Enterprise.
	ExternalQACheckIDs []int `json:"externalQaCheckIds,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *QACheckRevalidationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

type (
	// QACheckRevalidation represents a QA checks revalidation job.
	QACheckRevalidation struct {
		Identifier string                         `json:"identifier"`
		Status     string                         `json:"status"`
		Progress   int                            `json:"progress"`
		Attributes *QACheckRevalidationAttributes `json:"attributes"`
		CreatedAt  string                         `json:"createdAt"`
		UpdatedAt  string                         `json:"updatedAt"`
		StartedAt  *string                        `json:"startedAt"`
		FinishedAt *string                        `json:"finishedAt"`
	}

	// QACheckRevalidationAttributes represents the attributes
	// of a QA checks revalidation job.
	QACheckRevalidationAttributes struct {
		LanguageIDs        []string `json:"languageIds"`
		QACheckCategories  []string `json:"qaCheckCategories"`
		FailedOnly         bool     `json:"failedOnly"`
		ExternalQACheckIDs []int    `json:"externalQaCheckIds,omitempty"`
	}

	// QACheckRevalidationResponse defines the structure of a response
	// when revalidating QA checks or getting a revalidation status.
	QACheckRevalidationResponse struct {
		Data *QACheckRevalidation `json:"data"`
	}
)

// QACheckValidateRequest defines the structure of a single item of the
// request to validate text by QA checks.
type QACheckValidateRequest struct {
	// String Identifier.
	StringID int `json:"stringId"`
	// Language Identifier.
	LanguageID string `json:"languageId"`
	// Translation text.
	Text string `json:"text"`
	// Plural form. Enum: zero, one, two, few, many, other.
	PluralCategoryName string `json:"pluralCategoryName,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *QACheckValidateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StringID == 0 {
		return errors.New("stringId is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Text == "" {
		return errors.New("text is required")
	}

	return nil
}

// QACheckValidation represents a QA check issue found when validating text.
type QACheckValidation struct {
	StringID              int    `json:"stringId"`
	LanguageID            string `json:"languageId"`
	Category              string `json:"category"`
	CategoryDescription   string `json:"categoryDescription"`
	Validation            string `json:"validation"`
	ValidationDescription string `json:"validationDescription"`
	PluralID              int    `json:"pluralId"`
	PluralCategoryName    string `json:"pluralCategoryName"`
	Text                  string `json:"text"`
	Translation           string `json:"translation"`
}

// QACheckValidationsResponse defines the structure of a response
// when validating text by QA checks.
type QACheckValidationsResponse struct {
	Data []struct {
		Data *QACheckValidation `json:"data"`
	} `json:"data"`
}
