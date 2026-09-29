package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// FineTuningDataset represents a fine-tuning dataset.
type FineTuningDataset struct {
	Identifier string                       `json:"identifier"`
	Status     string                       `json:"status"`
	Progress   int                          `json:"progress"`
	Attributes *FineTuningDatasetAttributes `json:"attributes"`
	CreatedAt  string                       `json:"createdAt"`
	UpdatedAt  string                       `json:"updatedAt"`
	StartedAt  string                       `json:"startedAt"`
	FinishedAt string                       `json:"finishedAt"`
}

// FineTuningDatasetAttributes represents the attributes of a fine-tuning dataset.
// It is used to generate a fine-tuning dataset.
type FineTuningDatasetAttributes struct {
	// Project identifiers from which the dataset will be generated.
	// Note: Required if `tmIds` is not provided.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// TM identifiers from which the dataset will be generated.
	// Note: This parameter is not supported for the prompt with the
	// external configuraion.
	TMIDs []int `json:"tmIds,omitempty"`
	// Purpose of the dataset. Enum: training, validation. Default: training.
	Purpose string `json:"purpose,omitempty"`
	// Start date for dataset generation.
	DateFrom string `json:"dateFrom,omitempty"`
	// End date for dataset generation.
	DateTo string `json:"dateTo,omitempty"`
	// Maximum dataset file size in bytes.
	// Note: If not provided, default limits based on the model will be applied.
	MaxFileSize int `json:"maxFileSize,omitempty"`
	// Minimum number of examples in the dataset.
	// Note: If not provided, default limits based on the model will be applied.
	MinExamplesCount int `json:"minExamplesCount,omitempty"`
	// Maximum number of examples in the dataset.
	// Note: If not provided, default limits based on the model will be applied.
	MaxExamplesCount int `json:"maxExamplesCount,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *FineTuningDatasetAttributes) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if len(r.ProjectIDs) == 0 && len(r.TMIDs) == 0 {
		return errors.New("projectIds or tmIds are required")
	}

	return nil
}

// FineTuningDatasetResponse defines the structure of a response when
// getting a single fine-tuning dataset.
type FineTuningDatasetResponse struct {
	Data *FineTuningDataset `json:"data"`
}

type (
	// FineTuningEvent represents a fine-tuning event.
	FineTuningEvent struct {
		ID        string               `json:"id"`
		Type      string               `json:"type"` // enum: message, metrics
		Message   string               `json:"message"`
		Data      *FineTuningEventData `json:"data,omitempty"`
		CreatedAt string               `json:"createdAt"`
	}

	// FineTuningEventData represents the data of a fine-tuning event.
	FineTuningEventData struct {
		Step               int     `json:"step"`
		TotalSteps         int     `json:"totalSteps"`
		TrainingLoss       float64 `json:"trainingLoss"`
		ValidationLoss     float64 `json:"validationLoss"`
		FullValidationLoss float64 `json:"fullValidationLoss"`
	}
)

// FineTuningEventsListResponse defines the structure of a response when
// getting a list of fine-tuning events.
type FineTuningEventsListResponse struct {
	Data []struct {
		Data *FineTuningEvent `json:"data"`
	} `json:"data"`
}

type (
	// FineTuningJob represents a fine-tuning job.
	FineTuningJob struct {
		Identifier string                   `json:"identifier"`
		Status     string                   `json:"status"`
		Progress   int                      `json:"progress"`
		Attributes *FineTuningJobAttributes `json:"attributes"`
		CreatedAt  string                   `json:"createdAt"`
		UpdatedAt  string                   `json:"updatedAt"`
		StartedAt  string                   `json:"startedAt"`
		FinishedAt string                   `json:"finishedAt"`
	}

	// FineTuningJobAttributes represents the attributes of a fine-tuning job.
	FineTuningJobAttributes struct {
		DryRun               bool                          `json:"dryRun"`
		AIPromptID           int                           `json:"aiPromptId"`
		Hyperparameters      *FineTuningJobHyperparameters `json:"hyperparameters"`
		TrainingOptions      *FineTuningJobOptions         `json:"trainingOptions"`
		ValidationOptions    *FineTuningJobOptions         `json:"validationOptions"`
		BaseModel            string                        `json:"baseModel"`
		FineTunedModel       string                        `json:"fineTunedModel"`
		TrainedTokensCount   int                           `json:"trainedTokensCount"`
		TrainingDatasetURL   string                        `json:"trainingDatasetUrl"`
		ValidationDatasetURL string                        `json:"validationDatasetUrl"`
		Metadata             *FineTuningJobMetadata        `json:"metadata"`
	}

	// FineTuningJobHyperparameters represents the hyperparameters of a fine-tuning job.
	FineTuningJobHyperparameters struct {
		// Number of examples in each batch. A larger batch size means that model
		// parameters are updated less frequently, but with lower variance.
		// Note: This parameter is not supported by Mistral AI.
		BatchSize int `json:"batchSize,omitempty"`
		// Scaling factor for the learning rate. A smaller learning rate may be useful
		// to avoid overfitting. Note: This parameter is not supported by Mistral AI.
		LearningRateMultiplier float64 `json:"learningRateMultiplier,omitempty"`
		// The number of epochs to train the model for. An epoch refers to one full
		// cycle through the training dataset.
		NEpochs int `json:"nEpochs,omitempty"`
	}

	// FineTuningJobOptions represents the options of a fine-tuning job.
	FineTuningJobOptions struct {
		// Project identifiers from which the dataset will be generated.
		// Note: Required if `tmIds` is not provided.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// TM identifiers from which the dataset will be generated.
		// Note: This parameteris not supported for the prompt with
		// external configuraion.
		TMIDs []int `json:"tmIds,omitempty"`
		// Start date for the dataset generation.
		DateFrom string `json:"dateFrom,omitempty"`
		// End date for the dataset generation.
		DateTo string `json:"dateTo,omitempty"`
		// Maximum dataset file size in bytes.
		// Note: If not provided, default limits based on the model will be applied.
		MaxFileSize int `json:"maxFileSize,omitempty"`
		// Minimum number of examples in the dataset.
		// Note: If not provided, default limits based on the model will be applied.
		MinExamplesCount int `json:"minExamplesCount,omitempty"`
		// Maximum number of examples in the dataset.
		// Note: If not provided, default limits based on the model will be applied.
		MaxExamplesCount int `json:"maxExamplesCount,omitempty"`
	}

	// FineTuningJobMetadata represents the metadata of a fine-tuning job.
	FineTuningJobMetadata struct {
		Cost         float64 `json:"cost"`
		CostCurrency string  `json:"costCurrency"`
	}
)

// FineTuningJobResponse defines the structure of a response when
// getting a fine-tuning job.
type FineTuningJobResponse struct {
	Data *FineTuningJob `json:"data"`
}

// FineTuningJobsListResponse defines the structure of a response when
// getting a list of fine-tuning jobs.
type FineTuningJobsListResponse struct {
	Data []*FineTuningJobResponse `json:"data"`
}

// FineTuningJobsListOptions specifies the optional parameters to the
// AIService.ListFineTuningJobs method.
type FineTuningJobsListOptions struct {
	// Filter the collection by the specified status. It can be one status or
	// a list of comma-separated ones.
	// Enum: created, in_progress, canceled, failed, finished.
	Statuses []string `json:"statuses,omitempty"`
	// Sort the collection by the specified field. Example: orderBy=createdAt desc.
	// Enum: createdAt, updatedAt, startedAt, finishedAt.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of FineTuningJobsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *FineTuningJobsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if len(o.Statuses) > 0 {
		v.Add("statuses", JoinSlice(o.Statuses))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// FineTuningJobCreateRequest defines the structure of a request to create a fine-tuning job.
type FineTuningJobCreateRequest struct {
	// Options for generating the fine-tuning dataset.
	TrainingOptions *FineTuningJobOptions `json:"trainingOptions"`
	// Options for generating the fine-tuning dataset.
	ValidationOptions *FineTuningJobOptions `json:"validationOptions,omitempty"`
	// The hyperparameters used for the fine-tuning job.
	Hyperparameters *FineTuningJobHyperparameters `json:"hyperparameters,omitempty"`
	// Simulate the fine-tuning of job creation without actually
	// creating them. Default: false.
	DryRun *bool `json:"dryRun,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *FineTuningJobCreateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.TrainingOptions == nil {
		return errors.New("trainingOptions is required")
	}

	return nil
}

