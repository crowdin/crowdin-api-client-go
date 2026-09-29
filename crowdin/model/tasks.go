package model

import (
	"errors"
	"fmt"
	"net/url"
)

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusDone       TaskStatus = "done"
	TaskStatusClosed     TaskStatus = "closed"
)

type (
	// Task represents a task in Crowdin.
	Task struct {
		ID               int                 `json:"id"`
		ProjectID        int                 `json:"projectId"`
		CreatorID        int                 `json:"creatorId"`
		Type             int                 `json:"type"`
		Status           TaskStatus          `json:"status"`
		Title            string              `json:"title"`
		BatchID          *int                `json:"batchId,omitempty"`
		Assignees        []*TaskAssignee     `json:"assignees"`
		AssignedTeams    []*TaskAssignedTeam `json:"assignedTeams"`
		Progress         TaskProgress        `json:"progress"`
		SourceLanguageID string              `json:"sourceLanguageId"`
		TargetLanguageID string              `json:"targetLanguageId"`
		Description      string              `json:"description"`
		TranslationURL   string              `json:"translationUrl"`
		WebURL           string              `json:"webUrl"`
		// Number of words currently in the task.
		WordsCount int `json:"wordsCount"`
		// Number of words when the task was created.
		// Preserved when source strings are later removed.
		OriginalWordsCount           int         `json:"originalWordsCount"`
		CommentsCount                int         `json:"commentsCount"`
		Deadline                     string      `json:"deadline"`
		StartedAt                    string      `json:"startedAt"`
		ResolvedAt                   string      `json:"resolvedAt"`
		TimeRange                    string      `json:"timeRange"`
		TranslationsUpdatedTimeRange string      `json:"translationsUpdatedTimeRange,omitempty"`
		WorkflowStepID               int         `json:"workflowStepId"`
		BuyURL                       string      `json:"buyUrl"`
		CreatedAt                    string      `json:"createdAt"`
		UpdatedAt                    string      `json:"updatedAt"`
		SourceLanguage               *Language   `json:"sourceLanguage"`
		TargetLanguages              []*Language `json:"targetLanguages"`
		LabelIDs                     []int       `json:"labelIds"`
		// Match rule for labels. Enum: all, any.
		LabelMatchRule  *string `json:"labelMatchRule,omitempty"`
		ExcludeLabelIDs []int   `json:"excludeLabelIds"`
		// Match rule for excluded labels. Enum: all, any.
		ExcludeLabelMatchRule    *string   `json:"excludeLabelMatchRule,omitempty"`
		PrecedingTaskID          int       `json:"precedingTaskId"`
		EstimatedCost            *TaskCost `json:"estimatedCost,omitempty"`
		ActualCost               *TaskCost `json:"actualCost,omitempty"`
		GenerateCostEstimate     *bool     `json:"generateCostEstimate,omitempty"`
		GenerateTranslationCost  *bool     `json:"generateTranslationCost,omitempty"`
		ReportSettingsTemplateID *int      `json:"reportSettingsTemplateId,omitempty"`
		FilesCount               int       `json:"filesCount"`
		FileIDs                  []int     `json:"fileIds,omitempty"`
		Vendor                   string    `json:"vendor,omitempty"`
		BranchIDs                []int     `json:"branchIds,omitempty"`
		IsArchived               *bool     `json:"isArchived,omitempty"`
		Fields                   any       `json:"fields,omitempty"`
		// How a shared task's words split while it syncs to the vendor.
		// It is nil when the task is fully synced, not shared, or the
		// counterpart is unavailable. Crowdin Enterprise only.
		SyncScope *TaskSyncScope `json:"syncScope,omitempty"`
	}

	// TaskAssignee represents an assignee of a task.
	TaskAssignee struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		FullName   string `json:"fullName"`
		AvatarURL  string `json:"avatarUrl"`
		WordsCount int    `json:"wordsCount"`
		WordsLeft  int    `json:"wordsLeft"`
		TimeSpent  int    `json:"timeSpent,omitempty"`
	}

	// TaskAssignedTeam represents a team assigned to a task.
	TaskAssignedTeam struct {
		ID         int `json:"id"`
		WordsCount int `json:"wordsCount"`
		TimeSpent  int `json:"timeSpent,omitempty"`
	}

	// TaskProgress represents the progress of a task.
	TaskProgress struct {
		Total   int `json:"total"`
		Done    int `json:"done"`
		Percent int `json:"percent"`
	}

	// TaskCost represents the estimated or actual cost of a task.
	TaskCost struct {
		Cost     float64 `json:"cost"`
		Date     string  `json:"date"`
		Currency string  `json:"currency"`
	}

	// TaskSyncScope represents how a shared task's words split
	// while it syncs to the vendor.
	TaskSyncScope struct {
		// Words already synced into the vendor task.
		SyncedWords int `json:"syncedWords"`
		// Words routed to the vendor and awaiting sync.
		PendingWords int `json:"pendingWords"`
		// Words the client's workflow does not share with the vendor.
		SkippedWords int `json:"skippedWords"`
	}
)

