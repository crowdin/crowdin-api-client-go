package crowdin

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAIService_GetPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/prompts/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate",
				"aiProviderId": 2,
				"aiModelId": "gpt-3.5-turbo-instruct",
				"isEnabled": true,
				"enabledProjectIds": [1],
				"config": {
					"mode": "basic",
					"companyDescription": "string",
					"projectDescription": "string",
					"audienceDescription": "string",
					"otherLanguageTranslations": {
						"isEnabled": true,
						"languageIds": ["uk"]
					},
					"glossaryTerms": true,
					"tmSuggestions": true,
					"fileContent": true,
					"fileContext": true,
					"publicProjectDescription": true
				},
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	prompt, resp, err := client.AI.GetPrompt(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:                2,
		Name:              "Pre-translate prompt",
		Action:            "pre_translate",
		AIProviderID:      2,
		AIModelID:         "gpt-3.5-turbo-instruct",
		IsEnabled:         true,
		EnabledProjectIDs: []int{1},
		Config: model.PromptConfig{
			Mode:                model.ModeBasic,
			CompanyDescription:  ToPtr("string"),
			ProjectDescription:  ToPtr("string"),
			AudienceDescription: ToPtr("string"),
			OtherLanguageTranslations: &model.OtherLanguageTranslations{
				IsEnabled:   ToPtr(true),
				LanguageIDs: []string{"uk"},
			},
			GlossaryTerms:            ToPtr(true),
			TMSuggestions:            ToPtr(true),
			FileContent:              ToPtr(true),
			FileContext:              ToPtr(true),
			PublicProjectDescription: ToPtr(true),
		},
		CreatedAt: "2023-09-20T11:11:05+00:00",
		UpdatedAt: "2023-09-20T12:22:20+00:00",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_ListPrompts(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.AIPromtsListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.AIPromtsListOptions{},
		},
		{
			name: "with options",
			opts: &model.AIPromtsListOptions{
				ProjectID:   1,
				Action:      model.ActionAssist,
				ListOptions: model.ListOptions{Offset: 1, Limit: 25},
			},
			expectedQuery: "?action=assist&limit=25&offset=1&projectId=1",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for userID, tt := range tests {
		userID++
		path := fmt.Sprintf("/api/v2/users/%d/ai/prompts", userID)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path+tt.expectedQuery)

			fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": 2
						}
					},
					{
						"data": {
							"id": 4
						}
					}
				],
				"pagination": {
					"offset": 10,
					"limit": 25
				}
			}`)
		})

		prompts, resp, err := client.AI.ListPrompts(context.Background(), userID, tt.opts)
		require.NoError(t, err)

		expected := []*model.Prompt{{ID: 2}, {ID: 4}}
		assert.Equal(t, expected, prompts)

		assert.Equal(t, 10, resp.Pagination.Offset)
		assert.Equal(t, 25, resp.Pagination.Limit)
	}
}

func TestAIService_ListPrompts_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/prompts", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	prompts, _, err := client.AI.ListPrompts(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, prompts)
}

func TestAIService_AddPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"name":"Pre-translate prompt",
			"action":"pre_translate",
			"aiProviderId":1,
			"aiModelId":"gpt-3.5-turbo-instruct",
			"config":{
				"mode":"basic"
			}
		}`)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate"
			}
		}`)
	})

	req := &model.PromptAddRequest{
		Name:         "Pre-translate prompt",
		Action:       "pre_translate",
		AIProviderID: 1,
		AIModelID:    "gpt-3.5-turbo-instruct",
		Config: model.PromptConfig{
			Mode: model.ModeBasic,
		},
	}
	prompt, resp, err := client.AI.AddPrompt(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:     2,
		Name:   "Pre-translate prompt",
		Action: "pre_translate",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_EditPrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/prompts/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/name","value":"Pre-translate prompt"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Pre-translate prompt",
				"action": "pre_translate"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "Pre-translate prompt",
		},
	}
	prompt, resp, err := client.AI.EditPrompt(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Prompt{
		ID:     2,
		Name:   "Pre-translate prompt",
		Action: "pre_translate",
	}
	assert.Equal(t, expected, prompt)
}

func TestAIService_DeletePrompt(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	var userID = 1
	t.Run("path with user id", func(t *testing.T) {
		const path = "/api/v2/users/1/ai/prompts/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeletePrompt(context.Background(), 2, userID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	userID = 0
	t.Run("path without user id", func(t *testing.T) {
		const path = "/api/v2/ai/prompts/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeletePrompt(context.Background(), 2, userID)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestAIService_GetProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/providers/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai",
				"credentials": {
					"apiKey": "string"
				},
				"config": {
					"actionRules": [
						{
							"action": "pre_translate",
							"availableAiModelIds": [
								"gpt-3.5-turbo-instruct"
							]
						}
					]
				},
				"isEnabled": true,
				"useSystemCredentials": false,
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	provider, resp, err := client.AI.GetProvider(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"apiKey": "string",
		},
		Config: model.ProviderConfig{
			ActionRules: []model.ActionRule{
				{
					Action:              model.ActionPreTranslate,
					AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
				},
			},
		},
		IsEnabled:            true,
		UseSystemCredentials: false,
		CreatedAt:            "2023-09-20T11:11:05+00:00",
		UpdatedAt:            "2023-09-20T12:22:20+00:00",
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_ListProviders(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.ListOptions{},
		},
		{
			name:          "with options",
			opts:          &model.ListOptions{Offset: 1, Limit: 25},
			expectedQuery: "?limit=25&offset=1",
		},
	}

	client, mux, teardown := setupClient()
	defer teardown()

	for userID, tt := range tests {
		userID++
		path := fmt.Sprintf("/api/v2/users/%d/ai/providers", userID)
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			testURL(t, r, path+tt.expectedQuery)

			fmt.Fprint(w, `{
				"data": [
					{
						"data": {
							"id": 2
						}
					},
					{
						"data": {
							"id": 4
						}
					}
				],
				"pagination": {
					"offset": 10,
					"limit": 25
				}
			}`)
		})

		providers, resp, err := client.AI.ListProviders(context.Background(), userID, tt.opts)
		require.NoError(t, err)

		expected := []*model.Provider{{ID: 2}, {ID: 4}}
		assert.Equal(t, expected, providers)

		assert.Equal(t, 10, resp.Pagination.Offset)
		assert.Equal(t, 25, resp.Pagination.Limit)
	}
}

func TestAIService_ListProviders_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/users/1/ai/providers", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	providers, _, err := client.AI.ListProviders(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, providers)
}

func TestAIService_AddProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/providers"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"name":"OpenAI",
			"type":"open_ai",
			"credentials":{
				"resourceName":"resourceName",
				"apiKey":"apiKey123",
				"deploymentName":"deploymentName",
				"apiVersion":"v.1.0"
			},
			"config":{
				"actionRules":[
					{
						"action":"pre_translate",
						"availableAiModelIds":["gpt-3.5-turbo-instruct"]
					}
				]
			},
			"isEnabled":true,
			"useSystemCredentials":false
		}`)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai",
				"credentials": {
					"resourceName": "resourceName",
					"apiKey": "apiKey123",
					"deploymentName": "deploymentName",
					"apiVersion": "v.1.0"
				},
				"config": {
					"actionRules": [
						{
							"action": "pre_translate",
							"availableAiModelIds": [
								"gpt-3.5-turbo-instruct"
							]
						}
					]
				},
				"isEnabled": true,
				"useSystemCredentials": false,
				"createdAt": "2023-09-20T11:11:05+00:00",
				"updatedAt": "2023-09-20T12:22:20+00:00"
			}
		}`)
	})

	req := &model.ProviderAddRequest{
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"resourceName":   "resourceName",
			"apiKey":         "apiKey123",
			"deploymentName": "deploymentName",
			"apiVersion":     "v.1.0",
		},
		Config: model.ProviderConfig{
			ActionRules: []model.ActionRule{
				{
					Action:              model.ActionPreTranslate,
					AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
				},
			},
		},
		IsEnabled:            ToPtr(true),
		UseSystemCredentials: ToPtr(false),
	}
	provider, resp, err := client.AI.AddProvider(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
		Credentials: map[string]string{
			"resourceName":   "resourceName",
			"apiKey":         "apiKey123",
			"deploymentName": "deploymentName",
			"apiVersion":     "v.1.0",
		},
		Config: model.ProviderConfig{
			ActionRules: []model.ActionRule{
				{
					Action:              model.ActionPreTranslate,
					AvailableAIModelIDs: []string{"gpt-3.5-turbo-instruct"},
				},
			},
		},
		IsEnabled:            true,
		UseSystemCredentials: false,
		CreatedAt:            "2023-09-20T11:11:05+00:00",
		UpdatedAt:            "2023-09-20T12:22:20+00:00",
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_EditProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/12345/ai/providers/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/name","value":"OpenAI"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "OpenAI",
				"type": "open_ai"
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    "replace",
			Path:  "/name",
			Value: "OpenAI",
		},
	}
	provider, resp, err := client.AI.EditProvider(context.Background(), 2, 12345, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.Provider{
		ID:   2,
		Name: "OpenAI",
		Type: model.OpenAI,
	}
	assert.Equal(t, expected, provider)
}

func TestAIService_DeleteProvider(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	t.Run("path with user id", func(t *testing.T) {
		const path = "/api/v2/users/12345/ai/providers/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeleteProvider(context.Background(), 2, 12345)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("path without user id (enterprise client)", func(t *testing.T) {
		const path = "/api/v2/ai/providers/2"
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			testURL(t, r, path)
			w.WriteHeader(http.StatusNoContent)
		})

		resp, err := client.AI.DeleteProvider(context.Background(), 2, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestAIService_ListProviderModels(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/ai/providers/2/models"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": "gpt-3.5-turbo-instruct"
					}
				}
			]
		}`)
	})

	models, resp, err := client.AI.ListProviderModels(context.Background(), 2, 0)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.ProviderModel{
		{
			ID: "gpt-3.5-turbo-instruct",
		},
	}
	assert.Equal(t, expected, models)
}

func TestAIService_ListProviderModels_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/ai/providers/2/models", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	models, _, err := client.AI.ListProviderModels(context.Background(), 2, 0)
	require.Error(t, err)
	assert.Nil(t, models)
}

func TestAIService_CreateProxyChatCompletion(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/users/1/ai/providers/2/chat/completions"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, path)
		testBody(t, r, `{"modelId":"gpt-3.5-turbo-instruct","stream":false}`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"modelId": "gpt-3.5-turbo-instruct"
			}
		}`)
	})

	req := &model.CreateProxyChatCompletionRequest{
		ModelID: "gpt-3.5-turbo-instruct",
		Stream:  ToPtr(false),
	}
	completion, resp, err := client.AI.CreateProxyChatCompletion(context.Background(), 2, 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProxyChatCompletion{}
	assert.Equal(t, expected, completion)
}