// PromptAction represents the action of an AI prompt.
type PromptAction string

const (
	// ActionPreTranslate is the pre-translation action.
	ActionPreTranslate PromptAction = "pre_translate"
	// ActionAssist is the assist action.
	//
	// Deprecated: the assist action has been removed from the API.
	ActionAssist PromptAction = "assist"
	// ActionQACheck is the QA check action.
	ActionQACheck PromptAction = "qa_check"
	// ActionAlignment is the alignment action (Enterprise only).
	ActionAlignment PromptAction = "alignment"
)

// PromptMode represents the mode of an AI prompt configuration.
type PromptMode string

const (
	// ModeBasic is the basic prompt mode.
	ModeBasic PromptMode = "basic"
	// ModeAdvanced is the advanced prompt mode.
	ModeAdvanced PromptMode = "advanced"
	// ModeExternal is the external prompt mode (prompt provided by an application).
	ModeExternal PromptMode = "external"
)

// Prompt represents an AI prompt.
type Prompt struct {
	ID                int          `json:"id"`
	Name              string       `json:"name"`
	Action            string       `json:"action"`
	AIProviderID      int          `json:"aiProviderId"`
	AIModelID         string       `json:"aiModelId"`
	IsEnabled         bool         `json:"isEnabled"`
	EnabledProjectIDs []int        `json:"enabledProjectIds"`
	Config            PromptConfig `json:"config"`
	// Preview of the compiled prompt.
	PromptPreview *string `json:"promptPreview,omitempty"`
	// Whether fine-tuning is available for the prompt.
	IsFineTuningAvailable bool `json:"isFineTuningAvailable"`
	// Identifier of the user who created the prompt.
	CreatedBy *int `json:"createdBy,omitempty"`
	// Identifier of the user who last updated the prompt.
	UpdatedBy *int `json:"updatedBy,omitempty"`
	// Identifier of the user who last used the prompt.
	LastUsedBy *int `json:"lastUsedBy,omitempty"`
	// Date and time when the prompt was last used.
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	// Total number of times the prompt has been used.
	UsageCount int    `json:"usageCount"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type (
	// PromptConfig represents the configuration of an AI prompt.
	PromptConfig struct {
		// Prompt mode. Enum: basic, advanced, external.
		Mode PromptMode `json:"mode"`
		// Deprecated: use Snippets instead.
		CompanyDescription *string `json:"companyDescription,omitempty"`
		// Deprecated: use Snippets instead.
		ProjectDescription *string `json:"projectDescription,omitempty"`
		// Deprecated: use Snippets instead.
		AudienceDescription *string `json:"audienceDescription,omitempty"`
		// Array of snippet placeholders. Get via AIService.ListSnippets.
		Snippets                  []string                   `json:"snippets,omitempty"`
		OtherLanguageTranslations *OtherLanguageTranslations `json:"otherLanguageTranslations,omitempty"`
		GlossaryTerms             *bool                      `json:"glossaryTerms,omitempty"`
		TMSuggestions             *bool                      `json:"tmSuggestions,omitempty"`
		// Deprecated: this field is deprecated in the API.
		FileContent *bool `json:"fileContent,omitempty"`
		FileContext *bool `json:"fileContext,omitempty"`
		// Generate a file summary to use as context.
		GenerateFileSummary *bool `json:"generateFileSummary,omitempty"`
		// Include screenshots in the prompt context.
		Screenshots *bool `json:"screenshots,omitempty"`
		// Include the project name and description in the prompt context.
		ProjectContext *bool `json:"projectContext,omitempty"`
		// Include the organization name and internal description in the
		// prompt context. Enterprise only (basic mode).
		OrganizationContext *bool `json:"organizationContext,omitempty"`
		// Deprecated: use ProjectContext instead.
		PublicProjectDescription *bool `json:"publicProjectDescription,omitempty"`
		SiblingsStrings          *bool `json:"siblingsStrings,omitempty"`
		FilteredStrings          *bool `json:"filteredStrings,omitempty"`
		// Retry the translation when QA issues are found (pre-translate action).
		RetryOnQAIssues *bool `json:"retryOnQaIssues,omitempty"`
		// Evaluation steps (basic mode, QA check action).
		EvaluationSteps []string `json:"evaluationSteps,omitempty"`
		// Prompt text (advanced mode).
		Prompt *string `json:"prompt,omitempty"`
		// Application identifier (external mode).
		Identifier string `json:"identifier,omitempty"`
		// Module key (external mode).
		Key string `json:"key,omitempty"`
		// Options to compile the prompt (external mode).
		Options map[string]any `json:"options,omitempty"`
	}

	// OtherLanguageTranslations represents the other language translations
	// settings of an AI prompt configuration.
	OtherLanguageTranslations struct {
		IsEnabled   *bool    `json:"isEnabled,omitempty"`
		LanguageIDs []string `json:"languageIds,omitempty"`
	}
)

// PromptResponse defines the structure of a response when
// getting a single AI prompt.
type PromptResponse struct {
	Data *Prompt `json:"data"`
}

// PromptsListResponse defines the structure of a response when
// getting a list of AI prompts.
type PromptsListResponse struct {
	Data []*PromptResponse `json:"data"`
}

// AIPromtsListOptions specifies the optional parameters to the
// AIService.ListPrompts method.
type AIPromtsListOptions struct {
	// Allows to filter the prompts available for the specific action.
	ProjectID int `json:"projectId,omitempty"`
	// Allows to filter the prompts available for the specific action.
	// Enum: pre_translate, qa_check, alignment (Enterprise only).
	Action PromptAction `json:"action,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of AIPromtsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
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

// PromptAddRequest defines the structure of a request to add an AI prompt.
type PromptAddRequest struct {
	// AI prompt name.
	Name string `json:"name"`
	// AI prompt action. Enum: pre_translate, qa_check, alignment (Enterprise only).
	Action PromptAction `json:"action"`
	// AI Provider identifier.
	AIProviderID int `json:"aiProviderId"`
	// AI Model identifier.
	AIModelID string `json:"aiModelId"`
	// Is AI prompt enabled. Default: true.
	//
	// Deprecated: this field is deprecated in the API.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// List of enabled project IDs.
	EnabledProjectIDs []int `json:"enabledProjectIds,omitempty"`
	// AI prompt configuration.
	Config PromptConfig `json:"config"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PromptAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Action == "" {
		return errors.New("action is required")
	}
	if r.AIProviderID == 0 {
		return errors.New("aiProviderId is required")
	}
	if r.AIModelID == "" {
		return errors.New("aiModelId is required")
	}
	if r.Config.Mode == "" {
		return errors.New("config.mode is required")
	}

	return nil
}

// ProviderType represents the type of an AI provider.
type ProviderType string

const (
	OpenAI               ProviderType = "open_ai"
	AzureOpenAI          ProviderType = "azure_open_ai"
	GoogleGemini         ProviderType = "google_gemini"
	GoogleGeminiAIStudio ProviderType = "google_gemini_ai_studio"
	MistralAI            ProviderType = "mistral_ai"
	Anthropic            ProviderType = "anthropic"
	XAI                  ProviderType = "x_ai"
	Watsonx              ProviderType = "watsonx"
	DeepSeek             ProviderType = "deepseek"
	MicrosoftFoundry     ProviderType = "microsoft_foundry"
	CustomAI             ProviderType = "custom_ai"
)

// Provider represents an AI provider.
type Provider struct {
	ID   int          `json:"id"`
	Name string       `json:"name"`
	Type ProviderType `json:"type"`
	// Provider credentials with string values only (e.g. apiKey, baseUrl).
	// See RawCredentials for the full credentials object.
	Credentials map[string]string `json:"credentials"`
	// RawCredentials is the full credentials object. Depending on the provider
	// type, values can be strings, booleans (sendCustomHeadersOnly),
	// objects (headers, serviceAccountKey) or arrays (deployments).
	RawCredentials       map[string]any `json:"-"`
	Config               ProviderConfig `json:"config"`
	IsEnabled            bool           `json:"isEnabled"`
	UseSystemCredentials bool           `json:"useSystemCredentials"`
	CreatedAt            string         `json:"createdAt"`
	UpdatedAt            string         `json:"updatedAt"`
	// Number of prompts that use the provider.
	PromptsCount int `json:"promptsCount"`
}

// UnmarshalJSON decodes an AI provider. The full `credentials` object is stored
// in RawCredentials, and its string values are also copied to Credentials.
func (p *Provider) UnmarshalJSON(data []byte) error {
	type alias Provider
	aux := struct {
		*alias
		Credentials map[string]any `json:"credentials"`
	}{alias: (*alias)(p)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.RawCredentials, p.Credentials = aux.Credentials, nil
	if aux.Credentials != nil {
		p.Credentials = make(map[string]string, len(aux.Credentials))
		for k, v := range aux.Credentials {
			if s, ok := v.(string); ok {
				p.Credentials[k] = s
			}
		}
	}

	return nil
}

// ProviderResponse defines the structure of a response when
// getting a single AI provider.
type ProviderResponse struct {
	Data *Provider `json:"data"`
}

// ProvidersListResponse defines the structure of a response when
// getting a list of AI providers.
type ProvidersListResponse struct {
	Data []*ProviderResponse `json:"data"`
}

// ProviderAddRequest defines the structure of a request to add an AI provider.
type ProviderAddRequest struct {
	// AI provider name.
	Name string `json:"name"`
	// AI provider type.
	// Enum: open_ai, azure_open_ai, google_gemini, google_gemini_ai_studio,
	// mistral_ai, anthropic, x_ai, watsonx, deepseek, microsoft_foundry, custom_ai.
	Type ProviderType `json:"type"`
	// User’s own AI provider credentials.
	// Note: Use only if useSystemCredentials is set to `false`.
	Credentials map[string]string `json:"credentials,omitempty"`
	// AI provider configuration.
	//
	// Deprecated: this field is deprecated in the API.
	Config ProviderConfig `json:"config,omitempty"`
	// Defines whether to AI provider is enabled. Default: true.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Enables the paid service AI provider via Crowdin.
	// Note: Set to true if `credentials` is not provided. Not supported
	// for `custom_ai` type.
	UseSystemCredentials *bool `json:"useSystemCredentials,omitempty"`
}

// ProviderConfig represents the configuration of an AI provider.
type ProviderConfig struct {
	// Action rules.
	ActionRules []ActionRule `json:"actionRules"`
}

// ActionRule represents an action rule of an AI provider.
type ActionRule struct {
	// Action name.
	Action PromptAction `json:"action"`
	// Available AI provider model ids.
	AvailableAIModelIDs []string `json:"availableAiModelIds"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
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

// ProviderModel represents an AI provider model.
type ProviderModel struct {
	ID           string  `json:"id"`
	Provider     *string `json:"provider,omitempty"`
	ProviderName *string `json:"providerName,omitempty"`
	ProviderID   *int    `json:"providerId,omitempty"`
	// Context window size in tokens.
	ContextWindow *int `json:"contextWindow,omitempty"`
	// Maximum output tokens.
	MaxOutputTokens *int `json:"maxOutputTokens,omitempty"`
	// Whether the model supports streaming responses.
	SupportsStreaming *bool `json:"supportsStreaming,omitempty"`
	// Whether the model supports function/tool calling.
	SupportsFunctionCalling *bool `json:"supportsFunctionCalling,omitempty"`
	// Whether the model supports JSON mode.
	SupportsJSONMode *bool `json:"supportsJsonMode,omitempty"`
	// Whether the model supports structured output via JSON Schema.
	SupportsJSONSchema *bool `json:"supportsJsonSchema,omitempty"`
	// Whether the model accepts image input (vision).
	SupportsVision *bool `json:"supportsVision,omitempty"`
	// Whether the model can be used while AI cost limits are in effect.
	IsCompatibleWithAILimit bool `json:"isCompatibleWithAiLimit"`
}

// ProviderModelResponse defines the structure of a response when
// getting an AI provider model.
type ProviderModelResponse struct {
	Data *ProviderModel `json:"data"`
}

// ProviderModelsListResponse defines the structure of a response when
// getting a list of AI provider models.
type ProviderModelsListResponse struct {
	Data []*ProviderModelResponse `json:"data"`
}

// ProxyChatCompletion represents an AI proxy chat completion.
type ProxyChatCompletion struct{}

// ProxyChatCompletionResponse defines the structure of a response when
// getting an AI proxy chat completion.
type ProxyChatCompletionResponse struct {
	Data *ProxyChatCompletion `json:"data"`
}

// CreateProxyChatCompletionRequest defines the structure of a request
// to create an AI proxy chat completion.
type CreateProxyChatCompletionRequest struct {
	// ID of the model to use (required for Google Gemini providers).
	Model string `json:"model,omitempty"`
	// ID of the model to use.
	//
	// Deprecated: the API expects the `model` field; use Model instead.
	ModelID string `json:"modelId,omitempty"`
	// Tokens will be sent as data-only server-sent events as they become available,
	// with the stream terminated by a data: [DONE] message.
	Stream *bool `json:"stream,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *CreateProxyChatCompletionRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

// PromptCloneRequest defines the structure of a request to clone an AI prompt.
type PromptCloneRequest struct {
	// Name of the cloned AI prompt.
	Name string `json:"name,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PromptCloneRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

type (
	// PromptCompletion represents an AI prompt completion job.
	PromptCompletion struct {
		Identifier string                      `json:"identifier"`
		Status     string                      `json:"status"`
		Progress   int                         `json:"progress"`
		Attributes *PromptCompletionAttributes `json:"attributes"`
		CreatedAt  string                      `json:"createdAt"`
		UpdatedAt  string                      `json:"updatedAt"`
		StartedAt  string                      `json:"startedAt"`
		FinishedAt string                      `json:"finishedAt"`
	}

	// PromptCompletionAttributes represents the attributes of an AI prompt completion job.
	PromptCompletionAttributes struct {
		AIPromptID int `json:"aiPromptId"`
	}
)

// PromptCompletionResponse defines the structure of a response when
// getting an AI prompt completion job.
type PromptCompletionResponse struct {
	Data *PromptCompletion `json:"data"`
}

type (
	// PromptCompletionRequest defines the structure of a request to generate
	// an AI prompt completion.
	PromptCompletionRequest struct {
		// AI prompt context resources.
		Resources *PromptCompletionResources `json:"resources"`
		// List of AI tools.
		Tools []*PromptCompletionTool `json:"tools,omitempty"`
		// Controls which (if any) tool is called by the model.
		// Can be a string (e.g. "auto", "none") or an object.
		ToolChoice any `json:"tool_choice,omitempty"`
	}

	// PromptCompletionResources represents the context resources of an
	// AI prompt completion.
	PromptCompletionResources struct {
		// Project identifier.
		ProjectID int `json:"projectId"`
		// Source language identifier. If not specified, the source language
		// of the project will be used.
		SourceLanguageID string `json:"sourceLanguageId,omitempty"`
		// Target language identifier.
		TargetLanguageID string `json:"targetLanguageId,omitempty"`
		// String identifiers (up to 500).
		// Note: Must be used together with `targetLanguageId`.
		StringIDs []int `json:"stringIds,omitempty"`
		// Values to override the prompt placeholders with.
		OverridePromptValues *PromptCompletionOverridePromptValues `json:"overridePromptValues,omitempty"`
		// Custom instruction (prompts with a custom action).
		CustomInstruction string `json:"customInstruction,omitempty"`
	}

	// PromptCompletionOverridePromptValues represents the values used to
	// override the prompt placeholders. Available keys depend on the prompt action.
	PromptCompletionOverridePromptValues struct {
		SourceLanguage   string `json:"sourceLanguage,omitempty"`
		TargetLanguage   string `json:"targetLanguage,omitempty"`
		Strings          string `json:"strings,omitempty"`
		TranslationUnits string `json:"translationUnits,omitempty"`
		// Alignment pairs (Enterprise only, alignment action).
		AlignmentPairs     string `json:"alignmentPairs,omitempty"`
		TM                 string `json:"tm,omitempty"`
		Terms              string `json:"terms,omitempty"`
		FileName           string `json:"fileName,omitempty"`
		FileContext        string `json:"fileContext,omitempty"`
		FileContent        string `json:"fileContent,omitempty"`
		FilteredStrings    string `json:"filteredStrings,omitempty"`
		SiblingsStrings    string `json:"siblingsStrings,omitempty"`
		PluralForms        string `json:"pluralForms,omitempty"`
		ProjectName        string `json:"projectName,omitempty"`
		ProjectDescription string `json:"projectDescription,omitempty"`
		// Organization name (Enterprise only).
		OrganizationName string `json:"organizationName,omitempty"`
		// Organization description (Enterprise only).
		OrganizationDescription string `json:"organizationDescription,omitempty"`
	}

	// PromptCompletionTool represents an AI tool of a prompt completion request.
	PromptCompletionTool struct {
		Tool *AITool `json:"tool"`
	}

	// AITool represents an AI tool.
	AITool struct {
		// The type of the tool. Currently, only `function` is supported.
		Type string `json:"type"`
		// The function definition.
		Function *AIToolFunction `json:"function"`
	}

	// AIToolFunction represents a function that can be called by the AI.
	AIToolFunction struct {
		// A description of what the function does.
		Description string `json:"description,omitempty"`
		// The name of the function to be called.
		Name string `json:"name"`
		// The parameters the function accepts, described as a JSON Schema object.
		Parameters map[string]any `json:"parameters,omitempty"`
	}
)

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *PromptCompletionRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Resources == nil {
		return errors.New("resources is required")
	}
	if r.Resources.ProjectID == 0 {
		return errors.New("resources.projectId is required")
	}
	if len(r.Resources.StringIDs) > 0 && r.Resources.TargetLanguageID == "" {
		return errors.New("resources.targetLanguageId is required when stringIds is set")
	}
	for _, t := range r.Tools {
		if t == nil || t.Tool == nil || t.Tool.Function == nil || t.Tool.Function.Name == "" {
			return errors.New("tools: tool.function.name is required")
		}
		if t.Tool.Type == "" {
			return errors.New("tools: tool.type is required")
		}
	}

	return nil
}

type (
	// AIFileTranslation represents an AI file translation job.
	AIFileTranslation struct {
		Identifier string                       `json:"identifier"`
		Status     string                       `json:"status"`
		Progress   int                          `json:"progress"`
		Attributes *AIFileTranslationAttributes `json:"attributes"`
		CreatedAt  string                       `json:"createdAt"`
		UpdatedAt  *string                      `json:"updatedAt,omitempty"`
		StartedAt  *string                      `json:"startedAt,omitempty"`
		FinishedAt *string                      `json:"finishedAt,omitempty"`
	}

	// AIFileTranslationAttributes represents the attributes of an AI file translation job.
	AIFileTranslationAttributes struct {
		// Current stage (start, import, translate, export, done).
		Stage        string                  `json:"stage"`
		Error        *AIFileTranslationError `json:"error,omitempty"`
		DownloadName *string                 `json:"downloadName,omitempty"`
		// Source language code. If not passed in the request, it is resolved
		// after the file is parsed.
		SourceLanguageID *string `json:"sourceLanguageId,omitempty"`
		TargetLanguageID string  `json:"targetLanguageId"`
		OriginalFileName string  `json:"originalFileName"`
		// Detected base file type.
		DetectedType *string `json:"detectedType,omitempty"`
		// Parser version for the detected file type.
		ParserVersion *int `json:"parserVersion,omitempty"`
	}

	// AIFileTranslationError represents an error of an AI file translation job.
	AIFileTranslationError struct {
		Stage   string `json:"stage"`
		Message string `json:"message"`
	}
)

// AIFileTranslationResponse defines the structure of a response when
// getting an AI file translation job.
type AIFileTranslationResponse struct {
	Data *AIFileTranslation `json:"data"`
}

// AIFileTranslationRequest defines the structure of a request to
// translate a file with AI.
type AIFileTranslationRequest struct {
	// Storage identifier of the file to translate.
	StorageID int `json:"storageId"`
	// Source language identifier. If not specified, auto-detection is used.
	SourceLanguageID string `json:"sourceLanguageId,omitempty"`
	// Target language identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// File type. Default: auto.
	Type string `json:"type,omitempty"`
	// Parser version. Note: Must be used together with `type`.
	ParserVersion int `json:"parserVersion,omitempty"`
	// Translation memory identifiers.
	TMIDs []int `json:"tmIds,omitempty"`
	// Glossary identifiers.
	GlossaryIDs []int `json:"glossaryIds,omitempty"`
	// Style guide identifiers.
	StyleGuideIDs []int `json:"styleGuideIds,omitempty"`
	// Pre-translation prompt identifier.
	// Note: Can't be used with `aiProviderId` or `aiModelId`.
	AIPromptID int `json:"aiPromptId,omitempty"`
	// AI provider identifier. Note: Must be used together with `aiModelId`.
	AIProviderID int `json:"aiProviderId,omitempty"`
	// AI model identifier. Note: Must be used together with `aiProviderId`.
	AIModelID string `json:"aiModelId,omitempty"`
	// Custom instructions for translation.
	Instructions []string `json:"instructions,omitempty"`
	// Storage identifiers of images to pass to AI as attachments (max 10).
	AttachmentIDs []int `json:"attachmentIds,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AIFileTranslationRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.StorageID == 0 {
		return errors.New("storageId is required")
	}
	if r.TargetLanguageID == "" {
		return errors.New("targetLanguageId is required")
	}
	if r.ParserVersion > 0 && r.Type == "" {
		return errors.New("type is required when parserVersion is set")
	}
	if len(r.AttachmentIDs) > maxAIAttachments {
		return fmt.Errorf("attachmentIds cannot contain more than %d items", maxAIAttachments)
	}

	return validateAIPromptOrModel(r.AIPromptID, r.AIProviderID, r.AIModelID, false)
}

// maxAIAttachments is the maximum number of attachments per AI request.
const maxAIAttachments = 10

// maxAITranslateStrings is the maximum number of strings per AI translate request.
const maxAITranslateStrings = 500

// validateAIPromptOrModel validates the mutually exclusive `aiPromptId`
// and `aiProviderId` + `aiModelId` parameters.
func validateAIPromptOrModel(promptID, providerID int, modelID string, required bool) error {
	if promptID > 0 && (providerID > 0 || modelID != "") {
		return errors.New("aiPromptId can't be used with aiProviderId or aiModelId")
	}
	if (providerID > 0) != (modelID != "") {
		return errors.New("aiProviderId and aiModelId must be used together")
	}
	if required && promptID == 0 && providerID == 0 {
		return errors.New("aiPromptId or aiProviderId with aiModelId is required")
	}

	return nil
}

// AITranslateStrings represents the result of the AI strings translation.
type AITranslateStrings struct {
	SourceLanguageID string   `json:"sourceLanguageId"`
	TargetLanguageID string   `json:"targetLanguageId"`
	Translations     []string `json:"translations"`
}

// AITranslateStringsResponse defines the structure of a response when
// translating strings with AI.
type AITranslateStringsResponse struct {
	Data *AITranslateStrings `json:"data"`
}

// AITranslateStringsRequest defines the structure of a request to
// translate strings with AI.
type AITranslateStringsRequest struct {
	// Strings to translate (max 500 strings, each max 10000 characters).
	Strings []string `json:"strings"`
	// Source language identifier. If not specified, auto-detection is used.
	SourceLanguageID string `json:"sourceLanguageId,omitempty"`
	// Target language identifier.
	TargetLanguageID string `json:"targetLanguageId"`
	// Translation memory identifiers.
	TMIDs []int `json:"tmIds,omitempty"`
	// Glossary identifiers.
	GlossaryIDs []int `json:"glossaryIds,omitempty"`
	// Style guide identifiers.
	StyleGuideIDs []int `json:"styleGuideIds,omitempty"`
	// Custom instructions for translation.
	Instructions []string `json:"instructions,omitempty"`
	// Storage identifiers of images to pass to AI as attachments (max 10).
	AttachmentIDs []int `json:"attachmentIds,omitempty"`
	// Pre-translation prompt identifier.
	// Note: Required if `aiProviderId` and `aiModelId` are not provided.
	AIPromptID int `json:"aiPromptId,omitempty"`
	// AI provider identifier. Note: Must be used together with `aiModelId`.
	AIProviderID int `json:"aiProviderId,omitempty"`
	// AI model identifier. Note: Must be used together with `aiProviderId`.
	AIModelID string `json:"aiModelId,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AITranslateStringsRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if len(r.Strings) == 0 {
		return errors.New("strings is required")
	}
	if len(r.Strings) > maxAITranslateStrings {
		return fmt.Errorf("strings cannot contain more than %d items", maxAITranslateStrings)
	}
	if r.TargetLanguageID == "" {
		return errors.New("targetLanguageId is required")
	}
	if len(r.AttachmentIDs) > maxAIAttachments {
		return fmt.Errorf("attachmentIds cannot contain more than %d items", maxAIAttachments)
	}

	return validateAIPromptOrModel(r.AIPromptID, r.AIProviderID, r.AIModelID, true)
}

type (
	// SupportedModel represents a supported AI provider model.
	SupportedModel struct {
		// ID of the configured AI provider.
		ProviderID *int `json:"providerId,omitempty"`
		// Type of the AI provider (e.g., open_ai).
		ProviderType string `json:"providerType"`
		// Human-readable name of the AI provider.
		ProviderName string `json:"providerName"`
		// Model identifier.
		ID string `json:"id"`
		// Human-readable model name.
		DisplayName string `json:"displayName"`
		// Indicates if the model supports reasoning (thinking).
		SupportReasoning bool `json:"supportReasoning"`
		// Intelligence level.
		Intelligence int `json:"intelligence"`
		// Speed level.
		Speed int `json:"speed"`
		// Pricing per 1M tokens.
		Price *SupportedModelPrice `json:"price,omitempty"`
		// Supported input/output modalities.
		Modalities *SupportedModelModalities `json:"modalities,omitempty"`
		// Context window size.
		ContextWindow int `json:"contextWindow"`
		// Maximum output tokens.
		MaxOutputTokens int `json:"maxOutputTokens"`
		// Knowledge cutoff date.
		KnowledgeCutoff *string `json:"knowledgeCutoff,omitempty"`
		// Release date.
		ReleaseDate *string `json:"releaseDate,omitempty"`
		// Supported features.
		Features *SupportedModelFeatures `json:"features,omitempty"`
	}

	// SupportedModelPrice represents the pricing of a supported model per 1M tokens.
	SupportedModelPrice struct {
		Input  float64 `json:"input"`
		Output float64 `json:"output"`
	}

	// SupportedModelModalities represents the input/output modalities of a supported model.
	SupportedModelModalities struct {
		Input  *SupportedModelModality `json:"input,omitempty"`
		Output *SupportedModelModality `json:"output,omitempty"`
	}

	// SupportedModelModality represents a set of supported modalities.
	SupportedModelModality struct {
		Text  bool `json:"text"`
		Image bool `json:"image"`
		Audio bool `json:"audio"`
	}

	// SupportedModelFeatures represents the features of a supported model.
	SupportedModelFeatures struct {
		Streaming        bool `json:"streaming"`
		StructuredOutput bool `json:"structuredOutput"`
		FunctionCalling  bool `json:"functionCalling"`
	}
)

// SupportedModelResponse defines the structure of a response when
// getting a supported AI provider model.
type SupportedModelResponse struct {
	Data *SupportedModel `json:"data"`
}

// SupportedModelsListResponse defines the structure of a response when
// getting a list of supported AI provider models.
type SupportedModelsListResponse struct {
	Data []*SupportedModelResponse `json:"data"`
}

// SupportedModelsListOptions specifies the optional parameters to the
// AIService.ListSupportedProviderModels method.
type SupportedModelsListOptions struct {
	// Filter by provider type.
	ProviderType ProviderType `json:"providerType,omitempty"`
	// Filter by enabled providers.
	Enabled *bool `json:"enabled,omitempty"`
	// Sort the collection by the specified field. Example: orderBy=id desc.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of SupportedModelsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *SupportedModelsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.ProviderType != "" {
		v.Add("providerType", string(o.ProviderType))
	}
	if o.Enabled != nil {
		v.Add("enabled", strconv.FormatBool(*o.Enabled))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}

// AIReportType represents the type of an AI report.
type AIReportType string

const (
	// AIReportUsageRawData is a report with one record per AI call.
	AIReportUsageRawData AIReportType = "tokens-usage-raw-data"
	// AIReportCostsByUsers is a report with AI spend per user.
	AIReportCostsByUsers AIReportType = "costs-by-users"
)

type (
	// AIReport represents an AI report generation job.
	AIReport struct {
		Identifier string              `json:"identifier"`
		Status     string              `json:"status"`
		Progress   int                 `json:"progress"`
		Attributes *AIReportAttributes `json:"attributes"`
		CreatedAt  string              `json:"createdAt"`
		UpdatedAt  string              `json:"updatedAt"`
		StartedAt  string              `json:"startedAt"`
		FinishedAt string              `json:"finishedAt"`
		ETA        string              `json:"eta"`
	}

	// AIReportAttributes represents the attributes of an AI report.
	AIReportAttributes struct {
		Format     string         `json:"format"`
		ReportType string         `json:"reportType"`
		Schema     map[string]any `json:"schema"`
	}
)

// AIReportResponse defines the structure of a response when
// getting an AI report.
type AIReportResponse struct {
	Data *AIReport `json:"data"`
}

type (
	// AIReportGenerateRequest defines the structure of a request to generate an AI report.
	AIReportGenerateRequest struct {
		// Report type. Enum: tokens-usage-raw-data, costs-by-users.
		Type AIReportType `json:"type"`
		// Report schema.
		Schema *AIReportSchema `json:"schema"`
	}

	// AIReportSchema represents the schema of an AI report.
	AIReportSchema struct {
		// Report date from in UTC, ISO 8601.
		DateFrom string `json:"dateFrom"`
		// Report date to in UTC, ISO 8601.
		DateTo string `json:"dateTo"`
		// Report file format. Enum: json, csv. Default: json.
		Format string `json:"format,omitempty"`
		// Count only usage from these projects.
		ProjectIDs []int `json:"projectIds,omitempty"`
		// Count only usage made through these AI prompts.
		PromptIDs []int `json:"promptIds,omitempty"`
		// Count only usage by these users.
		UserIDs []int `json:"userIds,omitempty"`
	}
)

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AIReportGenerateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	switch r.Type {
	case AIReportUsageRawData, AIReportCostsByUsers: // valid
	default:
		return fmt.Errorf("invalid type: %q, must be one of %s, %s",
			r.Type, AIReportUsageRawData, AIReportCostsByUsers)
	}

	if r.Schema == nil {
		return errors.New("schema is required")
	}
	if r.Schema.DateFrom == "" {
		return errors.New("schema.dateFrom is required")
	}
	if r.Schema.DateTo == "" {
		return errors.New("schema.dateTo is required")
	}

	return nil
}