// TaskResponse defines the structure of the response
// when getting a task.
type TaskResponse struct {
	Data *Task `json:"data"`
}

// TasksListResponse defines the structure of the response
// when getting a list of tasks.
type TasksListResponse struct {
	Data []*TaskResponse `json:"data"`
}

// TasksListOptions specifies the optional parameters to the
// TasksService.List method.
type TasksListOptions struct {
	// Sort a list of tasks by a specified field.
	// Enum: id, type, title, status, description, createdAt,
	// updatedAt, deadline, startedAt, resolvedAt. Default: id.
	// Example: orderBy=createdAt desc,title
	OrderBy string `json:"orderBy,omitempty"`
	// List tasks with specified statuses. It can be one status
	// or a list of status values.
	// Enum: todo, in_progress, done, closed.
	Status []TaskStatus `json:"status,omitempty"`
	// List tasks for specified assignee.
	AssigneeID int `json:"assigneeId,omitempty"`
	// Filter by task batch.
	BatchID int `json:"batchId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of TasksListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *TasksListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if len(o.Status) > 0 {
		v.Add("status", JoinSlice(o.Status))
	}
	if o.AssigneeID > 0 {
		v.Add("assigneeId", fmt.Sprintf("%d", o.AssigneeID))
	}
	if o.BatchID > 0 {
		v.Add("batchId", fmt.Sprintf("%d", o.BatchID))
	}

	return v, len(v) > 0
}

// AllTasksListOptions specifies the optional parameters to the
// TasksService.ListAll method.
type AllTasksListOptions struct {
	// Sort a list of tasks by a specified field.
	// Example: orderBy=createdAt desc,title.
	OrderBy string `json:"orderBy,omitempty"`
	// List tasks with specified statuses. It can be one status
	// or a list of status values.
	// Enum: todo, in_progress, done, closed.
	Status []TaskStatus `json:"status,omitempty"`
	// Filter by task type. It can be one type or a list of type values.
	// Enum: 0 - translate, 1 - proofread, 2 - translate by vendor,
	// 3 - proofread by vendor.
	Type []TaskType `json:"type,omitempty"`
	// Filter by project IDs.
	// Cannot be used together with GroupIDs.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// Filter by group IDs.
	// Cannot be used together with ProjectIDs. Enterprise only.
	GroupIDs []int `json:"groupIds,omitempty"`
	// Filter by assignee user IDs.
	AssigneeIDs []int `json:"assigneeIds,omitempty"`
	// Filter by creator user IDs.
	CreatorIDs []int `json:"creatorIds,omitempty"`
	// Filter by target language IDs.
	TargetLanguageIDs []string `json:"targetLanguageIds,omitempty"`
	// Filter by source language IDs.
	SourceLanguageIDs []string `json:"sourceLanguageIds,omitempty"`
	// Filter tasks created from this date (ISO 8601).
	CreatedAtFrom string `json:"createdAtFrom,omitempty"`
	// Filter tasks created until this date (ISO 8601).
	CreatedAtTo string `json:"createdAtTo,omitempty"`
	// Filter tasks with deadline from this date (ISO 8601).
	DeadlineFrom string `json:"deadlineFrom,omitempty"`
	// Filter tasks with deadline until this date (ISO 8601).
	DeadlineTo string `json:"deadlineTo,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of AllTasksListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *AllTasksListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if len(o.Status) > 0 {
		v.Add("status", JoinSlice(o.Status))
	}
	if len(o.Type) > 0 {
		v.Add("type", JoinSlice(o.Type))
	}
	if len(o.ProjectIDs) > 0 {
		v.Add("projectIds", JoinSlice(o.ProjectIDs))
	}
	if len(o.GroupIDs) > 0 {
		v.Add("groupIds", JoinSlice(o.GroupIDs))
	}
	if len(o.AssigneeIDs) > 0 {
		v.Add("assigneeIds", JoinSlice(o.AssigneeIDs))
	}
	if len(o.CreatorIDs) > 0 {
		v.Add("creatorIds", JoinSlice(o.CreatorIDs))
	}
	if len(o.TargetLanguageIDs) > 0 {
		v.Add("targetLanguageIds", JoinSlice(o.TargetLanguageIDs))
	}
	if len(o.SourceLanguageIDs) > 0 {
		v.Add("sourceLanguageIds", JoinSlice(o.SourceLanguageIDs))
	}
	if o.CreatedAtFrom != "" {
		v.Add("createdAtFrom", o.CreatedAtFrom)
	}
	if o.CreatedAtTo != "" {
		v.Add("createdAtTo", o.CreatedAtTo)
	}
	if o.DeadlineFrom != "" {
		v.Add("deadlineFrom", o.DeadlineFrom)
	}
	if o.DeadlineTo != "" {
		v.Add("deadlineTo", o.DeadlineTo)
	}

	return v, len(v) > 0
}

