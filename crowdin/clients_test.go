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

func TestClientsService_List(t *testing.T) {
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
			opts:          &model.ListOptions{Offset: 10, Limit: 25},
			expectedQuery: "?limit=25&offset=10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/clients"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 52760,
								"name": "John Smith",
								"description": "John Smith Organization",
								"status": "pending",
								"webUrl": "https://example.crowdin.com/u/clients/1/rates"
							}
						}
					],
					"pagination": {
						"offset": 10,
						"limit": 25
					}
				}`)
			})

			clients, resp, err := client.Clients.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.Client{
				{
					ID:          52760,
					Name:        "John Smith",
					Description: "John Smith Organization",
					Status:      "pending",
					WebURL:      "https://example.crowdin.com/u/clients/1/rates",
				},
			}
			assert.Equal(t, expected, clients)

			assert.Equal(t, 10, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestClientsService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/clients", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Clients.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}
