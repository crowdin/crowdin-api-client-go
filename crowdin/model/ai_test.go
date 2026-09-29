package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFineTuningDatasetAttributesValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *FineTuningDatasetAttributes
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty projectIds and tmIds",
			req:  &FineTuningDatasetAttributes{},
			err:  "projectIds or tmIds are required",
		},
		{
			name: "valid request",
			req: &FineTuningDatasetAttributes{
				ProjectIDs:       []int{1, 2},
				TMIDs:            []int{3, 4},
				Purpose:          "training",
				DateFrom:         "2024-09-23T11:26:54+00:00",
				DateTo:           "2024-09-23T11:26:54+00:00",
				MaxFileSize:      100,
				MinExamplesCount: 10,
				MaxExamplesCount: 100,
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

func TestFineTuningJobsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *FineTuningJobsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &FineTuningJobsListOptions{},
		},
		{
			name: "with one status",
			opt:  &FineTuningJobsListOptions{Statuses: []string{"created"}},
			out:  "statuses=created",
		},
		{
			name: "with multiple statuses",
			opt:  &FineTuningJobsListOptions{Statuses: []string{"created", "finished"}},
			out:  "statuses=created%2Cfinished",
		},
		{
			name: "with orderBy",
			opt:  &FineTuningJobsListOptions{OrderBy: "createdAt"},
			out:  "orderBy=createdAt",
		},
		{
			name: "with all options",
			opt: &FineTuningJobsListOptions{Statuses: []string{"in_progress"}, OrderBy: "createdAt",
				ListOptions: ListOptions{Limit: 10, Offset: 20}},
			out: "limit=10&offset=20&orderBy=createdAt&statuses=in_progress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestFineTuningJobCreateRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *FineTuningJobCreateRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty trainingOptions",
			req:  &FineTuningJobCreateRequest{},
			err:  "trainingOptions is required",
		},
		{
			name: "valid request",
			req: &FineTuningJobCreateRequest{
				TrainingOptions: &FineTuningJobOptions{},
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

func TestAIPromtsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *AIPromtsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &AIPromtsListOptions{},
		},
		{
			name: "with project ID",
			opt:  &AIPromtsListOptions{ProjectID: 1},
			out:  "projectId=1",
		},
		{
			name: "with action",
			opt:  &AIPromtsListOptions{Action: ActionAssist},
			out:  "action=assist",
		},
		{
			name: "with all options",
			opt:  &AIPromtsListOptions{ProjectID: 2, Action: ActionPreTranslate},
			out:  "action=pre_translate&projectId=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestPromptAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PromptAddRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty name",
			req:  &PromptAddRequest{},
			err:  "name is required",
		},
		{
			name: "empty action",
			req:  &PromptAddRequest{Name: "Pre-translate prompt"},
			err:  "action is required",
		},
		{
			name: "empty aiProviderId",
			req:  &PromptAddRequest{Name: "Pre-translate prompt", Action: "pre_translate"},
			err:  "aiProviderId is required",
		},
		{
			name: "empty aiModelId",
			req:  &PromptAddRequest{Name: "Pre-translate prompt", Action: "pre_translate", AIProviderID: 1},
			err:  "aiModelId is required",
		},
		{
			name: "empty config mode",
			req: &PromptAddRequest{Name: "Pre-translate prompt", Action: "pre_translate", AIProviderID: 1,
				AIModelID: "gpt-3.5-turbo-instruct", Config: PromptConfig{},
			},
			err: "config.mode is required",
		},
		{
			name: "valid request",
			req: &PromptAddRequest{Name: "Pre-translate prompt", Action: "pre_translate", AIProviderID: 1,
				AIModelID: "gpt-3.5-turbo-instruct", Config: PromptConfig{Mode: "turbo"}},
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

func TestProviderAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ProviderAddRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty name",
			req:  &ProviderAddRequest{},
			err:  "name is required",
		},
		{
			name: "empty type",
			req:  &ProviderAddRequest{Name: "OpenAI"},
			err:  "type is required",
		},
		{
			name: "valid request",
			req: &ProviderAddRequest{
				Name:        "OpenAI",
				Type:        OpenAI,
				Credentials: map[string]string{"api_key": "value123"},
				Config: ProviderConfig{
					ActionRules: []ActionRule{
						{
							Action: "pre_translate", AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
						},
					},
				},
				IsEnabled:            toPtr(true),
				UseSystemCredentials: toPtr(false),
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

func TestCreateProxyChatCompletionRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *CreateProxyChatCompletionRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name:  "valid request",
			req:   &CreateProxyChatCompletionRequest{ModelID: "gpt-3.5-turbo-instruct", Stream: toPtr(false)},
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

func TestPromptCloneRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PromptCloneRequest
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
			req:   &PromptCloneRequest{},
			valid: true,
		},
		{
			name:  "with name",
			req:   &PromptCloneRequest{Name: "Cloned prompt"},
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

func TestPromptCompletionRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *PromptCompletionRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "missing resources",
			req:  &PromptCompletionRequest{},
			err:  "resources is required",
		},
		{
			name: "missing projectId",
			req:  &PromptCompletionRequest{Resources: &PromptCompletionResources{}},
			err:  "resources.projectId is required",
		},
		{
			name: "stringIds without targetLanguageId",
			req:  &PromptCompletionRequest{Resources: &PromptCompletionResources{ProjectID: 1, StringIDs: []int{1}}},
			err:  "resources.targetLanguageId is required when stringIds is set",
		},
		{
			name: "tool without function name",
			req: &PromptCompletionRequest{
				Resources: &PromptCompletionResources{ProjectID: 1},
				Tools:     []*PromptCompletionTool{{Tool: &AITool{Type: "function", Function: &AIToolFunction{}}}},
			},
			err: "tools: tool.function.name is required",
		},
		{
			name: "nil tool",
			req: &PromptCompletionRequest{
				Resources: &PromptCompletionResources{ProjectID: 1},
				Tools:     []*PromptCompletionTool{{}},
			},
			err: "tools: tool.function.name is required",
		},
		{
			name: "tool without type",
			req: &PromptCompletionRequest{
				Resources: &PromptCompletionResources{ProjectID: 1},
				Tools:     []*PromptCompletionTool{{Tool: &AITool{Function: &AIToolFunction{Name: "fn"}}}},
			},
			err: "tools: tool.type is required",
		},
		{
			name: "valid request",
			req: &PromptCompletionRequest{
				Resources: &PromptCompletionResources{
					ProjectID:        1,
					TargetLanguageID: "uk",
					StringIDs:        []int{1, 2},
					OverridePromptValues: &PromptCompletionOverridePromptValues{
						ProjectDescription: "Project description",
					},
				},
				Tools:      []*PromptCompletionTool{{Tool: &AITool{Type: "function", Function: &AIToolFunction{Name: "fn"}}}},
				ToolChoice: "auto",
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

func TestAIFileTranslationRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *AIFileTranslationRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "missing storageId",
			req:  &AIFileTranslationRequest{TargetLanguageID: "uk"},
			err:  "storageId is required",
		},
		{
			name: "missing targetLanguageId",
			req:  &AIFileTranslationRequest{StorageID: 1},
			err:  "targetLanguageId is required",
		},
		{
			name: "parserVersion without type",
			req:  &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk", ParserVersion: 1},
			err:  "type is required when parserVersion is set",
		},
		{
			name: "too many attachments",
			req: &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk",
				AttachmentIDs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			err: "attachmentIds cannot contain more than 10 items",
		},
		{
			name: "aiPromptId with aiProviderId",
			req:  &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk", AIPromptID: 1, AIProviderID: 2},
			err:  "aiPromptId can't be used with aiProviderId or aiModelId",
		},
		{
			name: "aiProviderId without aiModelId",
			req:  &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk", AIProviderID: 2},
			err:  "aiProviderId and aiModelId must be used together",
		},
		{
			name:  "valid request without prompt or model",
			req:   &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk"},
			valid: true,
		},
		{
			name: "valid request with provider and model",
			req: &AIFileTranslationRequest{StorageID: 1, TargetLanguageID: "uk", Type: "json", ParserVersion: 1,
				AIProviderID: 2, AIModelID: "gpt-4.1", StyleGuideIDs: []int{3}},
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

func TestAITranslateStringsRequestValidate(t *testing.T) {
	tooMany := make([]string, 501)

	tests := []struct {
		name  string
		req   *AITranslateStringsRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "missing strings",
			req:  &AITranslateStringsRequest{TargetLanguageID: "uk", AIPromptID: 1},
			err:  "strings is required",
		},
		{
			name: "too many strings",
			req:  &AITranslateStringsRequest{Strings: tooMany, TargetLanguageID: "uk", AIPromptID: 1},
			err:  "strings cannot contain more than 500 items",
		},
		{
			name: "missing targetLanguageId",
			req:  &AITranslateStringsRequest{Strings: []string{"a"}, AIPromptID: 1},
			err:  "targetLanguageId is required",
		},
		{
			name: "too many attachments",
			req: &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk", AIPromptID: 1,
				AttachmentIDs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
			err: "attachmentIds cannot contain more than 10 items",
		},
		{
			name: "missing prompt and model",
			req:  &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk"},
			err:  "aiPromptId or aiProviderId with aiModelId is required",
		},
		{
			name: "aiModelId without aiProviderId",
			req:  &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk", AIModelID: "gpt-4.1"},
			err:  "aiProviderId and aiModelId must be used together",
		},
		{
			name: "aiPromptId with aiModelId",
			req: &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk", AIPromptID: 1,
				AIModelID: "gpt-4.1"},
			err: "aiPromptId can't be used with aiProviderId or aiModelId",
		},
		{
			name:  "valid request with prompt",
			req:   &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk", AIPromptID: 1},
			valid: true,
		},
		{
			name: "valid request with provider and model",
			req: &AITranslateStringsRequest{Strings: []string{"a"}, TargetLanguageID: "uk", AIProviderID: 1,
				AIModelID: "gpt-4.1", StyleGuideIDs: []int{2}},
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

func TestSupportedModelsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *SupportedModelsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &SupportedModelsListOptions{},
		},
		{
			name: "with enabled false",
			opt:  &SupportedModelsListOptions{Enabled: toPtr(false)},
			out:  "enabled=false",
		},
		{
			name: "with all options",
			opt: &SupportedModelsListOptions{ProviderType: OpenAI, Enabled: toPtr(true), OrderBy: "id desc",
				ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "enabled=true&limit=10&offset=5&orderBy=id+desc&providerType=open_ai",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestAIReportGenerateRequestValidate(t *testing.T) {
	schema := &AIReportSchema{DateFrom: "2024-01-23T07:00:14+00:00", DateTo: "2024-09-27T07:00:14+00:00"}

	tests := []struct {
		name  string
		req   *AIReportGenerateRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "missing type",
			req:  &AIReportGenerateRequest{Schema: schema},
			err:  `invalid type: "", must be one of tokens-usage-raw-data, costs-by-users`,
		},
		{
			name: "invalid type",
			req:  &AIReportGenerateRequest{Type: "unknown", Schema: schema},
			err:  `invalid type: "unknown", must be one of tokens-usage-raw-data, costs-by-users`,
		},
		{
			name: "missing schema",
			req:  &AIReportGenerateRequest{Type: AIReportTokensUsageRawData},
			err:  "schema is required",
		},
		{
			name: "missing dateFrom",
			req:  &AIReportGenerateRequest{Type: AIReportCostsByUsers, Schema: &AIReportSchema{DateTo: "2024-09-27"}},
			err:  "schema.dateFrom is required",
		},
		{
			name: "missing dateTo",
			req:  &AIReportGenerateRequest{Type: AIReportCostsByUsers, Schema: &AIReportSchema{DateFrom: "2024-01-23"}},
			err:  "schema.dateTo is required",
		},
		{
			name:  "valid tokens usage raw data request",
			req:   &AIReportGenerateRequest{Type: AIReportTokensUsageRawData, Schema: schema},
			valid: true,
		},
		{
			name: "valid costs by users request",
			req: &AIReportGenerateRequest{Type: AIReportCostsByUsers, Schema: &AIReportSchema{
				DateFrom: "2024-01-23", DateTo: "2024-09-27", Format: "csv", ProjectIDs: []int{1}, UserIDs: []int{2},
			}},
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

func TestAIRequestLogsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *AIRequestLogsListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &AIRequestLogsListOptions{},
		},
		{
			name: "with statuses",
			opt:  &AIRequestLogsListOptions{Statuses: []string{"success", "timeout"}},
			out:  "statuses=success%2Ctimeout",
		},
		{
			name: "with all options",
			opt: &AIRequestLogsListOptions{
				RequestID:         "req-1",
				ProjectID:         1,
				UserID:            2,
				AIProviderID:      3,
				Model:             "gpt-4.1",
				SourceAction:      "pre_translate:manual",
				PromptAction:      "pre_translate",
				Statuses:          []string{"error"},
				SystemCredentials: toPtr(true),
				IsAutoTriggered:   toPtr(false),
				TokenName:         "token",
				OAuthClientID:     "client",
				CreatedAfter:      "2024-01-01T00:00:00+00:00",
				CreatedBefore:     "2024-02-01T00:00:00+00:00",
				ListOptions:       ListOptions{Limit: 10, Offset: 20},
			},
			out: "aiProviderId=3&createdAfter=2024-01-01T00%3A00%3A00%2B00%3A00&createdBefore=2024-02-01T00%3A00%3A00%2B00%3A00" +
				"&isAutoTriggered=false&limit=10&model=gpt-4.1&oauthClientId=client&offset=20&projectId=1&promptAction=pre_translate" +
				"&requestId=req-1&sourceAction=pre_translate%3Amanual&statuses=error&systemCredentials=true&tokenName=token&userId=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}

func TestAIRequestLogsExportRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *AIRequestLogsExportRequest
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
			req:   &AIRequestLogsExportRequest{},
			valid: true,
		},
		{
			name:  "with filters",
			req:   &AIRequestLogsExportRequest{Format: "csv", ProjectID: 1, Statuses: []string{"error"}},
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

func TestAISnippetAddRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *AISnippetAddRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "missing description",
			req:  &AISnippetAddRequest{Placeholder: "%custom:a%", Value: "value"},
			err:  "description is required",
		},
		{
			name: "missing placeholder",
			req:  &AISnippetAddRequest{Description: "Description", Value: "value"},
			err:  "placeholder is required",
		},
		{
			name: "missing value",
			req:  &AISnippetAddRequest{Description: "Description", Placeholder: "%custom:a%"},
			err:  "value is required",
		},
		{
			name:  "valid request",
			req:   &AISnippetAddRequest{Description: "Description", Placeholder: "%custom:a%", Value: "value"},
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

func TestAIUsageMembersListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opt  *AIUsageMembersListOptions
		out  string
	}{
		{
			name: "nil options",
			opt:  nil,
		},
		{
			name: "empty options",
			opt:  &AIUsageMembersListOptions{},
		},
		{
			name: "with userIds",
			opt:  &AIUsageMembersListOptions{UserIDs: []int{1, 2}},
			out:  "userIds=1%2C2",
		},
		{
			name: "with all options",
			opt: &AIUsageMembersListOptions{UserIDs: []int{1}, OrderBy: "id desc",
				ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out: "limit=10&offset=5&orderBy=id+desc&userIds=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opt.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}