// TaskType represents the type of a task.
type TaskType int

const (
	TaskTypeTranslate         TaskType = 0
	TaskTypeProofread         TaskType = 1
	TaskTypeTranslateByVendor TaskType = 2
	TaskTypeProofreadByVendor TaskType = 3
)

// TaskVendor represents the vendor of a task.
type TaskVendor string

const (
	TaskVendorCrowdinLanguageService   TaskVendor = "crowdin_language_service"
	TaskVendorOht                      TaskVendor = "oht"
	TaskVendorGengo                    TaskVendor = "gengo"
	TaskVendorUndertow                 TaskVendor = "undertow"
	TaskVendorManual                   TaskVendor = "manual"
	TaskVendorAlconost                 TaskVendor = "alconost"
	TaskVendorBabbleon                 TaskVendor = "babbleon"
	TaskVendorTomedes                  TaskVendor = "tomedes"
	TaskVendorE2f                      TaskVendor = "e2f"
	TaskVendorWritePathAdmin           TaskVendor = "write_path_admin"
	TaskVendorInlingo                  TaskVendor = "inlingo"
	TaskVendorAcclaro                  TaskVendor = "acclaro"
	TaskVendorTranslateByHumans        TaskVendor = "translate_by_humans"
	TaskVendorLingo24                  TaskVendor = "lingo24"
	TaskVendorAssertioLanguageServices TaskVendor = "assertio_language_services"
	TaskVendorGteLocalize              TaskVendor = "gte_localize"
	TaskVendorKettuSolutions           TaskVendor = "kettu_solutions"
	TaskVendorLanguageLineSolutions    TaskVendor = "languageline_solutions"
)

// TaskAddRequester is an interface encapsulating a request for task addition.
// The request body must conform to one of the following struct types:
//
//	TaskCreateForm
//	VendorTaskCreateForm
//	PendingTaskCreateForm
//
// The following forms are deprecated and kept for backward compatibility
// (use VendorTaskCreateForm or PendingTaskCreateForm instead):
//
//	LanguageServiceTaskCreateForm
//	VendorOhtTaskCreateForm
//	VendorGengoTaskCreateForm
//	VendorManualTaskCreateForm
//	LanguageServicePendingTaskCreateForm
//	VendorManualPendingTaskCreateForm
//
// For the Enterprise API, the request body should be one of the following structs:
//
//	EnterpriseTaskCreateForm
//	EnterpriseVendorTaskCreateForm
//	EnterprisePendingTaskCreateForm
type TaskAddRequester interface {
	ValidateRequest() error
}

type CrowdinTaskAssignee struct {
	// Project member identifier.
	ID int `json:"id"`
	// Defines how many words (starting from 1) are assigned
	// to each task assignee. Note: Can be used only when
	// `splitContent` parameter is specified.
	WordsCount int `json:"wordsCount,omitempty"`
}

