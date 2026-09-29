package model

import (
	"errors"
	"fmt"
	"net/url"
)

type (
	// Project represents a Crowdin project.
	Project struct {
		ID                   int         `json:"id"`
		GroupID              int         `json:"groupId,omitempty"`
		Type                 int         `json:"type"`
		UserID               int         `json:"userId"`
		SourceLanguageID     string      `json:"sourceLanguageId"`
		TargetLanguageIDs    []string    `json:"targetLanguageIds"`
		LanguageAccessPolicy string      `json:"languageAccessPolicy"`
		Name                 string      `json:"name"`
		Cname                string      `json:"cname,omitempty"`
		Identifier           string      `json:"identifier"`
		Description          string      `json:"description"`
		Visibility           string      `json:"visibility"`
		Logo                 string      `json:"logo"`
		IsExternal           bool        `json:"isExternal,omitempty"`
		ExternalType         string      `json:"externalType,omitempty"`
		WorkflowID           int         `json:"workflowId,omitempty"`
		HasCrowdsourcing     bool        `json:"hasCrowdsourcing,omitempty"`
		PublicDownloads      bool        `json:"publicDownloads"`
		CreatedAt            string      `json:"createdAt"`
		UpdatedAt            string      `json:"updatedAt"`
		LastActivity         string      `json:"lastActivity"`
		SourceLanguage       *Language   `json:"sourceLanguage"`
		TargetLanguages      []*Language `json:"targetLanguages"`
		WebURL               string      `json:"webUrl"`
		Fields               any         `json:"fields,omitempty"`

		ExternalProjectID               *int   `json:"externalProjectId,omitempty"`
		ExternalOrganizationID          *int   `json:"externalOrganizationId,omitempty"`
		PublicURL                       string `json:"publicUrl,omitempty"`
		SavingsReportSettingsTemplateID int    `json:"savingsReportSettingsTemplateId,omitempty"`

		// Deprecated: use GlossaryAccessOption instead.
		GlossaryAccess                  bool                       `json:"glossaryAccess,omitempty"`
		ClientOrganizationID            int                        `json:"clientOrganizationId,omitempty"`
		TranslateDuplicates             int                        `json:"translateDuplicates,omitempty"`
		TagsDetection                   int                        `json:"tagsDetection,omitempty"`
		GlossaryAccessOption            string                     `json:"glossaryAccessOption,omitempty"`
		IsMTAllowed                     bool                       `json:"isMtAllowed,omitempty"`
		TaskBasedAccessControl          bool                       `json:"taskBasedAccessControl,omitempty"`
		TaskReviewerIDs                 []int                      `json:"taskReviewerIds,omitempty"`
		HiddenStringsProofreadersAccess bool                       `json:"hiddenStringsProofreadersAccess,omitempty"`
		AutoSubstitution                bool                       `json:"autoSubstitution,omitempty"`
		ExportTranslatedOnly            bool                       `json:"exportTranslatedOnly,omitempty"`
		SkipUntranslatedStrings         bool                       `json:"skipUntranslatedStrings,omitempty"`
		ExportApprovedOnly              bool                       `json:"exportApprovedOnly,omitempty"`
		ExportWithMinApprovalsCount     int                        `json:"exportWithMinApprovalsCount,omitempty"`
		ExportStringsThatPassedWorkflow bool                       `json:"exportStringsThatPassedWorkflow,omitempty"`
		AutoTranslateDialects           bool                       `json:"autoTranslateDialects,omitempty"`
		UseGlobalTM                     bool                       `json:"useGlobalTm,omitempty"`
		TMContextType                   string                     `json:"tmContextType,omitempty"`
		ShowTMSuggestionsDialects       bool                       `json:"showTmSuggestionsDialects,omitempty"`
		TmApprovedSuggestionsOnly       bool                       `json:"tmApprovedSuggestionsOnly,omitempty"`
		IsSuspended                     bool                       `json:"isSuspended,omitempty"`
		QACheckIsActive                 bool                       `json:"qaCheckIsActive,omitempty"`
		QAApprovalsCount                int                        `json:"qaApprovalsCount,omitempty"`
		QACheckCategories               map[string]bool            `json:"qaCheckCategories,omitempty"`
		QAChecksIgnorableCategories     map[string]bool            `json:"qaChecksIgnorableCategories,omitempty"`
		CustomQACheckIDs                []int                      `json:"customQaCheckIds,omitempty"`
		ExternalQACheckIDs              []int                      `json:"externalQaCheckIds,omitempty"`
		LanguageMapping                 map[string]LanguageMapping `json:"languageMapping,omitempty"`
		DelayedWorkflowStart            bool                       `json:"delayedWorkflowStart,omitempty"`
		NotificationSettings            *NotificationSettings      `json:"notificationSettings,omitempty"`
		DefaultTMID                     int                        `json:"defaultTmId,omitempty"`
		DefaultGlossaryID               int                        `json:"defaultGlossaryId,omitempty"`
		AssignedTMs                     map[int]map[string]int     `json:"assignedTms,omitempty"`
		AssignedGlossaries              []int                      `json:"assignedGlossaries,omitempty"`
		AssignedStyleGuides             []int                      `json:"assignedStyleGuides,omitempty"`
		TMPenalties                     any                        `json:"tmPenalties,omitempty"`
		NormalizePlaceholder            bool                       `json:"normalizePlaceholder,omitempty"`
		TMPreTranslate                  *ProjectTMPreTranslate     `json:"tmPreTranslate,omitempty"`
		MTPreTranslate                  *ProjectMTPreTranslate     `json:"mtPreTranslate,omitempty"`
		AiPreTranslate                  *ProjectAiPreTranslate     `json:"aiPreTranslate,omitempty"`
		EditorSuggestionAiPromptID      *int                       `json:"editorSuggestionAiPromptId,omitempty"`
		AlignmentActionAiPromptID       *int                       `json:"alignmentActionAiPromptId,omitempty"`
		QACheckActionAiPromptID         *int                       `json:"qaCheckActionAiPromptId,omitempty"`
		ContextReviewAiPromptID         *int                       `json:"contextReviewAiPromptId,omitempty"`
		SaveMetaInfoInSource            bool                       `json:"saveMetaInfoInSource,omitempty"`
		SkipUntranslatedFiles           bool                       `json:"skipUntranslatedFiles,omitempty"`
		InContext                       bool                       `json:"inContext,omitempty"`
		InContextProcessHiddenStrings   bool                       `json:"inContextProcessHiddenStrings,omitempty"`
		InContextPseudoLanguageID       string                     `json:"inContextPseudoLanguageId,omitempty"`
		InContextPseudoLanguage         *Language                  `json:"inContextPseudoLanguage,omitempty"`
	}

	ProjectTMPenalties struct {
		AutoSubstitution int `json:"autoSubstitution,omitempty"`
		TMPriority       struct {
			Priority int `json:"priority,omitempty"`
			Penalty  int `json:"penalty,omitempty"`
		} `json:"tmPriority,omitempty"`
		MultipleTranslations int `json:"multipleTranslations,omitempty"`
		TimeSinceLastUsage   struct {
			Months  int `json:"months,omitempty"`
			Penalty int `json:"penalty,omitempty"`
		} `json:"timeSinceLastUsage,omitempty"`
		TimeSinceLastModified struct {
			Months  int `json:"months,omitempty"`
			Penalty int `json:"penalty,omitempty"`
		} `json:"timeSinceLastModified,omitempty"`
	}

	ProjectTMPreTranslate struct {
		Enabled *bool `json:"enabled,omitempty"`
		// Enum: "all", "perfectMatchOnly", "exceptAutoSubstituted", "perfectMatchApprovedOnly", "none".
		AutoApproveOption string `json:"autoApproveOption,omitempty"`
		// Enum: "perfect", "100".
		MinimumMatchRatio string `json:"minimumMatchRatio,omitempty"`
	}

	ProjectMTPreTranslate struct {
		Enabled *bool        `json:"enabled,omitempty"`
		MTs     []ProjectMTs `json:"mts,omitempty"`
	}

	ProjectAiPreTranslate struct {
		Enabled   *bool             `json:"enabled,omitempty"`
		AiPrompts []ProjectAiPrompt `json:"aiPrompts,omitempty"`
	}

	ProjectMTs struct {
		MTID int `json:"mtId,omitempty"`
		// Specify an array of languageIds to use specific languages, or use the string all
		// to include all supported languages.
		// Retrieve languageIds via the `List Supported Languages` endpoint
		LanguageIDs []string `json:"languageIds,omitempty"`
	}

	ProjectAiPrompt struct {
		AiPromptID int `json:"aiPromptId,omitempty"`
		// Specify an array of languageIds to use specific languages, or use the string all
		// to include all supported languages.
		// Retrieve languageIds via the List Supported Languages endpoint
		LanguageIDs []string `json:"languageIds,omitempty"`
	}
)

