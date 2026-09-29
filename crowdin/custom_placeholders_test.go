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

func TestCustomPlaceholdersService_List(t *testing.T) {
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

			const path = "/api/v2/custom-placeholders"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 1,
								"definition": "start, then \"http\", end",
								"description": "URL validation",
								"argumentDelimiter": "\""
							}
						},
						{
							"data": {
								"id": 2,
								"description": null
							}
						}
					],
					"pagination": {
						"offset": 1,
						"limit": 10
					}
				}`)
			})

			placeholders, resp, err := client.CustomPlaceholders.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.CustomPlaceholder{
				{
					ID:                1,
					Definition:        `start, then "http", end`,
					Description:       ToPtr("URL validation"),
					ArgumentDelimiter: `"`,
				},
				{
					ID: 2,
				},
			}
			assert.Equal(t, expected, placeholders)

			assert.Equal(t, 1, resp.Pagination.Offset)
			assert.Equal(t, 10, resp.Pagination.Limit)
		})
	}
}

func TestCustomPlaceholdersService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/custom-placeholders", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.CustomPlaceholders.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestCustomPlaceholdersService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/custom-placeholders/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"definition": "start, then \"http\", end",
				"description": "URL validation",
				"argumentDelimiter": "\""
			}
		}`)
	})

	placeholder, resp, err := client.CustomPlaceholders.Get(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.CustomPlaceholder{
		ID:                1,
		Definition:        `start, then "http", end`,
		Description:       ToPtr("URL validation"),
		ArgumentDelimiter: `"`,
	}
	assert.Equal(t, expected, placeholder)
}

func TestCustomPlaceholdersService_Add(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/custom-placeholders"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"definition": "start, then \"http\", end",
			"description": "URL validation",
			"argumentDelimiter": "\""
		}`)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"definition": "start, then \"http\", end",
				"description": "URL validation",
				"argumentDelimiter": "\""
			}
		}`)
	})

	req := &model.CustomPlaceholderAddRequest{
		Definition:        `start, then "http", end`,
		Description:       "URL validation",
		ArgumentDelimiter: `"`,
	}
	placeholder, resp, err := client.CustomPlaceholders.Add(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	expected := &model.CustomPlaceholder{
		ID:                1,
		Definition:        `start, then "http", end`,
		Description:       ToPtr("URL validation"),
		ArgumentDelimiter: `"`,
	}
	assert.Equal(t, expected, placeholder)
}

func TestCustomPlaceholdersService_Add_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, _, err := client.CustomPlaceholders.Add(context.Background(), &model.CustomPlaceholderAddRequest{})
	require.EqualError(t, err, "definition is required")
}

func TestCustomPlaceholdersService_Edit(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/custom-placeholders/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testURL(t, r, path)
		testBody(t, r, `[{"op":"replace","path":"/description","value":"Link validation"}]`+"\n")

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"definition": "start, then \"http\", end",
				"description": "Link validation",
				"argumentDelimiter": "\""
			}
		}`)
	})

	req := []*model.UpdateRequest{
		{
			Op:    model.OpReplace,
			Path:  "/description",
			Value: "Link validation",
		},
	}
	placeholder, resp, err := client.CustomPlaceholders.Edit(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.CustomPlaceholder{
		ID:                1,
		Definition:        `start, then "http", end`,
		Description:       ToPtr("Link validation"),
		ArgumentDelimiter: `"`,
	}
	assert.Equal(t, expected, placeholder)
}

func TestCustomPlaceholdersService_Delete(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/custom-placeholders/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testURL(t, r, path)

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.CustomPlaceholders.Delete(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
