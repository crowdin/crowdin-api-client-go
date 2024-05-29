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

func TestWorkflowsService_GetTemplate(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/workflow-templates/1"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 2,
				"title": "In-house + Machine Translation",
				"description": "Combine the efforts of human translators and Machine Translation technology.",
				"groupId": 2,
				"isDefault": true,
				"steps": [
					{
						"id": 3,
						"languages": [2],
						"assignees": [5],
						"vendorId": 52760,
						"config": {
							"minRelevant": 0,
							"autoSubstitution": false
						},
						"mtId": 12
					}
				],
				"webUrl": "https://example.crowdin.com/u/workflows/1/read"
			}
		}`)
	})

	template, resp, err := client.Workflows.GetTemplate(context.Background(), 1)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.WorkflowTemplate{
		ID:          2,
		Title:       "In-house + Machine Translation",
		Description: "Combine the efforts of human translators and Machine Translation technology.",
		GroupID:     2,
		IsDefault:   true,
		Steps: []*model.WorkflowTemplateStep{
			{
				ID:        3,
				Languages: []int{2},
				Assignees: []int{5},
				VendorID:  52760,
				Config: model.WorkflowTemplateStepConfig{
					MinRelevant:      ToPtr(0),
					AutoSubstitution: ToPtr(false),
				},
				MTID: 12,
			},
		},
		WebURL: "https://example.crowdin.com/u/workflows/1/read",
	}
	assert.Equal(t, expected, template)
}

func TestWorkflowsService_ListTemplates(t *testing.T) {
	tests := []struct {
		name          string
		opts          *model.WorkflowTemplatesListOptions
		expectedQuery string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &model.WorkflowTemplatesListOptions{},
		},
		{
			name:          "groupId=0",
			opts:          &model.WorkflowTemplatesListOptions{GroupID: ToPtr(0)},
			expectedQuery: "?groupId=0",
		},
		{
			name: "all options",
			opts: &model.WorkflowTemplatesListOptions{
				GroupID:     ToPtr(1),
				ListOptions: model.ListOptions{Offset: 10, Limit: 25},
			},
			expectedQuery: "?groupId=1&limit=25&offset=10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/workflow-templates"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expectedQuery)

				fmt.Fprint(w, `{
					"data": [
						{
							"data": {
								"id": 2,
								"title": "In-house + Machine Translation",
								"description": "Combine the efforts of human translators and Machine Translation technology.",
								"groupId": 2,
								"isDefault": true,
								"steps": [
									{
										"id": 3,
										"languages": [2],
										"assignees": [5],
										"vendorId": 52760,
										"config": {
											"minRelevant": 60,
											"autoSubstitution": true
										},
										"mtId": 12
									}
								],
								"webUrl": "https://example.crowdin.com/u/workflows/1/read"
							}
						}
					],
					"pagination": {
						"offset": 10,
						"limit": 25
					}
				}`)
			})

			templates, resp, err := client.Workflows.ListTemplates(context.Background(), tt.opts)
			require.NoError(t, err)

			expected := []*model.WorkflowTemplate{
				{
					ID:          2,
					Title:       "In-house + Machine Translation",
					Description: "Combine the efforts of human translators and Machine Translation technology.",
					GroupID:     2,
					IsDefault:   true,
					Steps: []*model.WorkflowTemplateStep{
						{
							ID:        3,
							Languages: []int{2},
							Assignees: []int{5},
							VendorID:  52760,
							Config: model.WorkflowTemplateStepConfig{
								MinRelevant:      ToPtr(60),
								AutoSubstitution: ToPtr(true),
							},
							MTID: 12,
						},
					},
					WebURL: "https://example.crowdin.com/u/workflows/1/read",
				},
			}
			assert.Equal(t, expected, templates)

			assert.Equal(t, 10, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestWorkflowsService_ListTemplates_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/workflow-templates", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	templates, _, err := client.Workflows.ListTemplates(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, templates)
}

func TestWorkflowsService_GetStep(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/workflow-steps/2"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": {
				"id": 311,
				"title": "Translate",
				"type": "Translate",
				"languages": ["de","it"],
				"config": {
					"assignees": {
						"de": [346],
						"it": [43]
					}
				}
			}
		}`)
	})

	step, resp, err := client.Workflows.GetStep(context.Background(), 1, 2)
	require.NoError(t, err)
	assert.NotNil(t, resp)

	expected := &model.WorkflowStep{
		ID:        311,
		Title:     "Translate",
		Type:      "Translate",
		Languages: []string{"de", "it"},
		Config: map[string]any{
			"assignees": map[string]any{
				"de": []any{float64(346)},
				"it": []any{float64(43)},
			},
		},
	}
	assert.Equal(t, expected, step)
}

func TestWorkflowsService_ListSteps(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/workflow-steps"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprint(w, `{
			"data": [
				{
					"data": {
						"id": 311,
						"title": "Translate",
						"type": "Translate",
						"languages": ["de","it"],
						"config": {
							"minRelevant": "perfect",
							"autoSubstitution": 2
						}
					}
				}
			],
			"pagination": {
				"offset": 10,
				"limit": 25
			}
		}`)
	})

	steps, resp, err := client.Workflows.ListSteps(context.Background(), "1")
	require.NoError(t, err)

	expected := []*model.WorkflowStep{
		{
			ID:        311,
			Title:     "Translate",
			Type:      "Translate",
			Languages: []string{"de", "it"},
			Config: map[string]any{
				"minRelevant":      "perfect",
				"autoSubstitution": float64(2),
			},
		},
	}
	assert.Equal(t, expected, steps)

	assert.Equal(t, 10, resp.Pagination.Offset)
	assert.Equal(t, 25, resp.Pagination.Limit)
}

func TestWorkflowsService_ListSteps_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/workflow-steps", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	steps, _, err := client.Workflows.ListSteps(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, steps)
}
