package crowdin

import (
	"context"
	"fmt"
	"strings"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// AIService provides access to the AI related methods of the CrowdIn API.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/AI
type AIService struct {
	client *Client
}

// GenerateFineTuningDataset generates a new AI Prompt Fine-Tuning Dataset.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.ai.prompts.fine-tuning.datasets.post
func (s *AIService) GenerateFineTuningDataset(ctx context.Context, aiPromptID, userID int, req *model.FineTuningDatasetAttributes) (
	*model.FineTuningDataset, *Response, error,
) {
	res := new(model.FineTuningDatasetResponse)
	resp, err := s.client.Post(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/datasets", aiPromptID), userID), req, res)

	return res.Data, resp, err
}

// GetFineTuningDatasetGenerationStatus returns the status of the AI Prompt Fine-Tuning Dataset generation.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.users.ai.prompts.fine-tuning.datasets.get
func (s *AIService) GetFineTuningDatasetGenerationStatus(ctx context.Context, aiPromptID int, jobIdentifier string, userID int) (
	*model.FineTuningDataset, *Response, error,
) {
	res := new(model.FineTuningDatasetResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/datasets/%s", aiPromptID, jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// DownloadFineTuningDataset returns a download link for the AI Prompt Fine-Tuning Dataset.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.users.ai.prompts.fine-tuning.datasets.download.get
func (s *AIService) DownloadFineTuningDataset(ctx context.Context, aiPromptID int, jobIdentifier string, userID int) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/datasets/%s/download", aiPromptID, jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// ListFineTuningJobs returns a list of AI Prompt Fine-Tuning Jobs.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.ai.prompts.fine-tuning.jobs.getMany
func (s *AIService) ListFineTuningJobs(ctx context.Context, userID int, opts *model.FineTuningJobsListOptions) (
	[]*model.FineTuningJob, *Response, error,
) {
	res := new(model.FineTuningJobsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("prompts/fine-tuning/jobs", userID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.FineTuningJob, 0, len(res.Data))
	for _, job := range res.Data {
		list = append(list, job.Data)
	}

	return list, resp, err
}

// ListFineTuningEvents returns a list of AI Prompt Fine-Tuning Events.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.ai.prompts.fine-tuning.jobs.events.getMany
func (s *AIService) ListFineTuningEvents(ctx context.Context, aiPromptID int, jobIdentifier string, userID int) (
	[]*model.FineTuningEvent, *Response, error,
) {
	res := new(model.FineTuningEventsListResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/jobs/%s/events", aiPromptID, jobIdentifier), userID), nil, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.FineTuningEvent, 0, len(res.Data))
	for _, event := range res.Data {
		list = append(list, event.Data)
	}

	return list, resp, err
}

// CreateFineTuningJob creates a new AI Prompt Fine-Tuning Job.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.ai.prompts.fine-tuning.jobs.post
func (s *AIService) CreateFineTuningJob(ctx context.Context, aiPromptID, userID int, req *model.FineTuningJobCreateRequest) (
	*model.FineTuningJob, *Response, error,
) {
	res := new(model.FineTuningJobResponse)
	resp, err := s.client.Post(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/jobs", aiPromptID), userID), req, res)

	return res.Data, resp, err
}

// GetFineTuningJobStatus returns the status of the AI Prompt Fine-Tuning Job.
//
// https://support.crowdin.com/developer/api/v2/#tag/AI/operation/api.users.ai.prompts.fine-tuning.jobs.get
func (s *AIService) GetFineTuningJobStatus(ctx context.Context, aiPromptID int, jobIdentifier string, userID int) (
	*model.FineTuningJob, *Response, error,
) {
	res := new(model.FineTuningJobResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/fine-tuning/jobs/%s", aiPromptID, jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// ListPrompts returns a list of AI prompts.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.ai.prompts.getMany
func (s *AIService) ListPrompts(ctx context.Context, userID int, opt *model.AIPromtsListOptions) ([]*model.Prompt, *Response, error) {
	res := new(model.PromptsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("prompts", userID), opt, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Prompt, 0, len(res.Data))
	for _, promt := range res.Data {
		list = append(list, promt.Data)
	}

	return list, resp, err
}

// GetPrompt retrieves a single AI prompt.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.get
func (s *AIService) GetPrompt(ctx context.Context, promptID, userID int) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d", promptID), userID), nil, res)

	return res.Data, resp, err
}

// AddPrompt adds a new AI prompt.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.post
func (s *AIService) AddPrompt(ctx context.Context, userID int, req *model.PromptAddRequest) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Post(ctx, s.getPath("prompts", userID), req, res)

	return res.Data, resp, err
}

// EditPrompt updates an existing AI prompt.
// For the Enterprise client, set the userID to 0.
//
// Request body:
//   - Op (string): operation to perform. Enum: replace, test.
//   - Path (string<json-pointer>): path to the field to update. Enum: "/name", "/action",
//     "/aiProviderId", "/aiModelId", "/enabledProjectIds", "/config".
//   - Value (any): new value to set.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.patch
func (s *AIService) EditPrompt(ctx context.Context, promptID, userID int, req []*model.UpdateRequest) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Patch(ctx, s.getPath(fmt.Sprintf("prompts/%d", promptID), userID), req, res)

	return res.Data, resp, err
}

// DeletePrompt deletes an existing AI prompt.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.delete
func (s *AIService) DeletePrompt(ctx context.Context, promptID, userID int) (*Response, error) {
	return s.client.Delete(ctx, s.getPath(fmt.Sprintf("prompts/%d", promptID), userID), nil)
}

// ListProviders returns a list of AI providers.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.ai.providers.getMany
func (s *AIService) ListProviders(ctx context.Context, userID int, opt *model.ListOptions) ([]*model.Provider, *Response, error) {
	res := new(model.ProvidersListResponse)
	resp, err := s.client.Get(ctx, s.getPath("providers", userID), opt, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Provider, 0, len(res.Data))
	for _, provider := range res.Data {
		list = append(list, provider.Data)
	}

	return list, resp, err
}

// GetProvider returns a single AI provider.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.get
func (s *AIService) GetProvider(ctx context.Context, providerID, userID int) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("providers/%d", providerID), userID), nil, res)

	return res.Data, resp, err
}

// AddProvider adds a new AI provider.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.post
func (s *AIService) AddProvider(ctx context.Context, userID int, req *model.ProviderAddRequest) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Post(ctx, s.getPath("providers", userID), req, res)

	return res.Data, resp, err
}

// EditProvider updates an existing AI provider.
// For the Enterprise client, set the userID to 0.
//
// Request body:
//   - Op (string): operation to perform. Enum: replace, test.
//   - Path (string<json-pointer>): path to the field to update. Enum: "/name", "/type",
//     "/credentials", "/config", "/isEnabled", "/useSystemCredentials".
//   - Value (any): new value to set.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.patch
func (s *AIService) EditProvider(ctx context.Context, providerID, userID int, req []*model.UpdateRequest) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Patch(ctx, s.getPath(fmt.Sprintf("providers/%d", providerID), userID), req, res)

	return res.Data, resp, err
}

// DeleteProvider deletes an existing AI provider.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.delete
func (s *AIService) DeleteProvider(ctx context.Context, providerID, userID int) (*Response, error) {
	return s.client.Delete(ctx, s.getPath(fmt.Sprintf("providers/%d", providerID), userID), nil)
}

// ListProviderModels returns a list of AI provider models.
// For the Enterprise client, set the userID to 0.
//
// https://developer.crowdin.com/api/v2/#operation/api.ai.providers.models.getMany
func (s *AIService) ListProviderModels(ctx context.Context, providerID, userID int) ([]*model.ProviderModel, *Response, error) {
	res := new(model.ProviderModelsListResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("providers/%d/models", providerID), userID), nil, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ProviderModel, 0, len(res.Data))
	for _, model := range res.Data {
		list = append(list, model.Data)
	}

	return list, resp, err
}