// ProjectsListOptions specifies the optional parameters to the ProjectsService.List method.
type ProjectsListOptions struct {
	ListOptions

	// Order projects by.
	// Enum: id, name, identifier, description, createdAt, updatedAt, lastActivity. Default: id.
	// Example: orderBy=createdAt desc,name,id.
	OrderBy string `json:"orderBy,omitempty"`
	// User Identifier.
	UserID int `json:"userId,omitempty"`
	// Projects with Manager Access. Enum: 0, 1. Default: 0.
	HasManagerAccess *int `json:"hasManagerAccess,omitempty"`
	// Set type to 0 to get all file based projects. Enum: 0, 1.
	Type *int `json:"type,omitempty"`
	// Filter projects by name.
	Filter string `json:"filter,omitempty"`
	// Group Identifier (Enterprise only).
	// Note: Set 0 to see items of root group.
	GroupID *int `json:"groupId,omitempty"`
}

// Values returns the url.Values representation of ProjectsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ProjectsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.UserID > 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}
	if o.HasManagerAccess != nil &&
		(*o.HasManagerAccess == 0 || *o.HasManagerAccess == 1) {
		v.Add("hasManagerAccess", fmt.Sprintf("%d", *o.HasManagerAccess))
	}
	if o.Type != nil && (*o.Type == 0 || *o.Type == 1) {
		v.Add("type", fmt.Sprintf("%d", *o.Type))
	}
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}
	if o.GroupID != nil {
		v.Add("groupId", fmt.Sprintf("%d", *o.GroupID))
	}
	return v, len(v) > 0
}

// ProjectGetResponse defines the structure of a response when retrieving a project.
type ProjectsGetResponse struct {
	Data *Project `json:"data"`
}

// GroupListResponse defines the structure of a response when getting a list of groups.
type ProjectsListResponse struct {
	Data       []*ProjectsGetResponse `json:"data"`
	Pagination *Pagination            `json:"pagination"`
}