type (
	TaskCreateForm struct {
		// Task title
		Title string `json:"title"`
		// Task language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 0 - translate, 1 - proofread.
		Type *TaskType `json:"type"`
		// Branch identifiers.
		// One of branchIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// Task string identifiers.
		// One of branchIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// Task file identifiers.
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Directory identifiers (file-based projects only).
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Match rule for labels. Enum: all, any.
		// Note: Can only be used when `labelIds` parameter is provided.
		LabelMatchRule string `json:"labelMatchRule,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Match rule for excluded labels. Enum: all, any.
		// Note: Can only be used when `excludeLabelIds` parameter is provided.
		ExcludeLabelMatchRule string `json:"excludeLabelMatchRule,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Split content for task.
		SplitContent *bool `json:"splitContent,omitempty"`
		// Skip strings already included in other tasks. Default: false.
		SkipAssignedStrings *bool `json:"skipAssignedStrings,omitempty"`
		// Defines whether to export only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false`,
		// `type=0` or `type=2` in same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
		// Task started date. Format: UTC, ISO 8601.
		StartedAt string `json:"startedAt,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Start date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateFrom string `json:"translationsUpdatedDateFrom,omitempty"`
		// End date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateTo string `json:"translationsUpdatedDateTo,omitempty"`
		// Generate cost estimate report for the task. Requires `reportSettingsTemplateId`.
		GenerateCostEstimate *bool `json:"generateCostEstimate,omitempty"`
		// Generate translation cost report for the task. Requires `reportSettingsTemplateId`.
		GenerateTranslationCost *bool `json:"generateTranslationCost,omitempty"`
		// Report settings template identifier. Required when `generateCostEstimate`
		// or `generateTranslationCost` is `true`.
		ReportSettingsTemplateID int `json:"reportSettingsTemplateId,omitempty"`
		// Task batch identifier. Allows grouping tasks into a batch.
		BatchID int `json:"batchId,omitempty"`
	}

	// VendorTaskCreateForm defines the structure of the request to create
	// a task assigned to a vendor (Crowdin only).
	VendorTaskCreateForm struct {
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 2 - translate by vendor, 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Vendor identifier. Built-in vendors: crowdin_language_service, oht,
		// gengo, undertow. For any other vendor, use its identifier from the
		// Crowdin Store.
		Vendor TaskVendor `json:"vendor"`
		// Branch identifiers.
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// String identifiers.
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers (file-based projects only).
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Directory identifiers (file-based projects only).
		// One of branchIds, directoryIds, stringIds or fileIds is required.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Match rule for labels. Enum: all, any.
		LabelMatchRule string `json:"labelMatchRule,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Match rule for excluded labels. Enum: all, any.
		ExcludeLabelMatchRule string `json:"excludeLabelMatchRule,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Defines whether to include only pretranslated strings.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Skip strings already included in other tasks.
		// Note: Not supported by all vendors.
		SkipAssignedStrings *bool `json:"skipAssignedStrings,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		// Note: Not supported by all vendors.
		Deadline string `json:"deadline,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Start date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateFrom string `json:"translationsUpdatedDateFrom,omitempty"`
		// End date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateTo string `json:"translationsUpdatedDateTo,omitempty"`
		// Generate cost estimate report for the task. Requires `reportSettingsTemplateId`.
		GenerateCostEstimate *bool `json:"generateCostEstimate,omitempty"`
		// Generate translation cost report for the task. Requires `reportSettingsTemplateId`.
		GenerateTranslationCost *bool `json:"generateTranslationCost,omitempty"`
		// Report settings template identifier. Required when `generateCostEstimate`
		// or `generateTranslationCost` is `true`.
		ReportSettingsTemplateID int `json:"reportSettingsTemplateId,omitempty"`
	}

	// LanguageServiceTaskCreateForm defines the structure of the request to create
	// a Crowdin Language Service task.
	//
	// Deprecated: use VendorTaskCreateForm instead.
	LanguageServiceTaskCreateForm struct {
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 2 - translate by vendor, 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: "crowdin_language_service".
		Vendor TaskVendor `json:"vendor"`
		// Branch identifiers.
		// One of branchIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// String identifiers.
		// One of branchIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers.
		// One of branchIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Defines whether to include only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false` or
		// `includeUntranslatedStringsOnly=true` in the same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// VendorOhtTaskCreateForm defines the structure of the request to create
	// a OneHourTranslation vendor task.
	//
	// Deprecated: use VendorTaskCreateForm instead.
	VendorOhtTaskCreateForm struct {
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 2 - translate by vendor, 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: "oht" - OneHourTranslation.
		Vendor TaskVendor `json:"vendor"`
		// Branch identifiers.
		// One of branchIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// String identifiers.
		// One of branchIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers.
		// One of branchIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Task expertise. Default: standard.
		// Enum: standard, mobile-applications, software-it, gaming-video-games,
		// technical-engineering, marketing-consumer-media, business-finance,
		// legal-certificate, medical, ad-words-banners, automotive-aerospace,
		// scientific, scientific-academic, tourism, training-employee-handbooks,
		// forex-crypto.
		Expertise string `json:"expertise,omitempty"`
		// Enables Edit stage for all jobs. Default: false.
		EditService *bool `json:"editService,omitempty"`
		// Defines whether to include only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false` or
		// `includeUntranslatedStringsOnly=true` in the same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// VendorGengoTaskCreateForm defines the structure of the request to create
	// a Gengo vendor task.
	//
	// Deprecated: use VendorTaskCreateForm instead.
	VendorGengoTaskCreateForm struct {
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 2 - translate by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: "gengo" - Gengo.
		Vendor TaskVendor `json:"vendor"`
		// Branch identifiers.
		// One of branchIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// String identifiers.
		// One of branchIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers.
		// One of branchIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Task expertise. Default: standard.
		// Enum: standard, pro.
		Expertise string `json:"expertise,omitempty"`
		// Task tone. Default: "".
		// Enum: "", "Informal", "Friendly", "Business", "Formal", "other".
		Tone *string `json:"tone,omitempty"`
		// Task purpose. Default: "standard".
		// Enum: "Personal use", "Business", "Online content", "App/Web localization",
		// "Media content", "Semi-technical", "other".
		Purpose string `json:"purpose,omitempty"`
		// Instructions for translators.
		CustomerMessage string `json:"customerMessage,omitempty"`
		// Use preferred translators. Default: false.
		UsePreferred *bool `json:"usePreferred,omitempty"`
		// Enables Edit stage for all jobs. Default: false.
		EditService *bool `json:"editService,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	// VendorManualTaskCreateForm defines the structure of the request to create
	// a manual vendor task.
	//
	// Deprecated: use VendorTaskCreateForm instead.
	VendorManualTaskCreateForm struct {
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// Task type. Enum: 2 - translate by vendor, 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: alconost, babbleon, tomedes, e2f, write_path_admin,
		// inlingo, acclaro, translate_by_humans, lingo24, assertio_language_services,
		// gte_localize, kettu_solutions, languageline_solutions.
		Vendor TaskVendor `json:"vendor"`
		// Branch identifiers.
		// One of branchIds, stringIds or fileIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// String identifiers.
		// One of branchIds, stringIds or fileIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers.
		// One of branchIds, stringIds or fileIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Skip strings already included in other tasks. Default: false.
		SkipAssignedStrings *bool `json:"skipAssignedStrings,omitempty"`
		// Defines whether to export only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false`,
		// `type=0` or `type=2` in same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
		// Task started date. Format: UTC, ISO 8601.
		StartedAt string `json:"startedAt,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
	}

	PendingTaskCreateForm struct {
		// Translate task identifier.
		PrecedingTaskID int `json:"precedingTaskId"`
		// Task type. Enum: 1 - proofread, 3 - proofread by vendor
		// (requires `vendor` field).
		Type TaskType `json:"type"`
		// Task title.
		Title string `json:"title"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Vendor identifier. Required when `type` is 3.
		Vendor TaskVendor `json:"vendor,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
	}

	// LanguageServicePendingTaskCreateForm defines the structure of the request
	// to create a pending Crowdin Language Service task.
	//
	// Deprecated: use PendingTaskCreateForm with the Vendor field instead.
	LanguageServicePendingTaskCreateForm struct {
		// Translate task identifier.
		PrecedingTaskID int `json:"precedingTaskId"`
		// Task type. Enum: 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: "crowdin_language_service".
		Vendor TaskVendor `json:"vendor"`
		// Task title.
		Title string `json:"title"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
	}

	// VendorManualPendingTaskCreateForm defines the structure of the request
	// to create a pending manual vendor task.
	//
	// Deprecated: use PendingTaskCreateForm with the Vendor field instead.
	VendorManualPendingTaskCreateForm struct {
		// Translate task identifier.
		PrecedingTaskID int `json:"precedingTaskId"`
		// Task type. Enum: 3 - proofread by vendor.
		Type TaskType `json:"type"`
		// Task vendor. Enum: alconost, babbleon, tomedes, e2f, write_path_admin,
		// inlingo, acclaro, translate_by_humans, lingo24, assertio_language_services,
		// gte_localize, kettu_solutions, languageline_solutions.
		Vendor TaskVendor `json:"vendor"`
		// Task title.
		Title string `json:"title"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
	}
)

type (
	EnterpriseTaskCreateForm struct {
		// Task type. Enum: 0 - translate, 1 - proofread.
		// Note: Can't be used with `workflowStepId` in same request.
		Type *TaskType `json:"type"`
		// Task workflow step id.
		// Note: Can't be used with `type` in same request.
		WorkflowStepID int `json:"workflowStepId,omitempty"`
		// Task title
		Title string `json:"title"`
		// Task language identifier.
		LanguageID string `json:"languageId"`
		// Task string identifiers.
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// Task file identifiers (file-based projects only).
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Directory identifiers (file-based projects only).
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// Branch identifiers.
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Match rule for labels. Enum: all, any.
		LabelMatchRule string `json:"labelMatchRule,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Match rule for excluded labels. Enum: all, any.
		ExcludeLabelMatchRule string `json:"excludeLabelMatchRule,omitempty"`
		// Task status. Enum: todo, in_progress.
		Status TaskStatus `json:"status,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Split content for task.
		SplitContent *bool `json:"splitContent,omitempty"`
		// Skip strings already included in other tasks. Default: false.
		SkipAssignedStrings *bool `json:"skipAssignedStrings,omitempty"`
		// Defines which other tasks the `skipAssignedStrings` filter considers.
		// Enum: all, sameWorkflowStep.
		// Note: Can only be used when `skipAssignedStrings` is `true`.
		SkipAssignedStringsScope string `json:"skipAssignedStringsScope,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task assigned teams.
		AssignedTeams []TaskAssignedTeam `json:"assignedTeams,omitempty"`
		// Defines whether to export only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false`,
		// `type=0` or `type=2` in same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
		// Task started date. Format: UTC, ISO 8601.
		StartedAt string `json:"startedAt,omitempty"`
		// Start date for interval when strings were modified. Format: UTC, ISO 8601.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Start date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateFrom string `json:"translationsUpdatedDateFrom,omitempty"`
		// End date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateTo string `json:"translationsUpdatedDateTo,omitempty"`
		// Generate cost estimate report for the task. Requires `reportSettingsTemplateId`.
		GenerateCostEstimate *bool `json:"generateCostEstimate,omitempty"`
		// Generate translation cost report for the task. Requires `reportSettingsTemplateId`.
		GenerateTranslationCost *bool `json:"generateTranslationCost,omitempty"`
		// Report settings template identifier. Required when `generateCostEstimate`
		// or `generateTranslationCost` is `true`.
		ReportSettingsTemplateID int `json:"reportSettingsTemplateId,omitempty"`
		// Fields for task.
		Fields map[string]any `json:"fields,omitempty"`
		// Task batch identifier. Allows grouping tasks into a batch.
		BatchID int `json:"batchId,omitempty"`
	}

	EnterpriseVendorTaskCreateForm struct {
		// Task workflow step id with type `Translate by Vendor` or `Proofread by Vendor`.
		WorkflowStepID int `json:"workflowStepId"`
		// Task title.
		Title string `json:"title"`
		// Language identifier.
		LanguageID string `json:"languageId"`
		// String identifiers.
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		StringIDs []int `json:"stringIds,omitempty"`
		// File identifiers (file-based projects only).
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		FileIDs []int `json:"fileIds,omitempty"`
		// Directory identifiers (file-based projects only).
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		DirectoryIDs []int `json:"directoryIds,omitempty"`
		// Branch identifiers.
		// One of stringIds, fileIds, directoryIds or branchIds is required.
		BranchIDs []int `json:"branchIds,omitempty"`
		// Label identifiers.
		LabelIDs []int `json:"labelIds,omitempty"`
		// Match rule for labels. Enum: all, any.
		LabelMatchRule string `json:"labelMatchRule,omitempty"`
		// Exclude label identifiers.
		ExcludeLabelIDs []int `json:"excludeLabelIds,omitempty"`
		// Match rule for excluded labels. Enum: all, any.
		ExcludeLabelMatchRule string `json:"excludeLabelMatchRule,omitempty"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Skip strings already included in other tasks. Default: false.
		SkipAssignedStrings *bool `json:"skipAssignedStrings,omitempty"`
		// Defines which other tasks the `skipAssignedStrings` filter considers.
		// Enum: all, sameWorkflowStep.
		// Note: Can only be used when `skipAssignedStrings` is `true`.
		SkipAssignedStringsScope string `json:"skipAssignedStringsScope,omitempty"`
		// Defines whether to export only pretranslated strings. Default: false.
		// Note: `true` value can't be used with `skipUntranslatedStrings=false`,
		// `type=0` or `type=2` in same request.
		IncludePreTranslatedStringsOnly *bool `json:"includePreTranslatedStringsOnly,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
		// Task started date. Format: UTC, ISO 8601.
		StartedAt string `json:"startedAt,omitempty"`
		// End date for interval when strings were modified. Format: UTC, ISO 8601.
		DateTo string `json:"dateTo,omitempty"`
		// Start date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateFrom string `json:"translationsUpdatedDateFrom,omitempty"`
		// End date for interval when translations were updated (proofread tasks only).
		// Format: UTC, ISO 8601.
		TranslationsUpdatedDateTo string `json:"translationsUpdatedDateTo,omitempty"`
		// Generate cost estimate report for the task. Requires `reportSettingsTemplateId`.
		GenerateCostEstimate *bool `json:"generateCostEstimate,omitempty"`
		// Generate translation cost report for the task. Requires `reportSettingsTemplateId`.
		GenerateTranslationCost *bool `json:"generateTranslationCost,omitempty"`
		// Report settings template identifier. Required when `generateCostEstimate`
		// or `generateTranslationCost` is `true`.
		ReportSettingsTemplateID int `json:"reportSettingsTemplateId,omitempty"`
		// Fields for task.
		Fields map[string]any `json:"fields,omitempty"`
	}

	EnterprisePendingTaskCreateForm struct {
		// Translate task identifier.
		PrecedingTaskID int `json:"precedingTaskId"`
		// Task type. Enum: 1 - proofread, 3 - proofread by vendor
		// (requires `vendor` field).
		// Note: One of `type` or `workflowStepId` is required.
		// Both cannot be provided simultaneously.
		Type TaskType `json:"type,omitempty"`
		// Task workflow step identifier.
		// Note: One of `type` or `workflowStepId` is required.
		// Both cannot be provided simultaneously.
		WorkflowStepID int `json:"workflowStepId,omitempty"`
		// Task title.
		Title string `json:"title"`
		// Task description.
		Description string `json:"description,omitempty"`
		// Vendor identifier. Required when `type` is 3.
		Vendor TaskVendor `json:"vendor,omitempty"`
		// Task assignees.
		Assignees []CrowdinTaskAssignee `json:"assignees,omitempty"`
		// Task assigned teams.
		AssignedTeams []TaskAssignedTeam `json:"assignedTeams,omitempty"`
		// Task deadline date. Format: UTC, ISO 8601.
		Deadline string `json:"deadline,omitempty"`
	}
)

// Validate checks if the request is valid.
func (r *TaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *TaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type == nil || (*r.Type != TaskTypeTranslate && *r.Type != TaskTypeProofread) {
		return fmt.Errorf("type is required and must be one of %d, %d", TaskTypeTranslate, TaskTypeProofread)
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 && len(r.DirectoryIDs) == 0 {
		return errors.New("one of stringIds, fileIds, directoryIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *LanguageServiceTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *LanguageServiceTaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type != TaskTypeTranslateByVendor && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be one of %d, %d",
			TaskTypeTranslateByVendor, TaskTypeProofreadByVendor)
	}
	if r.Vendor != TaskVendorCrowdinLanguageService {
		return fmt.Errorf("vendor is required and must be %q", TaskVendorCrowdinLanguageService)
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *VendorOhtTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *VendorOhtTaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type != TaskTypeTranslateByVendor && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be one of %d, %d",
			TaskTypeTranslateByVendor, TaskTypeProofreadByVendor)
	}
	if r.Vendor != TaskVendorOht {
		return fmt.Errorf("vendor is required and must be %q", TaskVendorOht)
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *VendorGengoTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *VendorGengoTaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type != TaskTypeTranslateByVendor {
		return fmt.Errorf("type is required and must be %d", TaskTypeTranslateByVendor)
	}
	if r.Vendor != TaskVendorGengo {
		return fmt.Errorf("vendor is required and must be %q", TaskVendorGengo)
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *VendorManualTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *VendorManualTaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type != TaskTypeTranslateByVendor && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be one of %d, %d",
			TaskTypeTranslateByVendor, TaskTypeProofreadByVendor)
	}
	if r.Vendor == "" {
		return errors.New("vendor is required")
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *PendingTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *PendingTaskCreateForm) ValidateRequest() error {
	if r.PrecedingTaskID == 0 {
		return errors.New("precedingTaskId is required")
	}
	if r.Type != TaskTypeProofread && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be one of %d, %d", TaskTypeProofread, TaskTypeProofreadByVendor)
	}
	if r.Type == TaskTypeProofreadByVendor && r.Vendor == "" {
		return fmt.Errorf("vendor is required when type is %d", TaskTypeProofreadByVendor)
	}
	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *VendorTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *VendorTaskCreateForm) ValidateRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if r.Type != TaskTypeTranslateByVendor && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be one of %d, %d",
			TaskTypeTranslateByVendor, TaskTypeProofreadByVendor)
	}
	if r.Vendor == "" {
		return errors.New("vendor is required")
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.BranchIDs) == 0 && len(r.DirectoryIDs) == 0 {
		return errors.New("one of stringIds, fileIds, directoryIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *LanguageServicePendingTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *LanguageServicePendingTaskCreateForm) ValidateRequest() error {
	if r.PrecedingTaskID == 0 {
		return errors.New("precedingTaskId is required")
	}
	if r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be %d", TaskTypeProofreadByVendor)
	}
	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *VendorManualPendingTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *VendorManualPendingTaskCreateForm) ValidateRequest() error {
	if r.PrecedingTaskID == 0 {
		return errors.New("precedingTaskId is required")
	}
	if r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type is required and must be %d", TaskTypeProofreadByVendor)
	}
	if r.Vendor == "" {
		return errors.New("vendor is required")
	}
	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *EnterpriseTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *EnterpriseTaskCreateForm) ValidateRequest() error {
	if r.WorkflowStepID == 0 && r.Type == nil {
		return errors.New("workflowStepId or type is required")
	}
	if r.WorkflowStepID > 0 && r.Type != nil {
		return errors.New("workflowStepId and type can't be used in the same request")
	}
	if r.Type != nil && (*r.Type != TaskTypeTranslate && *r.Type != TaskTypeProofread) {
		return fmt.Errorf("type must be one of %d, %d", TaskTypeTranslate, TaskTypeProofread)
	}
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.DirectoryIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds, directoryIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *EnterpriseVendorTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *EnterpriseVendorTaskCreateForm) ValidateRequest() error {
	if r.WorkflowStepID == 0 {
		return errors.New("workflowStepId is required")
	}
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.LanguageID == "" {
		return errors.New("languageId is required")
	}
	if len(r.StringIDs) == 0 && len(r.FileIDs) == 0 && len(r.DirectoryIDs) == 0 && len(r.BranchIDs) == 0 {
		return errors.New("one of stringIds, fileIds, directoryIds or branchIds is required")
	}

	return nil
}