// CreateProxyChatCompletion creates a new chat completion.
//
// This API method serves as an intermediary, forwarding your requests directly to the selected provider.
// Please refer to the documentation for the specific provider you use to determine the required payload format.
//
// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.chat.completions.post
func (s *AIService) CreateProxyChatCompletion(ctx context.Context, providerID, userID int, req *model.CreateProxyChatCompletionRequest) (
	*model.ProxyChatCompletion, *Response, error,
) {
	res := new(model.ProxyChatCompletionResponse)
	resp, err := s.client.Post(ctx, s.getPath(fmt.Sprintf("providers/%d/chat/completions", providerID), userID), req, res)

	return res.Data, resp, err
}

// ClonePrompt clones an existing AI prompt.
// For the Enterprise client, set the userID to 0.
// If req is nil, the prompt is cloned with the default name.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.prompts.clones.post
func (s *AIService) ClonePrompt(ctx context.Context, promptID, userID int, req *model.PromptCloneRequest) (*model.Prompt, *Response, error) {
	if req == nil {
		req = &model.PromptCloneRequest{}
	}

	res := new(model.PromptResponse)
	resp, err := s.client.Post(ctx, s.getPath(fmt.Sprintf("prompts/%d/clones", promptID), userID), req, res)

	return res.Data, resp, err
}