// ProjectsAddRequest defines the structure of a request to add a project.
type ProjectsAddRequest struct {
	// Project Name.
	Name string `json:"name"`
	// Project Identifier.
	Identifier string `json:"identifier,omitempty"`
	// Source Language Identifier.
	SourceLanguageID string `json:"sourceLanguageId"`
	// Target Languages Identifiers.
	TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
	// Defines how users can join the project. Enum: open, private. Default: private.
	// open – anyone can join the project
	// private – only invited users can join the project
	Visibility string `json:"visibility,omitempty"`
	// Defines access to project languages. Enum: open, moderate. Default: open.
	// open – each project user can access all project languages
	// moderate – users should join each project language separately
	LangAccessPolicy string `json:"languageAccessPolicy,omitempty"`
	// Custom domain name.
	Cname string `json:"cname,omitempty"`
	// Project description.
	Description string `json:"description,omitempty"`
	// Values available: 0 - Auto, 1 - Count tags, 1 - Skip tags. Default: 0.
	TagsDetection *int `json:"tagsDetection,omitempty"`
	// Allows machine translations (Microsoft Translator, Google Translate) be visible
	// for translators in the Editor. Default: true.
	IsMTAllowed *bool `json:"isMtAllowed,omitempty"`
	// Allow project members work with tasks they assigned to, even if they do not have
	// full access to the language. Default: false.
	TaskBasedAccessControl *bool `json:"taskBasedAccessControl,omitempty"`
	// Allows auto-substitution. Default: true.
	AutoSubstitution *bool `json:"autoSubstitution,omitempty"`
	// Automatically fill in regional dialects. Default: false.
	// If true, all untranslated strings in regional dialects (e.g. Argentine Spanish)
	// will automatically include translations completed in the primary language (e.g. Spanish).
	AutoTranslateDialects *bool `json:"autoTranslateDialects,omitempty"`
	// Allows translators to download source files to their machines and upload translations back into the project.
	// Project owner and managers can always download sources and upload translations. Default: true.
	PublicDownloads *bool `json:"publicDownloads,omitempty"`
	// Allows proofreaders to work with hidden strings.
	// Project owner and managers can always access hidden strings. Default: true.
	HiddenStringsProofreadersAccess *bool `json:"hiddenStringsProofreadersAccess,omitempty"`
	// If true - machine translations from connected MT engines (e.g. Microsoft Translator, Google Translate)
	// will appear as suggestions in the Editor. Default: false.
	// Note: If your organization plan is free or opensource - default value of this one will be true
	UseGlobalTM *bool `json:"useGlobalTm,omitempty"`
	// If true - show primary language TM suggestions for dialects if there are no dialect-specific ones. Default: true.
	ShowTMSuggestionsDialects *bool `json:"showTmSuggestionsDialects,omitempty"`
	// If true - only approved suggestions will be saved to the project default TM.
	TmApprovedSuggestionsOnly *bool `json:"tmApprovedSuggestionsOnly,omitempty"`
	// Defines whether to skip untranslated strings.
	SkipUntranslatedStrings *bool `json:"skipUntranslatedStrings,omitempty"`
	// Defines whether to export only approved strings.
	ExportApprovedOnly *bool `json:"exportApprovedOnly,omitempty"`
	// If true - QA checks are active. Default: true.
	QACheckIsActive *bool `json:"qaCheckIsActive,omitempty"`
	// Acceptable categories are: empty, size, tags, spaces, variables, punctuation, symbolRegister,
	// specialSymbols, wrongTranslation, spellcheck, icu, terms, duplicate, ftl, android, numbers,
	// ai, outdated, mdx, unifiedPlaceholders.
	QACheckCategories map[string]bool `json:"qaCheckCategories,omitempty"`
	// Acceptable categories are: empty, size, tags, spaces, variables, punctuation, symbolRegister,
	// specialSymbols, wrongTranslation, spellcheck, icu, terms, duplicate, ftl, android, numbers,
	// ai, outdated, mdx, unifiedPlaceholders.
	QAChecksIgnorableCategories map[string]bool `json:"qaChecksIgnorableCategories,omitempty"`
	// Language Mapping.
	LanguageMapping map[string]LanguageMapping `json:"languageMapping,omitempty"`
	// Allow project members to manage glossary terms.
	// The project owner and managers always can add and edit terms. Default: false.
	//
	// Deprecated: use GlossaryAccessOption instead.
	GlossaryAccess *bool `json:"glossaryAccess,omitempty"`
	// Defines the glossary access level for project members.
	// Enum: readOnly, fullAccess, manageDrafts. Default: readOnly.
	GlossaryAccessOption string `json:"glossaryAccessOption,omitempty"`
	// Enable the transformation of the placeholders to the unified format to improve the work with TM suggestions.
	NormalizePlaceholder *bool `json:"normalizePlaceholder,omitempty"`
	// Notification Settings.
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	// TM perfect match searching mode. Enum: "segmentContext" "auto" "prevAndNextSegment". Default: "segmentContext".
	// segmentContext - searching by context.
	// auto - context search for key-value formats and segment search for others.
	// prevAndNextSegment - search by previous and next segment.
	TMContextType  string                 `json:"tmContextType,omitempty"`
	TMPreTranslate *ProjectTMPreTranslate `json:"tmPreTranslate,omitempty"`
	MTPreTranslate *ProjectMTPreTranslate `json:"mtPreTranslate,omitempty"`
	AiPreTranslate *ProjectAiPreTranslate `json:"aiPreTranslate,omitempty"`
	// AI Prompt ID to be used as prompt for Assist action.
	//
	// Deprecated: the Assist AI prompt action was removed from the API.
	AssistActionAiPromptID int `json:"assistActionAiPromptId,omitempty"`
	// AI Prompt ID to be used as prompt for Editor Suggestion action.
	EditorSuggestionAiPromptID int `json:"editorSuggestionAiPromptId,omitempty"`
	// AI Prompt ID to be used as prompt for Alignment action (Enterprise only).
	AlignmentActionAiPromptID int `json:"alignmentActionAiPromptId,omitempty"`
	// AI Prompt ID to be used as prompt for QA Check action.
	QACheckActionAiPromptID int `json:"qaCheckActionAiPromptId,omitempty"`
	// AI Prompt ID to be used as prompt for Context Review action.
	ContextReviewAiPromptID int `json:"contextReviewAiPromptId,omitempty"`
	// Savings Report Settings Template Identifier.
	SavingsReportSettingsTemplateID int `json:"savingsReportSettingsTemplateId,omitempty"`
	// Style Guide identifiers to assign to the project.
	AssignedStyleGuides []int `json:"assignedStyleGuides,omitempty"`
	// Translation Memory ID.
	// Default: null
	DefaultTMID int `json:"defaultTmId,omitempty"`
	// Glossary ID.
	// Default: null
	DefaultGlossaryID int `json:"defaultGlossaryId,omitempty"`
	// Context and max.length added in Crowdin will be visible in the downloaded files.
	SaveMetaInfoInSource *bool `json:"saveMetaInfoInSource,omitempty"`
	// Defines the project type. Use 0 for a file-based project and 1 for a string-based project.
	// Enum: 0, 1. Default: 0.
	Type *int `json:"type,omitempty"`
	// Defines whether to export only translated file.
	SkipUntranslatedFiles *bool `json:"skipUntranslatedFiles,omitempty"`
	// Enable In-Context translations. Default: false.
	// Note: Must be used together with `inContextPseudoLanguageId`
	InContext *bool `json:"inContext,omitempty"`
	// Export hidden strings via pseudo-language. Default: true.
	// Note: If true - hidden strings included in the pseudo-language archive will be translatable via In-Context.
	InContextProcessHiddenStrings *bool `json:"inContextProcessHiddenStrings,omitempty"`
	// In-Context pseudo-language id.
	// Note: Must be different from project source and target languages
	InContextPseudoLanguageID string `json:"inContextPseudoLanguageId,omitempty"`

	// Workflow Template Step Identifier.
	ID int `json:"id,omitempty"`
	// Workflow Template Identifier.
	TemplateID int `json:"templateId,omitempty"`
	// Workflow Template Steps Configuration.
	// Note. Must be used together with `templateId`. Can't be used with
	//       `vendorId`, `mtEngineId` in same request.
	Steps []*WorkflowTemplateStep `json:"steps,omitempty"`
	// Group Identifier.
	GroupID int `json:"groupId,omitempty"`
	// Specify Vendor Identifier, if no Vendor is assigned to Workflow step yet.
	VendorID int `json:"vendorId,omitempty"`
	// Specify Machine Translation engine Identifier, if no MT engine is
	// assigned to Workflow step yet.
	MTEngineID int `json:"mtEngineId,omitempty"`
	// Enum 0, 1, 2, 3, 4, 5. Default: 0.
	//  0 - Show – translators will translate each instance separately,
	//  1 - Hide (regular detection) – all duplicates will share the same translation
	//  2 - Show, but auto-translate them,
	//  3 - Show within a version branch (regular detection) - duplicates will be hidden only
	//      between versions branches
	//  4 - Hide (strict detection) – all duplicates will share the same translation
	//  5 - Show within a version branch (strict detection) - duplicates will be hidden only
	//      between versions branches
	TranslateDuplicates *int `json:"translateDuplicates,omitempty"`
	// Delay workflow start after project creation. Default: false.
	DelayedWorkflowStart *bool `json:"delayedWorkflowStart,omitempty"`
	// Defines whether to export only approved strings.
	// Note: value greater than 0 can't be used with `exportStringsThatPassedWorkflow=true`
	//       in same request.
	ExportWithMinApprovalsCount *int `json:"exportWithMinApprovalsCount,omitempty"`
	// Defines whether to export only strings that passed workflow.
	// Note: `true` value can't be used with `exportWithMinApprovalsCount>0` in same request
	// or in projects without an assigned workflow.
	ExportStringsThatPassedWorkflow *bool `json:"exportStringsThatPassedWorkflow,omitempty"`
	// Clear QA checks for translations with specific number of approvals. Default: 1.
	QAApprovalsCount *int `json:"qaApprovalsCount,omitempty"`
	// Custom QA checks identifiers.
	CustomQACheckIDs []int `json:"customQaCheckIds,omitempty"`
	// External QA checks identifiers (Enterprise only).
	ExternalQACheckIDs []int `json:"externalQaCheckIds,omitempty"`
	// Task reviewer identifiers (Enterprise only).
	TaskReviewerIDs []int `json:"taskReviewerIds,omitempty"`
	// MT Engine Identifier.
	MTID int `json:"mtId,omitempty"`
	// Fields.
	Fields map[string]any `json:"fields,omitempty"`
	// Target Languages Identifiers.
	Languages []string `json:"languages,omitempty"`
}

