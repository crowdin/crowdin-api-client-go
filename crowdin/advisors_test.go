package crowdin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/crowdin/crowdin-api-client-go/crowdin/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const advisorCheckJSON = `{
	"identifier": "550e8400-e29b-41d4-a716-446655440000",
	"status": "created",
	"progress": 0,
	"attributes": {
		"category": null,
		"inspectors": [{"key": "string_context_relevance"}]
	},
	"createdAt": "2026-04-16T12:00:00+00:00",
	"updatedAt": "2026-04-16T12:00:00+00:00",
	"startedAt": null,
	"finishedAt": null
}`

const advisorInsightJSON = `{
	"id": 10,
	"inspectorKey": "string_context_relevance",
	"category": "context",
	"isDismissed": false,
	"status": "done",
	"outcome": "flagged",
	"severity": "high",
	"refreshPolicy": "hourly",
	"metrics": [
		{
			"key": "quality_score",
			"value": 12.5,
			"unit": "percent",
			"threshold": 50,
			"tone": "danger",
			"source": "deterministic",
			"checkedAt": "2026-04-06T10:30:00+00:00"
		}
	],
	"recommendations": [
		{"id": "run_ai_context_validation", "primary": true, "params": {"projectId": 1}}
	],
	"checkedAt": "2026-04-06T10:30:00+00:00",
	"lastAiRun": {"mode": "auto", "promptId": 42, "at": "2026-04-06T10:30:00+00:00"},
	"payload": {"foo": "bar"}
}`

func expectedAdvisorCheck() *model.AdvisorCheck {
	return &model.AdvisorCheck{
		Identifier: "550e8400-e29b-41d4-a716-446655440000",
		Status:     "created",
		Progress:   0,
		Attributes: &model.AdvisorCheckAttributes{
			Inspectors: []*model.AdvisorCheckInspector{{Key: "string_context_relevance"}},
		},
		CreatedAt: "2026-04-16T12:00:00+00:00",
		UpdatedAt: "2026-04-16T12:00:00+00:00",
	}
}

func expectedAdvisorInsight() *model.AdvisorInsight {
	checkedAt := "2026-04-06T10:30:00+00:00"

	return &model.AdvisorInsight{
		ID:            10,
		InspectorKey:  "string_context_relevance",
		Category:      ToPtr("context"),
		IsDismissed:   false,
		Status:        "done",
		Outcome:       ToPtr(model.AdvisorInsightOutcomeFlagged),
		Severity:      ToPtr("high"),
		RefreshPolicy: ToPtr("hourly"),
		Metrics: []*model.AdvisorInsightMetric{
			{
				Key:       "quality_score",
				Value:     12.5,
				Unit:      "percent",
				Threshold: ToPtr(50.0),
				Tone:      "danger",
				Source:    "deterministic",
				CheckedAt: ToPtr(checkedAt),
			},
		},
		Recommendations: []*model.AdvisorInsightRecommendation{
			{ID: "run_ai_context_validation", Primary: true, Params: map[string]any{"projectId": float64(1)}},
		},
		CheckedAt: ToPtr(checkedAt),
		LastAIRun: &model.AdvisorAIRun{Mode: model.AdvisorAIModeAuto, PromptID: 42, At: checkedAt},
		Payload:   map[string]any{"foo": "bar"},
	}
}