// GeneratePromptCompletion generates a new AI prompt completion.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.completions.post
func (s *AIService) GeneratePromptCompletion(ctx context.Context, promptID, userID int, req *model.PromptCompletionRequest) (
	*model.PromptCompletion, *Response, error,
) {
	res := new(model.PromptCompletionResponse)
	resp, err := s.client.Post(ctx, s.getPath(fmt.Sprintf("prompts/%d/completions", promptID), userID), req, res)

	return res.Data, resp, err
}

// GetPromptCompletionStatus returns the status of an AI prompt completion.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.prompts.completions.get
func (s *AIService) GetPromptCompletionStatus(ctx context.Context, promptID int, completionID string, userID int) (
	*model.PromptCompletion, *Response, error,
) {
	res := new(model.PromptCompletionResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/completions/%s", promptID, completionID), userID), nil, res)

	return res.Data, resp, err
}

// CancelPromptCompletion cancels an AI prompt completion.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.prompts.completions.delete
func (s *AIService) CancelPromptCompletion(ctx context.Context, promptID int, completionID string, userID int) (*Response, error) {
	return s.client.Delete(ctx, s.getPath(fmt.Sprintf("prompts/%d/completions/%s", promptID, completionID), userID), nil)
}

// DownloadPromptCompletion returns a download link for an AI prompt completion.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.prompts.completions.download.download
func (s *AIService) DownloadPromptCompletion(ctx context.Context, promptID int, completionID string, userID int) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("prompts/%d/completions/%s/download", promptID, completionID), userID), nil, res)

	return res.Data, resp, err
}

// TranslateFile starts an AI file translation job.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.file-translations.post
func (s *AIService) TranslateFile(ctx context.Context, userID int, req *model.AIFileTranslationRequest) (
	*model.AIFileTranslation, *Response, error,
) {
	res := new(model.AIFileTranslationResponse)
	resp, err := s.client.Post(ctx, s.getPath("file-translations", userID), req, res)

	return res.Data, resp, err
}