// AIRequestLog represents an AI request log entry.
type AIRequestLog struct {
	ID                int      `json:"id"`
	RequestID         string   `json:"requestId"`
	CreatedAt         string   `json:"createdAt"`
	Status            string   `json:"status"`
	HTTPStatus        *int     `json:"httpStatus,omitempty"`
	Model             string   `json:"model"`
	SourceAction      string   `json:"sourceAction"`
	PromptAction      *string  `json:"promptAction,omitempty"`
	SystemCredentials bool     `json:"systemCredentials"`
	IsAutoTriggered   bool     `json:"isAutoTriggered"`
	DurationMs        *int     `json:"durationMs,omitempty"`
	InputTokens       *int     `json:"inputTokens,omitempty"`
	OutputTokens      *int     `json:"outputTokens,omitempty"`
	TotalCost         *float64 `json:"totalCost,omitempty"`
	UserID            *int     `json:"userId,omitempty"`
	ProjectID         *int     `json:"projectId,omitempty"`
	PromptID          *int     `json:"promptId,omitempty"`
	AIProviderID      int      `json:"aiProviderId"`
	TokenName         *string  `json:"tokenName,omitempty"`
	OAuthClientID     *string  `json:"oauthClientId,omitempty"`
	OAuthClientName   *string  `json:"oauthClientName,omitempty"`
	IP                *string  `json:"ip,omitempty"`
	UserAgent         *string  `json:"userAgent,omitempty"`
	Error             *string  `json:"error,omitempty"`
}

