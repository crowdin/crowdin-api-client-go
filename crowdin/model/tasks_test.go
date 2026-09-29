package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTasksListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *TasksListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &TasksListOptions{},
		},
		{
			name: "all options",
			opts: &TasksListOptions{OrderBy: "createdAt desc,name", Status: []TaskStatus{TaskStatusTodo, TaskStatusDone},
				AssigneeID: 1, BatchID: 7, ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "assigneeId=1&batchId=7&limit=10&offset=5&orderBy=createdAt+desc%2Cname&status=todo%2Cdone",
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

func TestAllTasksListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *AllTasksListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &AllTasksListOptions{},
		},
		{
			name: "with type translate",
			opts: &AllTasksListOptions{Type: []TaskType{TaskTypeTranslate}},
			out:  "type=0",
		},
		{
			name: "all options",
			opts: &AllTasksListOptions{
				OrderBy:           "createdAt desc,title",
				Status:            []TaskStatus{TaskStatusTodo, TaskStatusInProgress},
				Type:              []TaskType{TaskTypeTranslate, TaskTypeProofreadByVendor},
				ProjectIDs:        []int{1, 2},
				GroupIDs:          []int{3},
				AssigneeIDs:       []int{4, 5},
				CreatorIDs:        []int{6},
				TargetLanguageIDs: []string{"uk", "de"},
				SourceLanguageIDs: []string{"en"},
				CreatedAtFrom:     "2024-01-01T00:00:00+00:00",
				CreatedAtTo:       "2024-02-01T00:00:00+00:00",
				DeadlineFrom:      "2024-03-01T00:00:00+00:00",
				DeadlineTo:        "2024-04-01T00:00:00+00:00",
				ListOptions:       ListOptions{Limit: 10, Offset: 5},
			},
			out: "assigneeIds=4%2C5&createdAtFrom=2024-01-01T00%3A00%3A00%2B00%3A00&createdAtTo=2024-02-01T00%3A00%3A00%2B00%3A00" +
				"&creatorIds=6&deadlineFrom=2024-03-01T00%3A00%3A00%2B00%3A00&deadlineTo=2024-04-01T00%3A00%3A00%2B00%3A00" +
				"&groupIds=3&limit=10&offset=5&orderBy=createdAt+desc%2Ctitle&projectIds=1%2C2&sourceLanguageIds=en" +
				"&status=todo%2Cin_progress&targetLanguageIds=uk%2Cde&type=0%2C3",
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

func TestUserTasksListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *UserTasksListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &UserTasksListOptions{},
		},
		{
			name: "with isArchived = 0",
			opts: &UserTasksListOptions{IsArchived: toPtr(0)},
			out:  "isArchived=0",
		},
		{
			name: "all options",
			opts: &UserTasksListOptions{OrderBy: "createdAt desc,name", Status: []TaskStatus{TaskStatusTodo, TaskStatusDone},
				IsArchived: toPtr(1), ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "isArchived=1&limit=10&offset=5&orderBy=createdAt+desc%2Cname&status=todo%2Cdone",
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

func TestTaskSettingsTemplateAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *TaskSettingsTemplateAddRequest
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
			req:  &TaskSettingsTemplateAddRequest{},
			err:  "name is required",
		},
		{
			name: "empty languages",
			req:  &TaskSettingsTemplateAddRequest{Name: "Default template", Config: TaskSettingsTemplateConfig{}},
			err:  "config languages is required",
		},
		{
			name: "valid request",
			req: &TaskSettingsTemplateAddRequest{
				Name: "Default template",
				Config: TaskSettingsTemplateConfig{
					Languages: []TaskSettingsTemplateLanguage{
						{
							LanguageID: "uk",
							UserIDs:    []UserID{1, 2},
							TeamIDs:    []int{3, 4},
						},
					},
				},
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

func TestTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *TaskCreateForm
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
			req:  &TaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "missing languageId",
			req:  &TaskCreateForm{Title: "Test task"},
			err:  "languageId is required",
		},
		{
			name: "missing type",
			req:  &TaskCreateForm{Title: "Test task", LanguageID: "uk"},
			err:  "type is required and must be one of 0, 1",
		},
		{
			name: "invalid type",
			req:  &TaskCreateForm{Title: "Test task", LanguageID: "uk", Type: toPtr(TaskType(100))},
			err:  "type is required and must be one of 0, 1",
		},
		{
			name: "missing one of stringIds, fileIds, branchIds",
			req:  &TaskCreateForm{Title: "Test task", LanguageID: "uk", Type: toPtr(TaskTypeProofread)},
			err:  "one of stringIds, fileIds, directoryIds or branchIds is required",
		},
		{
			name: "valid request",
			req: &TaskCreateForm{Title: "Test task", LanguageID: "uk", Type: toPtr(TaskTypeProofread),
				FileIDs: []int{1, 2}},
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

func TestLanguageServiceTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *LanguageServiceTaskCreateForm
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
			req:  &LanguageServiceTaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "invalid language",
			req:  &LanguageServiceTaskCreateForm{Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "required type",
			req:  &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en"},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "invalid type",
			req:  &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en", Type: 0},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "required vendor",
			req:  &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor},
			err:  "vendor is required and must be \"crowdin_language_service\"",
		},
		{
			name: "invalid vendor",
			req: &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: "invalid"},
			err: "vendor is required and must be \"crowdin_language_service\"",
		},
		{
			name: "required fileIds/branchIds/stringIds",
			req: &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorCrowdinLanguageService},
			err: "one of stringIds, fileIds or branchIds is required",
		},
		{
			name: "valid validation",
			req: &LanguageServiceTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorCrowdinLanguageService, BranchIDs: []int{1, 2, 3}},
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

func TestVendorOhtTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *VendorOhtTaskCreateForm
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
			req:  &VendorOhtTaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "invalid language",
			req:  &VendorOhtTaskCreateForm{Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "required type",
			req:  &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en"},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "invalid type",
			req:  &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en", Type: 0},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "required vendor",
			req:  &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor},
			err:  "vendor is required and must be \"oht\"",
		},
		{
			name: "invalid vendor",
			req:  &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor, Vendor: "invalid"},
			err:  "vendor is required and must be \"oht\"",
		},
		{
			name: "required fileIds/branchIds/stringIds",
			req:  &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor, Vendor: TaskVendorOht},
			err:  "one of stringIds, fileIds or branchIds is required",
		},
		{
			name: "valid validation",
			req: &VendorOhtTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorOht, StringIDs: []int{1, 2, 3}},
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

func TestVendorGengoTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *VendorGengoTaskCreateForm
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
			req:  &VendorGengoTaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "invalid language",
			req:  &VendorGengoTaskCreateForm{Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "required type",
			req:  &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en"},
			err:  "type is required and must be 2",
		},
		{
			name: "invalid type",
			req:  &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en", Type: 0},
			err:  "type is required and must be 2",
		},
		{
			name: "required vendor",
			req:  &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor},
			err:  "vendor is required and must be \"gengo\"",
		},
		{
			name: "invalid vendor",
			req:  &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor, Vendor: "invalid"},
			err:  "vendor is required and must be \"gengo\"",
		},
		{
			name: "required fileIds/branchIds/stringIds",
			req:  &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor, Vendor: TaskVendorGengo},
			err:  "one of stringIds, fileIds or branchIds is required",
		},
		{
			name: "valid validation",
			req: &VendorGengoTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorGengo, FileIDs: []int{1, 2, 3}},
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

func TestVendorManualTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *VendorManualTaskCreateForm
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
			req:  &VendorManualTaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "invalid language",
			req:  &VendorManualTaskCreateForm{Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "required type",
			req:  &VendorManualTaskCreateForm{Title: "French", LanguageID: "en"},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "invalid type",
			req:  &VendorManualTaskCreateForm{Title: "French", LanguageID: "en", Type: 0},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "required vendor",
			req:  &VendorManualTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor},
			err:  "vendor is required",
		},
		{
			name: "required fileIds/branchIds/stringIds",
			req:  &VendorManualTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor, Vendor: TaskVendorAlconost},
			err:  "one of stringIds, fileIds or branchIds is required",
		},

		{
			name: "valid validation",
			req: &VendorManualTaskCreateForm{Title: "French", LanguageID: "en", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorAlconost, BranchIDs: []int{1, 2, 3}},
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

func TestPendingTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PendingTaskCreateForm
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
			req:  &PendingTaskCreateForm{},
			err:  "precedingTaskId is required",
		},
		{
			name: "invalid type",
			req:  &PendingTaskCreateForm{PrecedingTaskID: 1},
			err:  "type is required and must be one of 1, 3",
		},
		{
			name: "required title",
			req:  &PendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofread},
			err:  "title is required",
		},
		{
			name: "vendor required for type 3",
			req:  &PendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor, Title: "French"},
			err:  "vendor is required when type is 3",
		},
		{
			name:  "valid validation",
			req:   &PendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofread, Title: "French"},
			valid: true,
		},
		{
			name: "valid validation with vendor",
			req: &PendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor, Title: "French",
				Vendor: TaskVendorGengo},
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

func TestVendorTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *VendorTaskCreateForm
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
			req:  &VendorTaskCreateForm{},
			err:  "title is required",
		},
		{
			name: "required languageId",
			req:  &VendorTaskCreateForm{Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "invalid type",
			req:  &VendorTaskCreateForm{Title: "French", LanguageID: "fr", Type: TaskTypeProofread},
			err:  "type is required and must be one of 2, 3",
		},
		{
			name: "required vendor",
			req:  &VendorTaskCreateForm{Title: "French", LanguageID: "fr", Type: TaskTypeTranslateByVendor},
			err:  "vendor is required",
		},
		{
			name: "required content",
			req: &VendorTaskCreateForm{Title: "French", LanguageID: "fr", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorUndertow},
			err: "one of stringIds, fileIds, directoryIds or branchIds is required",
		},
		{
			name: "valid with directoryIds",
			req: &VendorTaskCreateForm{Title: "French", LanguageID: "fr", Type: TaskTypeProofreadByVendor,
				Vendor: TaskVendor("custom_vendor"), DirectoryIDs: []int{1}},
			valid: true,
		},
		{
			name: "valid with stringIds",
			req: &VendorTaskCreateForm{Title: "French", LanguageID: "fr", Type: TaskTypeTranslateByVendor,
				Vendor: TaskVendorCrowdinLanguageService, StringIDs: []int{1, 2}},
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

func TestLanguageServicePendingTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *LanguageServicePendingTaskCreateForm
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
			req:  &LanguageServicePendingTaskCreateForm{},
			err:  "precedingTaskId is required",
		},
		{
			name: "invalid type",
			req:  &LanguageServicePendingTaskCreateForm{PrecedingTaskID: 1},
			err:  "type is required and must be 3",
		},
		{
			name: "required title",
			req:  &LanguageServicePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor},
			err:  "title is required",
		},
		{
			name:  "valid validation",
			req:   &LanguageServicePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor, Title: "French"},
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

func TestVendorManualPendingTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *VendorManualPendingTaskCreateForm
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
			req:  &VendorManualPendingTaskCreateForm{},
			err:  "precedingTaskId is required",
		},
		{
			name: "invalid type",
			req:  &VendorManualPendingTaskCreateForm{PrecedingTaskID: 1},
			err:  "type is required and must be 3",
		},
		{
			name: "required vendor",
			req:  &VendorManualPendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor},
			err:  "vendor is required",
		},
		{
			name: "required title",
			req: &VendorManualPendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor,
				Vendor: TaskVendorAlconost},
			err: "title is required",
		},
		{
			name: "valid validation",
			req: &VendorManualPendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor,
				Vendor: TaskVendorAlconost, Title: "French"},
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

func TestEnterpriseTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *EnterpriseTaskCreateForm
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
			req:  &EnterpriseTaskCreateForm{},
			err:  "workflowStepId or type is required",
		},
		{
			name: "one of workflowStepId or type is required",
			req:  &EnterpriseTaskCreateForm{WorkflowStepID: 1, Type: toPtr(TaskTypeProofread)},
			err:  "workflowStepId and type can't be used in the same request",
		},
		{
			name: "invalid type",
			req:  &EnterpriseTaskCreateForm{Type: toPtr(TaskType(10))},
			err:  "type must be one of 0, 1",
		},
		{
			name: "required title",
			req:  &EnterpriseTaskCreateForm{Type: toPtr(TaskTypeTranslate)},
			err:  "title is required",
		},
		{
			name: "required language",
			req:  &EnterpriseTaskCreateForm{WorkflowStepID: 1, Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "stringIds or fileIds is required",
			req:  &EnterpriseTaskCreateForm{WorkflowStepID: 1, Title: "French", LanguageID: "en"},
			err:  "one of stringIds, fileIds, directoryIds or branchIds is required",
		},
		{
			name: "valid data validation",
			req: &EnterpriseTaskCreateForm{
				WorkflowStepID: 1,
				Title:          "French",
				LanguageID:     "en",
				FileIDs:        []int{1},
				Fields: map[string]any{
					"key_1": "value",
					"key_2": 2,
					"key_3": true,
					"key_4": []string{"a", "b"},
				},
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

func TestEnterpriseVendorTaskCreateFormValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *EnterpriseVendorTaskCreateForm
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
			req:  &EnterpriseVendorTaskCreateForm{},
			err:  "workflowStepId is required",
		},
		{
			name: "required title",
			req:  &EnterpriseVendorTaskCreateForm{WorkflowStepID: 1},
			err:  "title is required",
		},
		{
			name: "required language",
			req:  &EnterpriseVendorTaskCreateForm{WorkflowStepID: 1, Title: "French"},
			err:  "languageId is required",
		},
		{
			name: "stringIds or fileIds is required",
			req:  &EnterpriseVendorTaskCreateForm{WorkflowStepID: 1, Title: "French", LanguageID: "en"},
			err:  "one of stringIds, fileIds, directoryIds or branchIds is required",
		},
		{
			name: "valid data validation",
			req: &EnterpriseVendorTaskCreateForm{
				WorkflowStepID: 1,
				Title:          "French",
				LanguageID:     "en",
				StringIDs:      []int{1},
				Fields:         nil,
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

func TestTasksService_Add_EnterprisePendingTaskCreateForm_WithRequestValidation(t *testing.T) {
	tests := []struct {
		name  string
		req   *EnterprisePendingTaskCreateForm
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
			req:  &EnterprisePendingTaskCreateForm{},
			err:  "precedingTaskId is required",
		},
		{
			name: "invalid type",
			req:  &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1},
			err:  "type (one of 1, 3) or workflowStepId is required",
		},
		{
			name: "invalid title",
			req:  &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofread},
			err:  "title is required",
		},
		{
			name: "type and workflowStepId together",
			req: &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofread, WorkflowStepID: 2,
				Title: "French"},
			err: "workflowStepId and type can't be used in the same request",
		},
		{
			name: "vendor required for type 3",
			req:  &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor, Title: "French"},
			err:  "vendor is required when type is 3",
		},
		{
			name:  "pass validation",
			req:   &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofread, Title: "French"},
			valid: true,
		},
		{
			name:  "pass validation with workflowStepId",
			req:   &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, WorkflowStepID: 2, Title: "French"},
			valid: true,
		},
		{
			name: "pass validation with vendor",
			req: &EnterprisePendingTaskCreateForm{PrecedingTaskID: 1, Type: TaskTypeProofreadByVendor,
				Vendor: "gengo", Title: "French"},
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

func TestTaskCommentAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *TaskCommentAddRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "valid request",
			req: &TaskCommentAddRequest{
				Text:      "Work in task",
				TimeSpent: 3600,
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