// GetFileTranslationStatus returns the status of an AI file translation job.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.file-translations.get
func (s *AIService) GetFileTranslationStatus(ctx context.Context, jobIdentifier string, userID int) (
	*model.AIFileTranslation, *Response, error,
) {
	res := new(model.AIFileTranslationResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("file-translations/%s", jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// CancelFileTranslation cancels an AI file translation job.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.file-translations.delete
func (s *AIService) CancelFileTranslation(ctx context.Context, jobIdentifier string, userID int) (*Response, error) {
	return s.client.Delete(ctx, s.getPath(fmt.Sprintf("file-translations/%s", jobIdentifier), userID), nil)
}

// DownloadTranslatedFile returns a download link for the file translated by AI.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.file-translations.download
func (s *AIService) DownloadTranslatedFile(ctx context.Context, jobIdentifier string, userID int) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("file-translations/%s/download", jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// DownloadTranslatedFileStrings returns a download link for the strings
// of the file translated by AI.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.file-translations.download-strings
func (s *AIService) DownloadTranslatedFileStrings(ctx context.Context, jobIdentifier string, userID int) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("file-translations/%s/translations", jobIdentifier), userID), nil, res)

	return res.Data, resp, err
}

// TranslateStrings translates strings with AI.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.translate.strings.post
func (s *AIService) TranslateStrings(ctx context.Context, userID int, req *model.AITranslateStringsRequest) (
	*model.AITranslateStrings, *Response, error,
) {
	res := new(model.AITranslateStringsResponse)
	resp, err := s.client.Post(ctx, s.getPath("translate", userID), req, res)

	return res.Data, resp, err
}

// ListAllProviderModels returns a list of models of all configured AI providers.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.models.crowdin.getMany
func (s *AIService) ListAllProviderModels(ctx context.Context, userID int) ([]*model.ProviderModel, *Response, error) {
	res := new(model.ProviderModelsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("providers/models", userID), nil, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.ProviderModel, 0, len(res.Data))
	for _, m := range res.Data {
		list = append(list, m.Data)
	}

	return list, resp, err
}

// ListSupportedProviderModels returns a list of supported AI provider models.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.supported-models.crowdin.getMany
func (s *AIService) ListSupportedProviderModels(ctx context.Context, userID int, opts *model.SupportedModelsListOptions) (
	[]*model.SupportedModel, *Response, error,
) {
	res := new(model.SupportedModelsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("providers/supported-models", userID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.SupportedModel, 0, len(res.Data))
	for _, m := range res.Data {
		list = append(list, m.Data)
	}

	return list, resp, err
}

// GatewayGet sends a GET request to the AI provider API through the AI Gateway.
// For the Enterprise client, set the userID to 0.
//
// The path is the raw provider API path after `/gateway/` and may contain
// slashes (e.g. `models`, `chat/completions`). The provider response is
// returned as is.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.gateway.crowdin.get
func (s *AIService) GatewayGet(ctx context.Context, providerID, userID int, path string) (map[string]any, *Response, error) {
	var res map[string]any
	resp, err := s.client.Get(ctx, s.gatewayPath(providerID, userID, path), nil, &res)

	return res, resp, err
}

// GatewayPost sends a POST request to the AI provider API through the AI Gateway.
// For the Enterprise client, set the userID to 0.
//
// The path is the raw provider API path after `/gateway/` and may contain
// slashes (e.g. `chat/completions`, `messages`). The body is forwarded to
// the provider and the provider response is returned as is.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.gateway.crowdin.post
func (s *AIService) GatewayPost(ctx context.Context, providerID, userID int, path string, body any) (map[string]any, *Response, error) {
	var res map[string]any
	resp, err := s.client.Post(ctx, s.gatewayPath(providerID, userID, path), body, &res)

	return res, resp, err
}

// GatewayPut sends a PUT request to the AI provider API through the AI Gateway.
// For the Enterprise client, set the userID to 0.
//
// The path is the raw provider API path after `/gateway/` and may contain
// slashes. The body is forwarded to the provider and the provider response
// is returned as is.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.gateway.crowdin.put
func (s *AIService) GatewayPut(ctx context.Context, providerID, userID int, path string, body any) (map[string]any, *Response, error) {
	var res map[string]any
	resp, err := s.client.Put(ctx, s.gatewayPath(providerID, userID, path), body, &res)

	return res, resp, err
}

// GatewayPatch sends a PATCH request to the AI provider API through the AI Gateway.
// For the Enterprise client, set the userID to 0.
//
// The path is the raw provider API path after `/gateway/` and may contain
// slashes. The body is forwarded to the provider and the provider response
// is returned as is.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.gateway.crowdin.patch
func (s *AIService) GatewayPatch(ctx context.Context, providerID, userID int, path string, body any) (map[string]any, *Response, error) {
	var res map[string]any
	resp, err := s.client.Patch(ctx, s.gatewayPath(providerID, userID, path), body, &res)

	return res, resp, err
}

// GatewayDelete sends a DELETE request to the AI provider API through the AI Gateway.
// For the Enterprise client, set the userID to 0.
//
// The path is the raw provider API path after `/gateway/` and may contain
// slashes. The provider response is returned as is.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.providers.gateway.crowdin.delete
func (s *AIService) GatewayDelete(ctx context.Context, providerID, userID int, path string) (map[string]any, *Response, error) {
	var res map[string]any
	resp, err := s.client.Delete(ctx, s.gatewayPath(providerID, userID, path), &res)

	return res, resp, err
}

// GenerateReport generates a new AI report.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.reports.post
func (s *AIService) GenerateReport(ctx context.Context, userID int, req *model.AIReportGenerateRequest) (*model.AIReport, *Response, error) {
	res := new(model.AIReportResponse)
	resp, err := s.client.Post(ctx, s.getPath("reports", userID), req, res)

	return res.Data, resp, err
}

// CheckReportStatus returns the generation status of an AI report.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.reports.get
func (s *AIService) CheckReportStatus(ctx context.Context, reportID string, userID int) (*model.AIReport, *Response, error) {
	res := new(model.AIReportResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("reports/%s", reportID), userID), nil, res)

	return res.Data, resp, err
}

// DownloadReport returns a download link for an AI report.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.reports.download.download
func (s *AIService) DownloadReport(ctx context.Context, reportID string, userID int) (*model.DownloadLink, *Response, error) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("reports/%s/download", reportID), userID), nil, res)

	return res.Data, resp, err
}

