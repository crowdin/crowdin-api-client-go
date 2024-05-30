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

// https://developer.crowdin.com/api/v2/#operation/api.ai.prompts.getMany
func (s *AIService) ListPrompts(ctx context.Context, opt *model.AIPromtsListOptions) ([]*model.Prompt, *Response, error) {
	res := new(model.PromptsListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/users/ai/prompts", opt, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Prompt, 0, len(res.Data))
	for _, promt := range res.Data {
		list = append(list, promt.Data)
	}

	return list, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.get
func (s *AIService) GetPrompt(ctx context.Context, userID, promptID int) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/users/%d/ai/prompts/%d", userID, promptID), nil, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.post
func (s *AIService) AddPrompt(ctx context.Context, userID int, req *model.PromptAddRequest) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/users/%d/ai/prompts", userID), req, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.patch
func (s *AIService) EditPrompt(ctx context.Context, userID, promptID int, req []*model.UpdateRequest) (*model.Prompt, *Response, error) {
	res := new(model.PromptResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/users/%d/ai/prompts/%d", userID, promptID), req, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.prompts.delete
func (s *AIService) DeletePrompt(ctx context.Context, userID, promptID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/users/%d/ai/prompts/%d", userID, promptID), nil)
}

// https://developer.crowdin.com/api/v2/#operation/api.ai.providers.getMany
func (s *AIService) ListProviders(ctx context.Context, opt *model.ListOptions) ([]*model.Provider, *Response, error) {
	res := new(model.ProvidersListResponse)
	resp, err := s.client.Get(ctx, "/api/v2/users/ai/providers", opt, res)
	if err != nil {
		return nil, resp, err
	}

	list := make([]*model.Provider, 0, len(res.Data))
	for _, provider := range res.Data {
		list = append(list, provider.Data)
	}

	return list, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.get
func (s *AIService) GetProvider(ctx context.Context, userID, providerID int) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers/%d", userID, providerID), nil, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.post
func (s *AIService) AddProvider(ctx context.Context, userID int, req *model.ProviderAddRequest) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers", userID), req, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.patch
func (s *AIService) EditProvider(ctx context.Context, userID, providerID int, req []*model.UpdateRequest) (*model.Provider, *Response, error) {
	res := new(model.ProviderResponse)
	resp, err := s.client.Patch(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers/%d", userID, providerID), req, res)

	return res.Data, resp, err
}

// https://developer.crowdin.com/api/v2/#operation/api.users.ai.providers.delete
func (s *AIService) DeleteProvider(ctx context.Context, userID, providerID int) (*Response, error) {
	return s.client.Delete(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers/%d", userID, providerID), nil)
}

// https://developer.crowdin.com/api/v2/#operation/api.ai.providers.models.getMany
func (s *AIService) ListProviderModels(ctx context.Context, userID, providerID int) ([]*model.ProviderModel, *Response, error) {
	res := new(model.ProviderModelsListResponse)
	resp, err := s.client.Get(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers/%d/models", userID, providerID), nil, res)
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
func (s *AIService) CreateProxyChatCompletion(ctx context.Context, userID, providerID int, req *model.CreateProxyChatCompletionRequest) (
	*model.ProxyChatCompletion, *Response, error,
) {
	res := new(model.ProxyChatCompletionResponse)
	resp, err := s.client.Post(ctx, fmt.Sprintf("/api/v2/users/%d/ai/providers/%d/chat/completions", userID, providerID), req, res)

	return res.Data, resp, err
}