// AIRequestLogResponse defines the structure of a response when
// getting an AI request log entry.
type AIRequestLogResponse struct {
	Data *AIRequestLog `json:"data"`
}

// AIRequestLogsListResponse defines the structure of a response when
// getting a list of AI request logs.
type AIRequestLogsListResponse struct {
	Data []*AIRequestLogResponse `json:"data"`
}

// AIRequestLogsListOptions specifies the optional parameters to the
// AIService.ListRequestLogs method.
type AIRequestLogsListOptions struct {
	// Filter by request identifier.
	RequestID string `json:"requestId,omitempty"`
	// Filter by project.
	ProjectID int `json:"projectId,omitempty"`
	// Filter by the user attributed to the AI request.
	UserID int `json:"userId,omitempty"`
	// Filter by AI provider.
	AIProviderID int `json:"aiProviderId,omitempty"`
	// Filter by model name.
	Model string `json:"model,omitempty"`
	// Filter by the feature or channel that produced the request.
	// Enum: ai_gateway, ai_proxy, ai_translate_strings, ai_file_translate,
	// ai_prompt_completion, pre_translate:manual, pre_translate:workflow,
	// ai_alignment, qa_check, ai_suggestion, advisor.
	SourceAction string `json:"sourceAction,omitempty"`
	// Filter by prompt action.
	PromptAction string `json:"promptAction,omitempty"`
	// Filter by status. Enum: pending, success, error, timeout.
	Statuses []string `json:"statuses,omitempty"`
	// Filter by whether the request used system-provided credentials.
	SystemCredentials *bool `json:"systemCredentials,omitempty"`
	// Filter by whether the request was triggered automatically.
	IsAutoTriggered *bool `json:"isAutoTriggered,omitempty"`
	// Filter by the name of the personal access token.
	TokenName string `json:"tokenName,omitempty"`
	// Filter by the OAuth application client_id.
	OAuthClientID string `json:"oauthClientId,omitempty"`
	// Return logs created after this date (ISO 8601).
	CreatedAfter string `json:"createdAfter,omitempty"`
	// Return logs created before this date (ISO 8601).
	CreatedBefore string `json:"createdBefore,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of AIRequestLogsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *AIRequestLogsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if o.RequestID != "" {
		v.Add("requestId", o.RequestID)
	}
	if o.ProjectID > 0 {
		v.Add("projectId", strconv.Itoa(o.ProjectID))
	}
	if o.UserID > 0 {
		v.Add("userId", strconv.Itoa(o.UserID))
	}
	if o.AIProviderID > 0 {
		v.Add("aiProviderId", strconv.Itoa(o.AIProviderID))
	}
	if o.Model != "" {
		v.Add("model", o.Model)
	}
	if o.SourceAction != "" {
		v.Add("sourceAction", o.SourceAction)
	}
	if o.PromptAction != "" {
		v.Add("promptAction", o.PromptAction)
	}
	if len(o.Statuses) > 0 {
		v.Add("statuses", JoinSlice(o.Statuses))
	}
	if o.SystemCredentials != nil {
		v.Add("systemCredentials", strconv.FormatBool(*o.SystemCredentials))
	}
	if o.IsAutoTriggered != nil {
		v.Add("isAutoTriggered", strconv.FormatBool(*o.IsAutoTriggered))
	}
	if o.TokenName != "" {
		v.Add("tokenName", o.TokenName)
	}
	if o.OAuthClientID != "" {
		v.Add("oauthClientId", o.OAuthClientID)
	}
	if o.CreatedAfter != "" {
		v.Add("createdAfter", o.CreatedAfter)
	}
	if o.CreatedBefore != "" {
		v.Add("createdBefore", o.CreatedBefore)
	}

	return v, len(v) > 0
}

type (
	// AIRequestLogsExport represents an AI request logs export job.
	AIRequestLogsExport struct {
		Identifier string                         `json:"identifier"`
		Status     string                         `json:"status"`
		Progress   int                            `json:"progress"`
		Attributes *AIRequestLogsExportAttributes `json:"attributes"`
		CreatedAt  string                         `json:"createdAt"`
		UpdatedAt  string                         `json:"updatedAt"`
		StartedAt  *string                        `json:"startedAt,omitempty"`
		FinishedAt *string                        `json:"finishedAt,omitempty"`
		ETA        *string                        `json:"eta,omitempty"`
	}

	// AIRequestLogsExportAttributes represents the attributes of an AI request logs export.
	AIRequestLogsExportAttributes struct {
		Format  string         `json:"format"`
		Filters map[string]any `json:"filters"`
	}
)

// AIRequestLogsExportResponse defines the structure of a response when
// getting an AI request logs export.
type AIRequestLogsExportResponse struct {
	Data *AIRequestLogsExport `json:"data"`
}

// AIRequestLogsExportRequest defines the structure of a request to export
// AI request logs. All fields are optional; passing no filters exports
// every log entry.
type AIRequestLogsExportRequest struct {
	// Format of the exported file. Enum: csv. Default: csv.
	Format string `json:"format,omitempty"`
	// Filter by request identifier.
	RequestID string `json:"requestId,omitempty"`
	// Filter by project.
	ProjectID int `json:"projectId,omitempty"`
	// Filter by the user attributed to the AI request.
	UserID int `json:"userId,omitempty"`
	// Filter by AI provider.
	AIProviderID int `json:"aiProviderId,omitempty"`
	// Filter by AI model.
	Model string `json:"model,omitempty"`
	// Filter by the feature or channel that produced the request.
	SourceAction string `json:"sourceAction,omitempty"`
	// Filter by prompt action.
	PromptAction string `json:"promptAction,omitempty"`
	// Filter by statuses. Enum: pending, success, error, timeout.
	Statuses []string `json:"statuses,omitempty"`
	// Filter by the name of the personal access token.
	TokenName string `json:"tokenName,omitempty"`
	// Filter by the OAuth application client_id.
	OAuthClientID string `json:"oauthClientId,omitempty"`
	// Export only requests that used system-provided credentials.
	SystemCredentials *bool `json:"systemCredentials,omitempty"`
	// Export only requests triggered automatically.
	IsAutoTriggered *bool `json:"isAutoTriggered,omitempty"`
	// Filter by requests logged after this date (ISO 8601).
	CreatedAfter string `json:"createdAfter,omitempty"`
	// Filter by requests logged before this date (ISO 8601).
	CreatedBefore string `json:"createdBefore,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AIRequestLogsExportRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	return nil
}