// ListRequestLogs returns a list of AI request logs.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.requestLogs.getMany
func (s *AIService) ListRequestLogs(ctx context.Context, userID int, opts *model.AIRequestLogsListOptions) (
	[]*model.AIRequestLog, *Response, error,
) {
	res := new(model.AIRequestLogsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("request-logs", userID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.AIRequestLog, 0, len(res.Data))
	for _, log := range res.Data {
		list = append(list, log.Data)
	}

	return list, resp, err
}

// ExportRequestLogs starts an export of the AI request logs.
// For the Enterprise client, set the userID to 0.
// If req is nil, every log entry is exported.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.requestLogs.exports.post
func (s *AIService) ExportRequestLogs(ctx context.Context, userID int, req *model.AIRequestLogsExportRequest) (
	*model.AIRequestLogsExport, *Response, error,
) {
	if req == nil {
		req = &model.AIRequestLogsExportRequest{}
	}

	res := new(model.AIRequestLogsExportResponse)
	resp, err := s.client.Post(ctx, s.getPath("request-logs/exports", userID), req, res)

	return res.Data, resp, err
}

// CheckRequestLogsExportStatus returns the status of an AI request logs export.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.requestLogs.exports.get
func (s *AIService) CheckRequestLogsExportStatus(ctx context.Context, exportID string, userID int) (
	*model.AIRequestLogsExport, *Response, error,
) {
	res := new(model.AIRequestLogsExportResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("request-logs/exports/%s", exportID), userID), nil, res)

	return res.Data, resp, err
}

// DownloadRequestLogsExport returns a download link for an AI request logs export.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.requestLogs.exports.download
func (s *AIService) DownloadRequestLogsExport(ctx context.Context, exportID string, userID int) (
	*model.DownloadLink, *Response, error,
) {
	res := new(model.DownloadLinkResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("request-logs/exports/%s/download", exportID), userID), nil, res)

	return res.Data, resp, err
}

// GetSettings returns the AI settings.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.settings.get
func (s *AIService) GetSettings(ctx context.Context, userID int) (*model.AISettings, *Response, error) {
	res := new(model.AISettingsResponse)
	resp, err := s.client.Get(ctx, s.getPath("settings", userID), nil, res)

	return res.Data, resp, err
}

// EditSettings updates the AI settings.
// For the Enterprise client, set the userID to 0.
//
// Request body:
//   - Op (string): operation to perform. Enum: replace, remove, test.
//   - Path (string<json-pointer>): path to the field to update. Enum: "/preTranslationAiPromptId",
//     "/editorSuggestionAiPromptId", "/alignmentActionAiPromptId" (Enterprise only),
//     "/qaCheckActionAiPromptId", "/contextReviewAiPromptId", "/dailyCostLimit",
//     "/monthlyCostLimit", "/userDailyCostLimit", "/userMonthlyCostLimit", "/perUserOverrides".
//   - Value (any): new value to set.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.users.ai.settings.patch
func (s *AIService) EditSettings(ctx context.Context, userID int, req []*model.UpdateRequest) (*model.AISettings, *Response, error) {
	res := new(model.AISettingsResponse)
	resp, err := s.client.Patch(ctx, s.getPath("settings", userID), req, res)

	return res.Data, resp, err
}