// Validate checks if the request is valid.
func (r *EnterprisePendingTaskCreateForm) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return r.ValidateRequest()
}

// ValidateRequest validates the request.
func (r *EnterprisePendingTaskCreateForm) ValidateRequest() error {
	if r.PrecedingTaskID == 0 {
		return errors.New("precedingTaskId is required")
	}
	if r.WorkflowStepID > 0 && r.Type != 0 {
		return errors.New("workflowStepId and type can't be used in the same request")
	}
	if r.WorkflowStepID == 0 && r.Type != TaskTypeProofread && r.Type != TaskTypeProofreadByVendor {
		return fmt.Errorf("type (one of %d, %d) or workflowStepId is required", TaskTypeProofread, TaskTypeProofreadByVendor)
	}
	if r.Type == TaskTypeProofreadByVendor && r.Vendor == "" {
		return fmt.Errorf("vendor is required when type is %d", TaskTypeProofreadByVendor)
	}
	if r.Title == "" {
		return errors.New("title is required")
	}

	return nil
}

// UserTasksListOptions specifies the optional parameters to the
// TasksService.ListUserTasks method.
type UserTasksListOptions struct {
	// Sort a list of user tasks by a specified field.
	// Enum: id, title, description, createdAt, updatedAt,
	// deadline, startedAt, resolvedAt. Default: id.
	// Example: orderBy=createdAt desc,title
	OrderBy string `json:"orderBy,omitempty"`
	// List tasks with specified statuses. It can be one status
	// or a list of status values.
	// Enum: todo, in_progress, done, closed.
	Status []TaskStatus `json:"status,omitempty"`
	// List archived/not archived tasks for the authorized user.
	// Enum: 1 - archived, 0 - not archived. Default: 0.
	IsArchived *int `json:"isArchived,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of UserTasksListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *UserTasksListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if len(o.Status) > 0 {
		v.Add("status", JoinSlice(o.Status))
	}
	if o.IsArchived != nil {
		v.Add("isArchived", fmt.Sprintf("%d", *o.IsArchived))
	}

	return v, len(v) > 0
}

type (
	// TaskSettingsTemplate represents a task settings template.
	TaskSettingsTemplate struct {
		ID        int                        `json:"id"`
		Name      string                     `json:"name"`
		Config    TaskSettingsTemplateConfig `json:"config"`
		CreatedAt string                     `json:"createdAt"`
		UpdatedAt string                     `json:"updatedAt"`
	}

	// TaskSettingsTemplateConfig represents the configuration of a task
	// settings template.
	TaskSettingsTemplateConfig struct {
		Languages []TaskSettingsTemplateLanguage `json:"languages"`
	}

	// TaskSettingsTemplateLanguage represents the language settings of a
	// task settings template.
	TaskSettingsTemplateLanguage struct {
		LanguageID string   `json:"languageId"`
		UserIDs    []UserID `json:"userIds"`
		TeamIDs    []int    `json:"teamIds,omitempty"`
	}
)

// TaskSettingsTemplateResponse defines the structure of the response
// when getting a task settings template.
type TaskSettingsTemplateResponse struct {
	Data *TaskSettingsTemplate `json:"data"`
}

// TaskSettingsTemplatesListResponse defines the structure of the response
// when getting a list of task settings templates.
type TaskSettingsTemplatesListResponse struct {
	Data []*TaskSettingsTemplateResponse `json:"data"`
}

// TaskSettingsTemplateAddRequest defines the structure of the request
// when adding a new task settings template.
type TaskSettingsTemplateAddRequest struct {
	// Template name
	Name string `json:"name"`
	// Defines task config
	Config TaskSettingsTemplateConfig `json:"config"`
}

// Validate checks if the request is valid.
// It implements the crowdin.Validator interface.
func (r *TaskSettingsTemplateAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	if r.Name == "" {
		return errors.New("name is required")
	}
	if len(r.Config.Languages) == 0 {
		return errors.New("config languages is required")
	}

	return nil
}

// TaskComment represents a comment on a task.
type TaskComment struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	TaskID    int    `json:"taskId"`
	Text      string `json:"text"`
	TimeSpent int    `json:"timeSpent"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// TaskCommentResponse defines the structure of the response
// when getting a task comment.
type TaskCommentResponse struct {
	Data *TaskComment `json:"data"`
}

// TaskCommentsListResponse defines the structure of the response
// when getting a list of task comments.
type TaskCommentsListResponse struct {
	Data []*TaskCommentResponse `json:"data"`
}

// TaskCommentAddRequest defines the structure of the request
// when adding a new comment to a task.
type TaskCommentAddRequest struct {
	// Comment text
	Text string `json:"text,omitempty"`
	// Specifies the time spent on the task in seconds
	TimeSpent int `json:"timeSpent,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.Validator interface.
func (r *TaskCommentAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}