type (
	// AISettings represents the AI settings.
	AISettings struct {
		PreTranslationAIPromptID   int  `json:"preTranslationAiPromptId"`
		EditorSuggestionAIPromptID int  `json:"editorSuggestionAiPromptId"`
		QACheckActionAIPromptID    *int `json:"qaCheckActionAiPromptId,omitempty"`
		ContextReviewAIPromptID    *int `json:"contextReviewAiPromptId,omitempty"`
		// Alignment action AI prompt identifier (Enterprise only).
		AlignmentActionAIPromptID *int `json:"alignmentActionAiPromptId,omitempty"`
		// Maximum total AI cost the organization can spend per day.
		DailyCostLimit *float64 `json:"dailyCostLimit,omitempty"`
		// Maximum total AI cost the organization can spend per month.
		MonthlyCostLimit *float64 `json:"monthlyCostLimit,omitempty"`
		// Default maximum AI cost a single user can spend per day.
		UserDailyCostLimit *float64 `json:"userDailyCostLimit,omitempty"`
		// Default maximum AI cost a single user can spend per month.
		UserMonthlyCostLimit *float64 `json:"userMonthlyCostLimit,omitempty"`
		// Whether any AI cost limit is currently in force.
		IsLimitingActive bool `json:"isLimitingActive"`
		// Per-user cost limit overrides.
		PerUserOverrides []*AIUserCostLimitOverride `json:"perUserOverrides"`
	}

	// AIUserCostLimitOverride represents a per-user AI cost limit override.
	AIUserCostLimitOverride struct {
		UserID int `json:"userId"`
		// Cost limit mode. Enum: custom, blocked, unlimited.
		CostLimitMode    string   `json:"costLimitMode"`
		DailyCostLimit   *float64 `json:"dailyCostLimit,omitempty"`
		MonthlyCostLimit *float64 `json:"monthlyCostLimit,omitempty"`
		CreatedAt        string   `json:"createdAt"`
		UpdatedAt        string   `json:"updatedAt"`
	}
)

