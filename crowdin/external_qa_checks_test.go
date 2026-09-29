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

func TestExternalQAChecksService_List(t *testing.T) {
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
			opts:          &model.ListOptions{Offset: 2, Limit: 20},
			expectedQuery: "?limit=20&offset=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/external-qa-checks"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 1,
								"name": "Test QA check",
								"description": "Do some QA checks",
								"config": {"identifier": "app-identifier"},
								"createdAt": "2024-09-19T14:14:00+00:00",
								"updatedAt": null
							}
						}
					],
					"pagination": {
						"offset": 2,
						"limit": 20
					}
				}`)
			})

			checks, resp, err := client.ExternalQAChecks.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.ExternalQACheck{
				{
					ID:          1,
					Name:        "Test QA check",
					Description: ToPtr("Do some QA checks"),
					Config:      map[string]any{"identifier": "app-identifier"},
					CreatedAt:   "2024-09-19T14:14:00+00:00",
				},
			}
			assert.Equal(t, expected, checks)

			assert.Equal(t, 2, resp.Pagination.Offset)
			assert.Equal(t, 20, resp.Pagination.Limit)
		})
	}
}

func TestExternalQAChecksService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/external-qa-checks", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.ExternalQAChecks.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestExternalQAChecksService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/external-qa-checks/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 1,
				"name": "Test QA check",
				"description": null,
				"config": {},
				"createdAt": "2024-09-19T14:14:00+00:00",
				"updatedAt": "2024-09-20T14:14:00+00:00"
			}
		}`)
	})

	check, resp, err := client.ExternalQAChecks.Get(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.ExternalQACheck{
		ID:        1,
		Name:      "Test QA check",
		Config:    map[string]any{},
		CreatedAt: "2024-09-19T14:14:00+00:00",
		UpdatedAt: ToPtr("2024-09-20T14:14:00+00:00"),
	}
	assert.Equal(t, expected, check)
}
