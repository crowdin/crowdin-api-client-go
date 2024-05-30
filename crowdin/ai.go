package crowdin

import (
	"context"
	"fmt"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
)

// AIService provides access to the AI related methods of the CrowdIn API.
//
// Crowdin API docs: https://developer.crowdin.com/api/v2/#tag/AI
type AIService struct {
	client *Client
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
//     "/aiProviderId", "/aiModelId", "/isEnabled", "/enabledProjectIds", "/config".
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

// getPath returns the path for the AI methods based on the user ID.
// If userID is 0 and organization is set, the Enterprise API path is used.
func (s *AIService) getPath(path string, userID int) string {
	if userID == 0 && s.client.organization != "" {
		return fmt.Sprintf("/api/v2/ai/%s", path)
	}

	return fmt.Sprintf("/api/v2/users/%d/ai/%s", userID, path)
}