func TestAdvisorsService_CreateCheck(t *testing.T) {
	tests := []struct {
		name string
		req  *model.AdvisorCheckCreateRequest
		body string
	}{
		{
			name: "nil request",
			req:  nil,
			body: `{}`,
		},
		{
			name: "by category",
			req:  &model.AdvisorCheckCreateRequest{Category: "context"},
			body: `{"category": "context"}`,
		},
		{
			name: "by inspectors",
			req: &model.AdvisorCheckCreateRequest{
				Inspectors: []*model.AdvisorCheckInspector{
					{
						Key:     "string_context_relevance",
						Options: &model.AdvisorInspectorOptions{Mode: model.AdvisorAIModeAuto, PromptID: 42},
					},
					{Key: "glossary_coverage"},
				},
			},
			body: `{"inspectors": [
				{"key": "string_context_relevance", "options": {"mode": "auto", "promptId": 42}},
				{"key": "glossary_coverage"}
			]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/advisors/checks"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)
				testURL(t, r, path)
				testJSONBody(t, r, tt.body)

				w.WriteHeader(http.StatusAccepted)
				fmt.Fprintf(w, `{"data": %s}`, advisorCheckJSON)
			})

			check, resp, err := client.Advisors.CreateCheck(context.Background(), 1, tt.req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusAccepted, resp.StatusCode)
			assert.Equal(t, expectedAdvisorCheck(), check)
		})
	}
}

func TestAdvisorsService_CreateCheck_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	req := &model.AdvisorCheckCreateRequest{
		Category:   "context",
		Inspectors: []*model.AdvisorCheckInspector{{Key: "string_context_relevance"}},
	}
	_, _, err := client.Advisors.CreateCheck(context.Background(), 1, req)
	require.EqualError(t, err, "category and inspectors are mutually exclusive")
}

func TestAdvisorsService_GetCheck(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/advisors/checks/550e8400-e29b-41d4-a716-446655440000"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, path)

		fmt.Fprintf(w, `{"data": %s}`, advisorCheckJSON)
	})

	check, resp, err := client.Advisors.GetCheck(context.Background(), 1, "550e8400-e29b-41d4-a716-446655440000")
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedAdvisorCheck(), check)
}

func TestAdvisorsService_GetCheck_notFound(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/advisors/checks/unknown"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		http.Error(w, `{"error": {"code": 404, "message": "Check Not Found"}}`, http.StatusNotFound)
	})

	check, resp, err := client.Advisors.GetCheck(context.Background(), 1, "unknown")
	require.Error(t, err)
	assert.Nil(t, check)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var e *model.ErrorResponse
	assert.True(t, errors.As(err, &e))
}

func TestAdvisorsService_ListInsights(t *testing.T) {
	tests := []struct {
		name     string
		opts     *model.AdvisorInsightsListOptions
		expected string
	}{
		{
			name:     "nil options",
			opts:     nil,
			expected: "",
		},
		{
			name:     "empty options",
			opts:     &model.AdvisorInsightsListOptions{},
			expected: "",
		},
		{
			name: "with options",
			opts: &model.AdvisorInsightsListOptions{
				IsDismissed: ToPtr(false),
				Status:      []string{"pending", "done"},
				Outcome:     []model.AdvisorInsightOutcome{model.AdvisorInsightOutcomeFlagged, model.AdvisorInsightOutcomeClear},
				ListOptions: model.ListOptions{Offset: 1, Limit: 25},
			},
			expected: "?isDismissed=false&limit=25&offset=1&outcome=flagged%2Cclear&status=pending%2Cdone",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, teardown := setupClient()
			defer teardown()

			const path = "/api/v2/projects/1/advisors/insights"
			mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				testURL(t, r, path+tt.expected)

				fmt.Fprintf(w, `{
					"data": [{"data": %s}],
					"pagination": {"offset": 1, "limit": 25, "total": 1}
				}`, advisorInsightJSON)
			})

			insights, resp, err := client.Advisors.ListInsights(context.Background(), 1, tt.opts)
			require.NoError(t, err)

			assert.Equal(t, []*model.AdvisorInsight{expectedAdvisorInsight()}, insights)
			assert.Equal(t, 1, resp.Pagination.Offset)
			assert.Equal(t, 25, resp.Pagination.Limit)
		})
	}
}

func TestAdvisorsService_ListInsights_invalidJSON(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	mux.HandleFunc("/api/v2/projects/1/advisors/insights", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, `invalid json`)
	})

	res, _, err := client.Advisors.ListInsights(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Nil(t, res)
}

func TestAdvisorsService_EditInsight(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/advisors/insights/10"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testURL(t, r, path)
		testJSONBodyAny(t, r, `[{"op": "replace", "path": "/isDismissed", "value": false}]`)

		fmt.Fprintf(w, `{"data": %s}`, advisorInsightJSON)
	})

	req := []*model.UpdateRequest{{Op: model.OpReplace, Path: "/isDismissed", Value: false}}
	insight, resp, err := client.Advisors.EditInsight(context.Background(), 1, 10, req)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expectedAdvisorInsight(), insight)
}

func TestAdvisorsService_CreateOrUpdateApplicationInsight(t *testing.T) {
	client, mux, teardown := setupClient()
	defer teardown()

	const path = "/api/v2/projects/1/applications/my-app/modules/my-inspector/advisors/insights"
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testURL(t, r, path)
		testJSONBody(t, r, `{
			"outcome": "flagged",
			"checkedAt": "2026-05-21T14:30:00+00:00",
			"metrics": [{"key": "quality_score", "value": 12, "unit": "percent", "tone": "danger", "source": "app"}],
			"recommendations": [{"id": "run_ai_context_validation", "primary": false}],
			"payload": {"foo": "bar"}
		}`)

		w.WriteHeader(http.StatusNoContent)
	})

	req := &model.ApplicationAdvisorInsightRequest{
		Outcome:   model.AdvisorInsightOutcomeFlagged,
		CheckedAt: "2026-05-21T14:30:00+00:00",
		Metrics: []*model.AdvisorInsightMetric{
			{Key: "quality_score", Value: 12, Unit: "percent", Tone: "danger", Source: "app"},
		},
		Recommendations: []*model.AdvisorInsightRecommendation{{ID: "run_ai_context_validation"}},
		Payload:         map[string]any{"foo": "bar"},
	}
	resp, err := client.Advisors.CreateOrUpdateApplicationInsight(context.Background(), 1, "my-app", "my-inspector", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestAdvisorsService_CreateOrUpdateApplicationInsight_invalidRequest(t *testing.T) {
	client, _, teardown := setupClient()
	defer teardown()

	_, err := client.Advisors.CreateOrUpdateApplicationInsight(context.Background(), 1, "my-app", "my-inspector", nil)
	require.ErrorIs(t, err, model.ErrNilRequest)

	req := &model.ApplicationAdvisorInsightRequest{Outcome: "unknown"}
	_, err = client.Advisors.CreateOrUpdateApplicationInsight(context.Background(), 1, "my-app", "my-inspector", req)
	require.EqualError(t, err, `invalid outcome: "unknown", must be one of flagged, clear, not_applicable`)
}
