package model

import (
	"errors"
	"fmt"
	"net/url"
)

type PromptAction string

const (
	ActionPreTranslate PromptAction = "pre_translate"
	ActionAssist       PromptAction = "assist"
)

type PromptMode string

const (
	ModeBasic    PromptMode = "basic"
	ModeAdvanced PromptMode = "advanced"
)

type Prompt struct {
	ID                int          `json:"id"`
	Name              string       `json:"name"`
	Action            string       `json:"action"`
	AIProviderID      int          `json:"aiProviderId"`
	AIModelID         string       `json:"aiModelId"`
	IsEnabled         bool         `json:"isEnabled"`
	EnabledProjectIDs []int        `json:"enabledProjectIds"`
	Config            PromptConfig `json:"config"`
	CreatedAt         string       `json:"createdAt"`
	UpdatedAt         string       `json:"updatedAt"`
}

type PromptConfig struct {
	Mode                      PromptMode `json:"mode"`
	CompanyDescription        *string    `json:"companyDescription,omitempty"`
	ProjectDescription        *string    `json:"projectDescription,omitempty"`
	AudienceDescription       *string    `json:"audienceDescription,omitempty"`
	OtherLanguageTranslations struct {
		IsEnabled   *bool `json:"isEnabled,omitempty"`
		LanguageIDs []int `json:"languageIds,omitempty"`
	} `json:"otherLanguageTranslations,omitempty"`
	GlossaryTerms            *bool   `json:"glossaryTerms,omitempty"`
	TMSuggestions            *bool   `json:"tmSuggestions,omitempty"`
	FileContent              *bool   `json:"fileContent,omitempty"`
	FileContext              *bool   `json:"fileContext,omitempty"`
	PublicProjectDescription *bool   `json:"publicProjectDescription,omitempty"`
	SiblingsStrings          *bool   `json:"siblingsStrings,omitempty"`
	FilteredStrings          *bool   `json:"filteredStrings,omitempty"`
	Prompt                   *string `json:"prompt,omitempty"`
}

type PromptResponse struct {
	Data *Prompt `json:"data"`
}

type PromptsListResponse struct {
	Data []*PromptResponse `json:"data"`
}

type AIPromtsListOptions struct {
	// Allows to filter the prompts available for the specific action.
	ProjectID int `json:"projectId,omitempty"`
	// Allows to filter the prompts available for the specific action.
	// Enum: pre_translate, assist.
	Action PromptAction `json:"action,omitempty"`

	ListOptions
}

func (o *AIPromtsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.ProjectID > 0 {
		v.Add("projectId", fmt.Sprintf("%d", o.ProjectID))
	}
	if o.Action != "" {
		v.Add("action", string(o.Action))
	}

	return v, len(v) > 0
}

type PromptAddRequest struct {
	// AI prompt name.
	Name string `json:"name"`
	// AI prompt action. Enum: pre_translate, assist.
	Action PromptAction `json:"action"`
	// AI Provider identifier.
	AIProviderID int `json:"aiProviderId"`
	// AI Model identifier.
	AIModelID string `json:"aiModelId"`
	// Is AI prompt enabled. Default: true.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// List of enabled project IDs.
	EnabledProjectIDs []int `json:"enabledProjectIds,omitempty"`
	// AI prompt configuration.
	Config PromptConfig `json:"config"`
}

func (r *PromptAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

type Provider struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	Credentials          any    `json:"credentials"`
	Config               any    `json:"config"`
	IsEnabled            bool   `json:"isEnabled"`
	UseSystemCredentials bool   `json:"useSystemCredentials"`
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
}

type ProviderResponse struct {
	Data *Provider `json:"data"`
}

type ProvidersListResponse struct {
	Data []*ProviderResponse `json:"data"`
}

type ProviderType string

const (
	OpenAI       ProviderType = "open_ai"
	AzureOpenAI  ProviderType = "azure_open_ai"
	GoogleGemini ProviderType = "google_gemini"
	MistralAI    ProviderType = "mistral_ai"
	Anthropic    ProviderType = "anthropic"
	CustomAI     ProviderType = "custom_ai"
)

type ProviderAddRequest struct {
	// AI provider name.
	Name string `json:"name"`
	// AI provider type.
	// Enum: open_ai, azure_open_ai, google_gemini, mistral_ai, anthropic, custom_ai.
	Type ProviderType `json:"type"`
	// User’s own AI provider credentials.
	// Note: Use only if useSystemCredentials is set to `false`.
	Credentials any `json:"credentials,omitempty"`
	// AI provider configuration.
	Config any `json:"config,omitempty"`
	// Defines whether to AI provider is enabled. Default: true.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Enables the paid service AI provider via Crowdin.
	// Note: Set to true if `credentials` is not provided. Not supported
	// for `custom_ai` type.
	UseSystemCredentials *bool `json:"useSystemCredentials,omitempty"`
}

func (r *ProviderAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Type == "" {
		return errors.New("type is required")
	}

	return nil
}

type ProviderModel struct {
	ID int `json:"id"`
}

type ProviderModelResponse struct {
	Data *ProviderModel `json:"data"`
}

type ProviderModelsListResponse struct {
	Data []*ProviderModelResponse `json:"data"`
}

type ProxyChatCompletion struct{}

type ProxyChatCompletionResponse struct {
	Data *ProxyChatCompletion `json:"data"`
}

type CreateProxyChatCompletionRequest struct {
	// ID of the model to use.
	ModelID string `json:"modelId,omitempty"`
	// Tokens will be sent as data-only server-sent events as they become available,
	// with the stream terminated by a data: [DONE] message.
	Stream *bool `json:"stream,omitempty"`
}

func (r *CreateProxyChatCompletionRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}
