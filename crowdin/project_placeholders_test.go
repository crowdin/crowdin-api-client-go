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

func TestProjectPlaceholdersService_List(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name:          "with options",
			opts:          &model.ListOptions{Offset: 1, Limit: 10},
			expectedQuery: "?limit=10&offset=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/placeholders"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 1,
								"customPlaceholderId": 3,
								"type": "high",
								"index": 1,
								"isBlocking": false,
								"formats": ["docbook", "adoc"]
							}
						}
					],
					"pagination": {
						"offset": 1,
						"limit": 10
					}
				}`)
			})

			placeholders, resp, err := client.ProjectPlaceholders.List(context.Background(), 1, tt.opts)
			require.NoError(t, err)

			expected := []*model.ProjectPlaceholder{
				{
					ID:                  1,
					CustomPlaceholderID: 3,
					Type:                ToPtr("high"),
					Index:               ToPtr(1),
					IsBlocking:          ToPtr(false),
					Formats:             []string{"docbook", "adoc"},
				},
			}
			assert.Equal(t, expected, placeholders)

			assert.Equal(t, 1, resp.Pagination.Offset)
			assert.Equal(t, 10, resp.Pagination.Limit)
		})
	}
}

func TestProjectPlaceholdersService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/placeholders", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.ProjectPlaceholders.List(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestProjectPlaceholdersService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/placeholders/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"customPlaceholderId": 3,
				"type": null,
				"index": null,
				"isBlocking": null,
				"formats": []
			}
		}`)
	})

	placeholder, resp, err := client.ProjectPlaceholders.Get(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectPlaceholder{
		ID:                  2,
		CustomPlaceholderID: 3,
		Formats:             []string{},
	}
	assert.Equal(t, expected, placeholder)
}

func TestProjectPlaceholdersService_Add(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/placeholders"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"customPlaceholderId": 3,
			"type": "high",
			"index": 1,
			"isBlocking": false,
			"formats": ["docbook", "adoc"]
		}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"customPlaceholderId": 3,
				"type": "high",
				"index": 1,
				"isBlocking": false,
				"formats": ["docbook", "adoc"]
			}
		}`)
	})

	req := &model.ProjectPlaceholderAddRequest{
		CustomPlaceholderID: 3,
		Type:                model.ProjectPlaceholderTypeHigh,
		Index:               1,
		IsBlocking:          ToPtr(false),
		Formats:             []string{"docbook", "adoc"},
	}
	placeholder, resp, err := client.ProjectPlaceholders.Add(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := &model.ProjectPlaceholder{
		ID:                  1,
		CustomPlaceholderID: 3,
		Type:                ToPtr("high"),
		Index:               ToPtr(1),
		IsBlocking:          ToPtr(false),
		Formats:             []string{"docbook", "adoc"},
	}
	assert.Equal(t, expected, placeholder)
}

func TestProjectPlaceholdersService_Add_requiredFields(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/placeholders"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJSONBody(t, r, `{"customPlaceholderId": 3}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"data": {"id": 1, "customPlaceholderId": 3}}`)
	})

	req := &model.ProjectPlaceholderAddRequest{CustomPlaceholderID: 3}
	placeholder, _, err := client.ProjectPlaceholders.Add(context.Background(), 1, req)
	require.NoError(t, err)
	assert.Equal(t, &model.ProjectPlaceholder{ID: 1, CustomPlaceholderID: 3}, placeholder)
}

func TestProjectPlaceholdersService_Add_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.ProjectPlaceholders.Add(context.Background(), 1, &model.ProjectPlaceholderAddRequest{})
	require.EqualError(t, err, "customPlaceholderId is required")
}

func TestProjectPlaceholdersService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/placeholders/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/isBlocking","value":true}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"customPlaceholderId": 3,
				"type": "low",
				"index": 0,
				"isBlocking": true,
				"formats": []
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    model.OpReplace,
			Path:  "/isBlocking",
			Value: true,
		},
	}
	placeholder, resp, err := client.ProjectPlaceholders.Edit(context.Background(), 1, 2, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ProjectPlaceholder{
		ID:                  2,
		CustomPlaceholderID: 3,
		Type:                ToPtr("low"),
		Index:               ToPtr(0),
		IsBlocking:          ToPtr(true),
		Formats:             []string{},
	}
	assert.Equal(t, expected, placeholder)
}

func TestProjectPlaceholdersService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/placeholders/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.ProjectPlaceholders.Delete(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestProjectPlaceholdersService_ListSystemPlaceholders(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.ListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "",
		},
		{
			name:          "with options",
			opts:          &model.ListOptions{Offset: 0, Limit: 50},
			expectedQuery: "?limit=50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/system-placeholders"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": "bracesSingle",
								"isEnabled": true,
								"label": "Single braces",
								"examples": ["{0}"],
								"description": "Java MessageFormat and .NET"
							}
						}
					],
					"pagination": {
						"offset": 0,
						"limit": 50
					}
				}`)
			})

			placeholders, resp, err := client.ProjectPlaceholders.ListSystemPlaceholders(context.Background(), 1, tt.opts)
			require.NoError(t, err)

			expected := []*model.ProjectSystemPlaceholder{
				{
					ID:          "bracesSingle",
					IsEnabled:   true,
					Label:       "Single braces",
					Examples:    []string{"{0}"},
					Description: "Java MessageFormat and .NET",
				},
			}
			assert.Equal(t, expected, placeholders)

			assert.Equal(t, 0, resp.Pagination.Offset)
			assert.Equal(t, 50, resp.Pagination.Limit)
		})
	}
}

func TestProjectPlaceholdersService_ListSystemPlaceholders_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/system-placeholders", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.ProjectPlaceholders.ListSystemPlaceholders(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestProjectPlaceholdersService_SystemPlaceholderBatchOperations(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/system-placeholders"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, path)
		testJSONBodyAny(t, r, `[
			{"op": "replace", "path": "/bracesSingle/isEnabled", "value": false},
			{"op": "replace", "path": "/twig/isEnabled", "value": true}
		]`)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": "twig",
						"isEnabled": true,
						"label": "Twig",
						"examples": ["{{ name }}"],
						"description": "Twig templates"
					}
				},
				{
					"data": {
						"id": "bracesSingle",
						"isEnabled": false,
						"label": "Single braces",
						"examples": ["{0}"],
						"description": "Java MessageFormat and .NET"
					}
				}
			]
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    model.OpReplace,
			Path:  "/bracesSingle/isEnabled",
			Value: false,
		},
		{
			Op:    model.OpReplace,
			Path:  "/twig/isEnabled",
			Value: true,
		},
	}
	placeholders, resp, err := client.ProjectPlaceholders.SystemPlaceholderBatchOperations(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := []*model.ProjectSystemPlaceholder{
		{
			ID:          "twig",
			IsEnabled:   true,
			Label:       "Twig",
			Examples:    []string{"{{ name }}"},
			Description: "Twig templates",
		},
		{
			ID:          "bracesSingle",
			IsEnabled:   false,
			Label:       "Single braces",
			Examples:    []string{"{0}"},
			Description: "Java MessageFormat and .NET",
		},
	}
	assert.Equal(t, expected, placeholders)
}

func TestProjectPlaceholdersService_SystemPlaceholderBatchOperations_emptyRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	res, _, err := client.ProjectPlaceholders.SystemPlaceholderBatchOperations(context.Background(), 1, nil)
	require.EqualError(t, err, "body cannot be empty or nil")
	assert.Nil(t, res)
}