// LanguageMapping represents a project language mapping.
type LanguageMapping struct {
	Name                 string `json:"name"`
	TwoLettersCode       string `json:"two_letters_code"`
	ThreeLettersCode     string `json:"three_letters_code"`
	Locale               string `json:"locale"`
	LocaleWithUnderscore string `json:"locale_with_underscore"`
	AndroidCode          string `json:"android_code"`
	OSXCode              string `json:"osx_code"`
	OSXLocale            string `json:"osx_locale"`
}

// NotificationSettings represents a project notification settings.
type NotificationSettings struct {
	// Notify translators about new strings. Default: false.
	TranslatorNewStrings *bool `json:"translatorNewStrings,omitempty"`
	// Notify project managers about new strings. Default: false.
	ManagerNewStrings *bool `json:"managerNewStrings,omitempty"`
	// Notify project managers about language translation/validation completion.
	// Default: false.
	ManagerLanguageCompleted *bool `json:"managerLanguageCompleted,omitempty"`
}

// Validate checks if the add request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ProjectsAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.SourceLanguageID == "" {
		return errors.New("sourceLanguageId is required")
	}
	return nil
}

// ProjectsFileFormatSettings represents a Crowdin project file format settings.
type ProjectsFileFormatSettings struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	Format     string         `json:"format"`
	Extensions []string       `json:"extensions"`
	Settings   map[string]any `json:"settings"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
}

// ProjectsFileFormatSettingsResponse defines the structure of a response when
// retrieving a project file format settings.
type ProjectsFileFormatSettingsResponse struct {
	Data *ProjectsFileFormatSettings `json:"data"`
}

// ProjectsFileFormatSettingsListResponse defines the structure of a response when
// getting a list of project file format settings.
type ProjectsFileFormatSettingsListResponse struct {
	Data []*ProjectsFileFormatSettingsResponse `json:"data"`
}

type FileFormatSettings interface {
	ValidateSettings() error
}

type (
	// ProjectsFileFormatSettingsRequest defines the structure of a request
	// to add a project file format settings.
	ProjectsAddFileFormatSettingsRequest struct {
		// Defines file format.
		Format string `json:"format"`
		// Defines file format settings.
		Settings FileFormatSettings `json:"settings"`
	}

	CommonFileFormatSettings struct {
		// Defines whether to split long texts into smaller text segments. Default: true.
		// Important! This option disables the possibility to upload existing translations for XML files when enabled.
		ContentSegmentation *bool `json:"contentSegmentation,omitempty"`
		// Storage identifier of the SRX segmentation rules file. Default: null.
		SRXStorageID *int `json:"srxStorageId,omitempty"`
		// File format export pattern. Default: null.
		// Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	PropertyFileFormatSettings struct {
		// File export pattern. Default: null.
		// Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols
		ExportPattern *string `json:"exportPattern,omitempty"`
		// Enum: 0, 1, 2, 3. Default: 1.
		// 0 - Do not escape single quote.
		// 1 - Escape single quote by another single quote.
		// 2 - Escape single quote by a backslash.
		// 3 - Escape single quote by another single quote only in strings containing variables ({0}).
		EscapeQuotes *int `json:"escapeQuotes,omitempty"`
		// Enum: 0, 1. Default: 1.
		// Defines whether any special characters (=, :, ! and #) should be escaped by backslash in exported translations.
		// You can add escape_special_characters per-file option. *
		// Acceptable values are: 0, 1. Default is 0.
		// 0 - Do not escape special characters.
		// 1 - Escape special characters by a backslash.
		EscapeSpecialCharacters *int `json:"escapeSpecialCharacters,omitempty"`
	}

	XMLFileFormatSettings struct {
		// Defines whether to translate texts placed inside the tags. Default: true.
		TranslateContent *bool `json:"translateContent,omitempty"`
		// Defines whether to translate tags attributes. Default: true.
		TranslateAttributes *bool `json:"translateAttributes,omitempty"`
		// This is an array of strings, where each item is the XPaths to DOM element
		// that should be imported. Default: []. Enum: "/path/to/node", "/path/to/attribute[@attr]",
		// "//node", "//[@attr]", "nodeone/nodetwo", "/nodeone//nodetwo", "//node[@attr]"
		TranslatableElements []string `json:"translatableElements,omitempty"`
		// Defines whether to split long texts into smaller text segments. Default: true.
		// Important! This option disables the possibility to upload existing translations for XML files when enabled.
		ContentSegmentation *bool `json:"contentSegmentation,omitempty"`
		// Storage Identifier of the SRX segmentation rules file. Default: null.
		SRXStorageID *int `json:"srxStorageId,omitempty"`
		// File format export pattern. Defines file name and path in resulting translations bundle.
		// Default: null. Note: Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
		// Specify the inline tags.
		InlineTags []string `json:"inlineTags,omitempty"`
	}

	HTMLFileFormatSettings struct {
		CommonFileFormatSettings
		// Specify CSS selectors for elements that should not be imported.
		ExcludedElements []string `json:"excludedElements,omitempty"`
		// Specify the inline tags.
		InlineTags []string `json:"inlineTags,omitempty"`
	}

	AdocFileFormatSettings struct {
		CommonFileFormatSettings
		// Skip Include Directives. Default: false.
		ExcludeIncludeDirectives *bool `json:"excludeIncludeDirectives,omitempty"` // Default: false
	}

	MDXV1FileFormatSettings struct {
		CommonFileFormatSettings
		// Specify elements that should not be imported.
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Defines whether to import code blocks. Default: false.
		ExcludeCodeBlocks *bool `json:"excludeCodeBlocks,omitempty"`
		// Default: "mdx_v1". Enum: "mdx_v1", "mdx_v2"
		Type string `json:"type,omitempty"`
		// Defines the marker used for bold text. Enum: "asterisk", "underscore". Default: "asterisk".
		StrongMarker string `json:"strongMarker,omitempty"`
		// Defines the marker used for italic text. Enum: "asterisk", "underscore". Default: "underscore".
		EmphasisMarker string `json:"emphasisMarker,omitempty"`
		// Defines the bullet used for unordered lists. Enum: "asterisks", "plus", "dash". Default: "dash".
		UnorderedListBullet string `json:"unorderedListBullet,omitempty"`
		// Defines the table column width. Enum: "consolidate", "evenly_distribute_cells".
		// Default: "evenly_distribute_cells".
		TableColumnWidth string `json:"tableColumnWidth,omitempty"`
	}

	MDXV2FileFormatSettings struct {
		CommonFileFormatSettings
		// Specify elements that should not be imported.
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
		// Defines whether to import code blocks. Default: false.
		ExcludeCodeBlocks *bool `json:"excludeCodeBlocks,omitempty"`
		// Defines the marker used for bold text. Enum: "asterisk", "underscore". Default: "asterisk".
		StrongMarker string `json:"strongMarker,omitempty"`
		// Defines the marker used for italic text. Enum: "asterisk", "underscore". Default: "underscore".
		EmphasisMarker string `json:"emphasisMarker,omitempty"`
		// Defines the bullet used for unordered lists. Enum: "asterisks", "plus", "dash". Default: "dash".
		UnorderedListBullet string `json:"unorderedListBullet,omitempty"`
		// Defines the table column width. Enum: "consolidate", "evenly_distribute_cells".
		// Default: "evenly_distribute_cells".
		TableColumnWidth string `json:"tableColumnWidth,omitempty"`
	}

	// MDFileFormatSettings defines the Markdown file format settings.
	MDFileFormatSettings struct {
		CommonFileFormatSettings
		// Specify the inline tags.
		InlineTags []string `json:"inlineTags,omitempty"`
		// Defines the marker used for bold text. Enum: "asterisk", "underscore". Default: "asterisk".
		StrongMarker string `json:"strongMarker,omitempty"`
		// Defines the marker used for italic text. Enum: "asterisk", "underscore". Default: "underscore".
		EmphasisMarker string `json:"emphasisMarker,omitempty"`
		// Defines the bullet used for unordered lists. Enum: "asterisks", "plus", "dash". Default: "dash".
		UnorderedListBullet string `json:"unorderedListBullet,omitempty"`
		// Defines the table column width. Enum: "consolidate", "evenly_distribute_cells".
		// Default: "evenly_distribute_cells".
		TableColumnWidth string `json:"tableColumnWidth,omitempty"`
		// Defines the quotes used in front matter. Enum: "auto", "single", "double". Default: "auto".
		FrontMatterQuotes string `json:"frontMatterQuotes,omitempty"`
	}

	// FMHTMLFileFormatSettings defines the Front Matter HTML file format settings.
	FMHTMLFileFormatSettings struct {
		CommonFileFormatSettings
		// Specify the inline tags.
		InlineTags []string `json:"inlineTags,omitempty"`
		// Specify CSS selectors for elements that should not be imported.
		ExcludedElements []string `json:"excludedElements,omitempty"`
		// Specify front matter elements that should not be imported.
		ExcludedFrontMatterElements []string `json:"excludedFrontMatterElements,omitempty"`
	}

	// WebXMLFileFormatSettings defines the Web XML file format settings.
	WebXMLFileFormatSettings struct {
		CommonFileFormatSettings
		// Specify the inline tags.
		InlineTags []string `json:"inlineTags,omitempty"`
	}

	// IDMLFileFormatSettings defines the IDML file format settings.
	IDMLFileFormatSettings struct {
		CommonFileFormatSettings
		// Defines whether to inline hyperlink text. Default: false.
		InlineHyperlinkText *bool `json:"inlineHyperlinkText,omitempty"`
	}

	// VDFFileFormatSettings defines the VDF file format settings.
	VDFFileFormatSettings struct {
		// Defines whether to convert ICU messages. Default: true.
		ConvertICU *bool `json:"convertIcu,omitempty"`
		// Defines whether to add the gender argument. Default: false.
		AddGenderArgument *bool `json:"addGenderArgument,omitempty"`
		// File format export pattern. Defines file name and path in resulting translations bundle.
		// Default: null. Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	DocxFileFormatSettings struct {
		CommonFileFormatSettings
		// When checked, strips additional formatting tags related to text spacing. Default: false.
		// Note: Works only for files with the following extensions: *.docx, *.dotx, *.docm, *.dotm,
		// *.xlsx, *.xltx, *.xlsm, *.xltm, *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		CleanTagsAggressively *bool `json:"cleanTagsAggressively,omitempty"`
		// When checked, exposes hidden text for translation. Default: false.
		// Note: Works only for files with the following extensions: *.docx, *.dotx, *.docm, *.dotm.
		TranslateHiddenText *bool `json:"translateHiddenText,omitempty"`
		// When checked, exposes hidden hyperlinks for translation. Default: false.
		// Note: Works only for files with the following extensions: *.docx, *.dotx,
		// *.docm, *.dotm, *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		TranslateHyperlinkUrls *bool `json:"translateHyperlinkUrls,omitempty"`
		// When checked, exposes hidden rows and columns for translation. Default: false.
		// Note: Works only for files with the following extensions: *.xlsx, *.xltx, *.xlsm, *.xltm.
		TranslateHiddenRowsAndColumns *bool `json:"translateHiddenRowsAndColumns,omitempty"`
		// When checked, expose slide notes for translation. Default: true.
		// Note: Works only for files with the following extensions: *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		ImportNotes *bool `json:"importNotes,omitempty"`
		// When checked, exposes hidden slides for translation. Default: false.
		// Note: Works only for files with the following extensions: *.pptx, *.potx, *.ppsx, *.pptm, *.potm, *.ppsm.
		ImportHiddenSlides *bool `json:"importHiddenSlides,omitempty"`
		// Defines whether to translate document properties. Default: false.
		TranslateDocProperties *bool `json:"translateDocProperties,omitempty"`
		// Defines whether to translate comments. Default: true.
		TranslateComments *bool `json:"translateComments,omitempty"`
		// Defines whether to ignore whitespace styles. Default: false.
		IgnoreWhitespaceStyles *bool `json:"ignoreWhitespaceStyles,omitempty"`
		// Defines whether to add tab as a character. Default: false.
		AddTabAsCharacter *bool `json:"addTabAsCharacter,omitempty"`
		// Defines whether to add line separator as a character. Default: false.
		AddLineSeparatorAsCharacter *bool `json:"addLineSeparatorAsCharacter,omitempty"`
		// Defines the line separator replacement. Default: a newline character.
		LineSeparatorReplacement *string `json:"lineSeparatorReplacement,omitempty"`
		// Defines whether to replace the no-break hyphen tag. Default: false.
		ReplaceNoBreakHyphenTag *bool `json:"replaceNoBreakHyphenTag,omitempty"`
		// Defines whether to ignore the soft hyphen tag. Default: false.
		IgnoreSoftHyphenTag *bool `json:"ignoreSoftHyphenTag,omitempty"`
		// Complex field definitions to extract. Default: [].
		ComplexFieldDefinitionsToExtract []string `json:"complexFieldDefinitionsToExtract,omitempty"`
		// Defines whether to translate Word headers and footers. Default: true.
		TranslateWordHeadersFooters *bool `json:"translateWordHeadersFooters,omitempty"`
		// Defines whether to translate Word graphic names. Default: true.
		TranslateWordGraphicName *bool `json:"translateWordGraphicName,omitempty"`
		// Defines whether to translate Word graphic descriptions. Default: false.
		TranslateWordGraphicDescription *bool `json:"translateWordGraphicDescription,omitempty"`
		// Defines whether to ignore Word font colors. Default: false.
		IgnoreWordFontColors *bool `json:"ignoreWordFontColors,omitempty"`
		// Minimum Word font color ignorance threshold. Default: "".
		WordFontColorsMinIgnoranceThreshold *string `json:"wordFontColorsMinIgnoranceThreshold,omitempty"`
		// Maximum Word font color ignorance threshold. Default: "".
		WordFontColorsMaxIgnoranceThreshold *string `json:"wordFontColorsMaxIgnoranceThreshold,omitempty"`
		// Word styles to exclude. Default: [].
		ExcludeWordStyles []string `json:"excludeWordStyles,omitempty"`
		// Defines whether to translate Word content in exclude style mode. Default: true.
		TranslateWordInExcludeStyleMode *bool `json:"translateWordInExcludeStyleMode,omitempty"`
		// Word highlight colors. Default: [].
		WordHighlightColors []string `json:"wordHighlightColors,omitempty"`
		// Defines whether to translate Word content in exclude highlight mode. Default: true.
		TranslateWordInExcludeHighlightMode *bool `json:"translateWordInExcludeHighlightMode,omitempty"`
		// Defines whether to translate Word excluded colors. Default: false.
		TranslateWordExcludeColors *bool `json:"translateWordExcludeColors,omitempty"`
		// Word excluded colors. Default: [].
		WordExcludedColors []string `json:"wordExcludedColors,omitempty"`
		// Defines whether to translate copied Excel cells. Default: true.
		TranslateExcelCellsCopied *bool `json:"translateExcelCellsCopied,omitempty"`
		// Defines whether to translate Excel sheet names. Default: false.
		TranslateExcelSheetNames *bool `json:"translateExcelSheetNames,omitempty"`
		// Excel excluded colors. Default: [].
		ExcelExcludedColors []string `json:"excelExcludedColors,omitempty"`
		// Defines whether to translate Excel diagram data. Default: false.
		TranslateExcelDiagramData *bool `json:"translateExcelDiagramData,omitempty"`
		// Defines whether to translate Excel drawings. Default: false.
		TranslateExcelDrawings *bool `json:"translateExcelDrawings,omitempty"`
		// Defines whether to allow Word style optimization. Default: true.
		AllowWordStyleOptimization *bool `json:"allowWordStyleOptimization,omitempty"`
		// Defines whether to translate Excel excluded colors. Default: false.
		TranslateExcelExcludeColors *bool `json:"translateExcelExcludeColors,omitempty"`
	}

	MediaWikiFileFormatSettings struct {
		// Storage identifier of the SRX segmentation rules file. Default: null.
		SRXStorageID *int `json:"srxStorageId,omitempty"`
		// File format export pattern. Defines file name and path in resulting
		// translations bundle.
		// Default: null. Note: Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	JSONFileFormatSettings struct {
		CommonFileFormatSettings
		// Enum: "i18next_json", "nestjs_i18n".
		Type string `json:"type,omitempty"`
	}

	TXTFileFormatSettings struct {
		// Storage identifier of the SRX segmentation rules file. Default: null.
		SRXStorageID *int `json:"srxStorageId,omitempty"`
		// File format export pattern. Defines file name and path in resulting
		// translations bundle.
		// Default: null. Note: Can't contain : * ? " < > | symbols
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	JavaScriptFileFormatSettings struct {
		// File export pattern. Defines file name and path in resulting translations bundle.
		// Note: Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
		// Enum: "single" "double". Default: "single".
		// single - Output will be enclosed in single quotes.
		// double - Output will be enclosed in double quotes.
		ExportQuotes *string `json:"exportQuotes,omitempty"`
	}

	StringCatalogFileFormatSettings struct {
		// Determines whether to import the key as source string if it does not exist.
		// Default: true.
		ImportKeyAsSource *bool `json:"importKeyAsSource,omitempty"`
		// Determines whether to import existing translations from the file. Default: false.
		ImportTranslations *bool `json:"importTranslations,omitempty"`
		// File format export pattern. Defines file name and path in resulting translations bundle.
		// Default: null. Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	OtherFileFormatSettings struct {
		// File format export pattern. Defines file name and path in resulting translations bundle.
		// Default: null. Can't contain : * ? " < > | symbols.
		ExportPattern *string `json:"exportPattern,omitempty"`
	}

	AndroidFileFormatSettings     struct{ CommonFileFormatSettings }
	FMMDFileFormatSettings        struct{ CommonFileFormatSettings }
	MadCapFLSNPFileFormatSettings struct{ CommonFileFormatSettings }
	MIFFileFormatSettings         struct{ CommonFileFormatSettings }
	DitaFileFormatSettings        struct{ CommonFileFormatSettings }
	ARBFileFormatSettings         struct{ CommonFileFormatSettings }
	FJSFileFormatSettings         struct{ CommonFileFormatSettings }
	MacOSFileFormatSettings       struct{ CommonFileFormatSettings }
	ChromeFileFormatSettings      struct{ CommonFileFormatSettings }
	CSVFileFormatSettings         struct{ CommonFileFormatSettings }
	XLSXFileFormatSettings        struct{ CommonFileFormatSettings }
	ReactIntlFileFormatSettings   struct{ CommonFileFormatSettings }
	XLIFFFileFormatSettings       struct{ CommonFileFormatSettings }
)

// Validate checks if the add project file format settings request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ProjectsAddFileFormatSettingsRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Format == "" {
		return errors.New("format is required")
	}
	if r.Settings == nil {
		return errors.New("settings is required")
	}
	return r.Settings.ValidateSettings()
}

func (p *CommonFileFormatSettings) ValidateSettings() error        { return nil }
func (p *PropertyFileFormatSettings) ValidateSettings() error      { return nil }
func (p *XMLFileFormatSettings) ValidateSettings() error           { return nil }
func (p *MediaWikiFileFormatSettings) ValidateSettings() error     { return nil }
func (p *TXTFileFormatSettings) ValidateSettings() error           { return nil }
func (p *JavaScriptFileFormatSettings) ValidateSettings() error    { return nil }
func (p *StringCatalogFileFormatSettings) ValidateSettings() error { return nil }
func (p *OtherFileFormatSettings) ValidateSettings() error         { return nil }
func (p *VDFFileFormatSettings) ValidateSettings() error           { return nil }

// ProjectsStringsExporterSettings represents a Crowdin project strings
// exporter settings.
type ProjectsStringsExporterSettings struct {
	ID        int                     `json:"id"`
	Format    string                  `json:"format"`
	Settings  StringsExporterSettings `json:"settings"`
	CreatedAt string                  `json:"createdAt"`
	UpdatedAt string                  `json:"updatedAt"`
}

// ProjectsStringsExporterSettingsResponse defines the structure of a response when
// retrieving a project strings exporter settings.
type ProjectsStringsExporterSettingsResponse struct {
	Data *ProjectsStringsExporterSettings `json:"data"`
}

// ProjectsStringsExporterSettingsListResponse defines the structure of a response when
// getting a list of project strings exporter settings.
type ProjectsStringsExporterSettingsListResponse struct {
	Data []*ProjectsStringsExporterSettingsResponse `json:"data"`
}

// ProjectsStringsExporterSettingsRequest defines the structure of a request
// to update a project strings exporter settings.
type ProjectsStringsExporterSettingsRequest struct {
	// Defines strings exporter format. Enum: "android", "macosx", "xliff".
	Format string `json:"format"`
	// Defines strings exporter settings.
	Settings StringsExporterSettings `json:"settings"`
}

// StringsExporterSettings defines the structure of a strings exporter settings.
type StringsExporterSettings struct {
	// Convert placeholders to MacOSX format. Default: false.
	// Note: Only for Android and MacOSX formats.
	ConvertPlaceholders *bool `json:"convertPlaceholders,omitempty"`
	// Convert line breaks. Default: false.
	// Note: Only for Android and MacOSX formats.
	ConvertLineBreaks *bool `json:"convertLineBreaks,omitempty"`
	// Use CDATA for strings with tags. Default: false.
	// Note: Only for Android format.
	UseCdataForStringsWithTags *bool `json:"useCdataForStringsWithTags,omitempty"`
	// Export context. Default: true.
	// Note: Only for MacOSX format.
	ExportContext *bool `json:"exportContext,omitempty"`
	// Defines language pair mapping the target language for the specified source language.
	// Note: Only for XLIFF format.
	LanguagePairMapping map[string]string `json:"languagePairMapping,omitempty"`
	// Copy source to empty target. Default: true.
	// Note: Only for XLIFF format.
	CopySourceToEmptyTarget *bool `json:"copySourceToEmptyTarget,omitempty"`
	// Export translators comment. Default: true.
	// Note: Only for XLIFF format.
	ExportTranslatorsComment *bool `json:"exportTranslatorsComment,omitempty"`
}

// Validate checks if the update request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ProjectsStringsExporterSettingsRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Format == "" {
		return errors.New("format is required")
	}
	st := r.Settings
	if st.ConvertPlaceholders == nil && st.ConvertLineBreaks == nil && st.UseCdataForStringsWithTags == nil &&
		st.ExportContext == nil && len(st.LanguagePairMapping) == 0 && st.CopySourceToEmptyTarget == nil &&
		st.ExportTranslatorsComment == nil {
		return errors.New("settings is required")
	}
	return nil
}
