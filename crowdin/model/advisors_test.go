package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvisorCheckCreateRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *AdvisorCheckCreateRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name:  "empty request",
			req:   &AdvisorCheckCreateRequest{},
			valid: true,
		},
		{
			name:  "category only",
			req:   &AdvisorCheckCreateRequest{Category: "context"},
			valid: true,
		},
		{
			name: "category and inspectors",
			req: &AdvisorCheckCreateRequest{
				Category:   "context",
				Inspectors: []*AdvisorCheckInspector{{Key: "string_context_relevance"}},
			},
			err: "category and inspectors are mutually exclusive",
		},
		{
			name: "nil inspector",
			req:  &AdvisorCheckCreateRequest{Inspectors: []*AdvisorCheckInspector{nil}},
			err:  "inspector key is required",
		},
		{
			name: "missing inspector key",
			req:  &AdvisorCheckCreateRequest{Inspectors: []*AdvisorCheckInspector{{}}},
			err:  "inspector key is required",
		},
		{
			name: "invalid mode",
			req: &AdvisorCheckCreateRequest{
				Inspectors: []*AdvisorCheckInspector{
					{Key: "string_context_relevance", Options: &AdvisorInspectorOptions{Mode: "some"}},
				},
			},
			err: `invalid mode: "some", must be one of auto, all`,
		},
		{
			name: "valid inspectors",
			req: &AdvisorCheckCreateRequest{
				Inspectors: []*AdvisorCheckInspector{
					{Key: "string_context_relevance", Options: &AdvisorInspectorOptions{Mode: AdvisorAIModeAll, PromptID: 42}},
					{Key: "glossary_coverage", Options: &AdvisorInspectorOptions{}},
					{Key: "tm_coverage"},
				},
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.valid {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}

func TestAdvisorInsightsListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *AdvisorInsightsListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &AdvisorInsightsListOptions{},
		},
		{
			name: "dismissed false",
			opts: &AdvisorInsightsListOptions{IsDismissed: toPtr(false)},
			out:  "isDismissed=false",
		},
		{
			name: "with all options",
			opts: &AdvisorInsightsListOptions{
				IsDismissed: toPtr(true),
				Status:      []string{"pending", "outdated"},
				Outcome:     []AdvisorInsightOutcome{AdvisorInsightOutcomeNotApplicable, AdvisorInsightOutcomeClear},
				ListOptions: ListOptions{Limit: 10, Offset: 5},
			},
			out: "isDismissed=true&limit=10&offset=5&outcome=not_applicable%2Cclear&status=pending%2Coutdated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, val.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, val)
			}
		})
	}
}

func TestApplicationAdvisorInsightRequestValidate(t *testing.T) {
	tests := []struct {
		name  string
		req   *ApplicationAdvisorInsightRequest
		err   string
		valid bool
	}{
		{
			name: "nil request",
			req:  nil,
			err:  "request cannot be nil",
		},
		{
			name: "empty request",
			req:  &ApplicationAdvisorInsightRequest{},
			err:  `invalid outcome: "", must be one of flagged, clear, not_applicable`,
		},
		{
			name: "invalid outcome",
			req:  &ApplicationAdvisorInsightRequest{Outcome: "ok"},
			err:  `invalid outcome: "ok", must be one of flagged, clear, not_applicable`,
		},
		{
			name: "too many metrics",
			req: &ApplicationAdvisorInsightRequest{
				Outcome: AdvisorInsightOutcomeClear,
				Metrics: make([]*AdvisorInsightMetric, 101),
			},
			err: "metrics cannot contain more than 100 items",
		},
		{
			name: "too many recommendations",
			req: &ApplicationAdvisorInsightRequest{
				Outcome:         AdvisorInsightOutcomeFlagged,
				Recommendations: make([]*AdvisorInsightRecommendation, 21),
			},
			err: "recommendations cannot contain more than 20 items",
		},
		{
			name:  "outcome only",
			req:   &ApplicationAdvisorInsightRequest{Outcome: AdvisorInsightOutcomeNotApplicable},
			valid: true,
		},
		{
			name: "all fields",
			req: &ApplicationAdvisorInsightRequest{
				Outcome:         AdvisorInsightOutcomeFlagged,
				CheckedAt:       "2026-05-21T14:30:00+00:00",
				Metrics:         []*AdvisorInsightMetric{{Key: "quality_score", Value: 12, Threshold: toPtr(50.0)}},
				Recommendations: []*AdvisorInsightRecommendation{{ID: "run_ai_context_validation", Primary: true}},
				Payload:         map[string]any{"foo": "bar"},
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.req.Validate(); tt.valid {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