// AISettingsResponse defines the structure of a response when
// getting the AI settings.
type AISettingsResponse struct {
	Data *AISettings `json:"data"`
}

// ProjectAISettings represents the AI settings of a project.
type ProjectAISettings struct {
	EditorSuggestionAIPromptID int  `json:"editorSuggestionAiPromptId"`
	QACheckActionAIPromptID    *int `json:"qaCheckActionAiPromptId,omitempty"`
	ContextReviewAIPromptID    *int `json:"contextReviewAiPromptId,omitempty"`
	// Alignment action AI prompt identifier (Enterprise only).
	AlignmentActionAIPromptID *int `json:"alignmentActionAiPromptId,omitempty"`
}

// ProjectAISettingsResponse defines the structure of a response when
// getting the AI settings of a project.
type ProjectAISettingsResponse struct {
	Data *ProjectAISettings `json:"data"`
}

// AISnippet represents an AI snippet.
type AISnippet struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Placeholder string `json:"placeholder"`
	Value       string `json:"value"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// AISnippetResponse defines the structure of a response when
// getting an AI snippet.
type AISnippetResponse struct {
	Data *AISnippet `json:"data"`
}

// AISnippetsListResponse defines the structure of a response when
// getting a list of AI snippets.
type AISnippetsListResponse struct {
	Data []*AISnippetResponse `json:"data"`
}

// AISnippetAddRequest defines the structure of a request to add an AI snippet.
type AISnippetAddRequest struct {
	// Unique description, between 3 and 255 characters.
	Description string `json:"description"`
	// Unique placeholder, must start with `%custom:` and end with `%`.
	Placeholder string `json:"placeholder"`
	// The text that will be used in the prompt (max 4000 characters).
	Value string `json:"value"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AISnippetAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Description == "" {
		return errors.New("description is required")
	}
	if r.Placeholder == "" {
		return errors.New("placeholder is required")
	}
	if r.Value == "" {
		return errors.New("value is required")
	}

	return nil
}

