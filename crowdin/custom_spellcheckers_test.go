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

func TestCustomSpellcheckersService_List(t *testing.T) {
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
			opts:          &model.ListOptions{Offset: 5, Limit: 25},
			expectedQuery: "?limit=25&offset=5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/custom-spellcheckers"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 2,
								"name": "Custom Spellchecker",
								"config": {
									"identifier": "app-identifier",
									"key": "app-module-key",
									"realTimeCheckEnabled": true,
									"enabledLanguageIds": ["en", "es"]
								},
								"createdAt": "2019-09-23T07:19:47+00:00",
								"updatedAt": null
							}
						}
					],
					"pagination": {
						"offset": 5,
						"limit": 25
					}
				}`)
			})

			spellcheckers, resp, err := client.CustomSpellcheckers.List(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.CustomSpellchecker{
				{
					ID:   2,
					Name: "Custom Spellchecker",
					Config: &model.CustomSpellcheckerConfig{
						Identifier:           "app-identifier",
						Key:                  "app-module-key",
						RealTimeCheckEnabled: true,
						EnabledLanguageIDs:   []string{"en", "es"},
					},
					CreatedAt: "2019-09-23T07:19:47+00:00",
				},
			}
			assert.Equal(t, expected, spellcheckers)

			assert.Equal(t, 5, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestCustomSpellcheckersService_List_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/custom-spellcheckers", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.CustomSpellcheckers.List(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestCustomSpellcheckersService_Get(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/custom-spellcheckers/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"name": "Custom Spellchecker",
				"config": {
					"identifier": "app-identifier",
					"key": "app-module-key",
					"realTimeCheckEnabled": false,
					"enabledLanguageIds": ["de"]
				},
				"createdAt": "2019-09-23T07:19:47+00:00",
				"updatedAt": "2019-09-24T07:19:47+00:00"
			}
		}`)
	})

	spellchecker, resp, err := client.CustomSpellcheckers.Get(context.Background(), 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.CustomSpellchecker{
		ID:   2,
		Name: "Custom Spellchecker",
		Config: &model.CustomSpellcheckerConfig{
			Identifier:         "app-identifier",
			Key:                "app-module-key",
			EnabledLanguageIDs: []string{"de"},
		},
		CreatedAt: "2019-09-23T07:19:47+00:00",
		UpdatedAt: ToPtr("2019-09-24T07:19:47+00:00"),
	}
	assert.Equal(t, expected, spellchecker)
}

func TestCustomSpellcheckersService_Get_notFound(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/custom-spellcheckers/1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error": {"code": 404, "message": "Custom Spellchecker Not Found"}}`)
	})

	res, resp, err := client.CustomSpellcheckers.Get(context.Background(), 1)
	require.Error(t, err)

	var errResponse *model.ErrorResponse
	assert.ErrorAs(t, err, &errResponse)
	assert.Equal(t, "404 Custom Spellchecker Not Found", errResponse.Error())

	assert.Nil(t, res)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