// GetProjectAISettings returns the AI settings of a project.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.projects.ai.settings.get
func (s *AIService) GetProjectAISettings(ctx context.Context, projectID int) (*model.ProjectAISettings, *Response, error) {
	res := new(model.ProjectAISettingsResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/projects/%d/ai/settings", projectID), nil, res)

	return res.Data, resp, err
}

// ListSnippets returns a list of AI snippets.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.snippets.getMany
func (s *AIService) ListSnippets(ctx context.Context, userID int, opts *model.ListOptions) ([]*model.AISnippet, *Response, error) {
	res := new(model.AISnippetsListResponse)
	resp, err := s.client.Get(ctx, s.getPath("settings/snippets", userID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.AISnippet, 0, len(res.Data))
	for _, snippet := range res.Data {
		list = append(list, snippet.Data)
	}

	return list, resp, err
}

// GetSnippet returns a single AI snippet.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.snippets.get
func (s *AIService) GetSnippet(ctx context.Context, snippetID, userID int) (*model.AISnippet, *Response, error) {
	res := new(model.AISnippetResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("settings/snippets/%d", snippetID), userID), nil, res)

	return res.Data, resp, err
}

// AddSnippet adds a new AI snippet.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.snippets.post
func (s *AIService) AddSnippet(ctx context.Context, userID int, req *model.AISnippetAddRequest) (*model.AISnippet, *Response, error) {
	res := new(model.AISnippetResponse)
	resp, err := s.client.Post(ctx, s.getPath("settings/snippets", userID), req, res)

	return res.Data, resp, err
}

// EditSnippet updates an existing AI snippet.
// For the Enterprise client, set the userID to 0.
//
// Request body:
//   - Op (string): operation to perform. Enum: replace, test.
//   - Path (string<json-pointer>): path to the field to update. Enum: "/description",
//     "/placeholder", "/value".
//   - Value (any): new value to set.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.snippets.patch
func (s *AIService) EditSnippet(ctx context.Context, snippetID, userID int, req []*model.UpdateRequest) (*model.AISnippet, *Response, error) {
	res := new(model.AISnippetResponse)
	resp, err := s.client.Patch(ctx, s.getPath(fmt.Sprintf("settings/snippets/%d", snippetID), userID), req, res)

	return res.Data, resp, err
}

// DeleteSnippet deletes an existing AI snippet.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.prompts.snippets.delete
func (s *AIService) DeleteSnippet(ctx context.Context, snippetID, userID int) (*Response, error) {
	return s.client.Delete(ctx, s.getPath(fmt.Sprintf("settings/snippets/%d", snippetID), userID), nil)
}

// ListUsageMembers returns a list of organization members with their AI usage and limits.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.usage.members.getMany
func (s *AIService) ListUsageMembers(ctx context.Context, userID int, opts *model.AIUsageMembersListOptions) (
	[]*model.AIUsageMember, *Response, error,
) {
	res := new(model.AIUsageMembersListResponse)
	resp, err := s.client.Get(ctx, s.getPath("usage/members", userID), opts, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.AIUsageMember, 0, len(res.Data))
	for _, member := range res.Data {
		list = append(list, member.Data)
	}

	return list, resp, err
}

// GetUsageMember returns the AI usage and limits of a single organization member.
// For the Enterprise client, set the userID to 0.
//
// https://support.crowdin.com/developer/api/v2/#operation/api.ai.usage.members.get
func (s *AIService) GetUsageMember(ctx context.Context, memberID, userID int) (*model.AIUsageMember, *Response, error) {
	res := new(model.AIUsageMemberResponse)
	resp, err := s.client.Get(ctx, s.getPath(fmt.Sprintf("usage/members/%d", memberID), userID), nil, res)

	return res.Data, resp, err
}

// gatewayPath returns the AI Gateway path for the given provider and raw provider API path.
func (s *AIService) gatewayPath(providerID, userID int, path string) string {
	return s.getPath(fmt.Sprintf("providers/%d/gateway/%s", providerID, strings.TrimPrefix(path, "/")), userID)
}

// getPath returns the path for the AI methods based on the user ID.
// If userID is 0 and organization is set, the Enterprise API path is used.
func (s *AIService) getPath(path string, userID int) string {
	if userID == 0 && s.client.organization != "" {
		return fmt.Sprintf("/api/v2/ai/%s", path)
	}

	return fmt.Sprintf("/api/v2/users/%d/ai/%s", userID, path)
}