// AIUsageMember represents the AI usage and limits of an organization member.
type AIUsageMember struct {
	// Organization member the usage and limits belong to.
	User *ShortUser `json:"user"`
	// Effective daily AI cost limit in USD. `nil` means no limit.
	DailyCostLimit *float64 `json:"dailyCostLimit,omitempty"`
	// AI cost spent in the current day, in USD.
	DailyCostSpent float64 `json:"dailyCostSpent"`
	// Date and time when the daily cost limit resets.
	DailyResetAt string `json:"dailyResetAt"`
	// Effective monthly AI cost limit in USD. `nil` means no limit.
	MonthlyCostLimit *float64 `json:"monthlyCostLimit,omitempty"`
	// AI cost spent in the current month, in USD.
	MonthlyCostSpent float64 `json:"monthlyCostSpent"`
	// Date and time when the monthly cost limit resets.
	MonthlyResetAt string `json:"monthlyResetAt"`
}

// AIUsageMemberResponse defines the structure of a response when
// getting the AI usage of a member.
type AIUsageMemberResponse struct {
	Data *AIUsageMember `json:"data"`
}

// AIUsageMembersListResponse defines the structure of a response when
// getting a list of AI usage members.
type AIUsageMembersListResponse struct {
	Data []*AIUsageMemberResponse `json:"data"`
}

// AIUsageMembersListOptions specifies the optional parameters to the
// AIService.ListUsageMembers method.
type AIUsageMembersListOptions struct {
	// Filter by user identifiers.
	UserIDs []int `json:"userIds,omitempty"`
	// Sort the collection by the specified field.
	OrderBy string `json:"orderBy,omitempty"`

	ListOptions
}

// Values returns the url.Values encoding of AIUsageMembersListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *AIUsageMembersListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()

	if len(o.UserIDs) > 0 {
		v.Add("userIds", JoinSlice(o.UserIDs))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}
