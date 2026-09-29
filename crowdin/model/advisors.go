package model

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// AdvisorAIMode defines the AI validation mode of an advisor inspector.
type AdvisorAIMode string

const (
	// AdvisorAIModeAuto lets the backend decide which items need AI validation.
	AdvisorAIModeAuto AdvisorAIMode = "auto"
	// AdvisorAIModeAll runs AI validation for all items.
	AdvisorAIModeAll AdvisorAIMode = "all"
)

// AdvisorInsightOutcome defines the outcome of an advisor insight check.
type AdvisorInsightOutcome string

const (
	// AdvisorInsightOutcomeFlagged means an issue was detected.
	AdvisorInsightOutcomeFlagged AdvisorInsightOutcome = "flagged"
	// AdvisorInsightOutcomeClear means the project is on target.
	AdvisorInsightOutcomeClear AdvisorInsightOutcome = "clear"
	// AdvisorInsightOutcomeNotApplicable means the inspector does not apply to the project.
	AdvisorInsightOutcomeNotApplicable AdvisorInsightOutcome = "not_applicable"
)

const (
	advisorInsightMaxMetrics         = 100
	advisorInsightMaxRecommendations = 20
)

type (
	// AdvisorCheck represents an advisor check job.
	AdvisorCheck struct {
		Identifier string `json:"identifier"`
		// Status of the underlying job (e.g. created, in_progress, done, failed).
		Status     string                  `json:"status"`
		Progress   int                     `json:"progress"`
		Attributes *AdvisorCheckAttributes `json:"attributes"`
		CreatedAt  string                  `json:"createdAt"`
		UpdatedAt  string                  `json:"updatedAt"`
		StartedAt  *string                 `json:"startedAt"`
		FinishedAt *string                 `json:"finishedAt"`
	}

	// AdvisorCheckAttributes represents the scope of an advisor check.
	// Absence of both fields means an all-inspector check.
	AdvisorCheckAttributes struct {
		Category   *string                  `json:"category,omitempty"`
		Inspectors []*AdvisorCheckInspector `json:"inspectors,omitempty"`
	}

	// AdvisorCheckInspector represents an inspector to re-check.
	AdvisorCheckInspector struct {
		// Inspector key, e.g. string_context_relevance.
		Key string `json:"key"`
		// Per-inspector options. Currently only `string_context_relevance`
		// supports AI options.
		Options *AdvisorInspectorOptions `json:"options,omitempty"`
	}

	// AdvisorInspectorOptions represents per-inspector AI options.
	AdvisorInspectorOptions struct {
		// AI validation mode. Enum: auto, all.
		Mode AdvisorAIMode `json:"mode,omitempty"`
		// AI prompt identifier.
		PromptID int `json:"promptId,omitempty"`
	}
)

// AdvisorCheckResponse defines the structure of a response when
// creating or getting an advisor check.
type AdvisorCheckResponse struct {
	Data *AdvisorCheck `json:"data"`
}

// AdvisorCheckCreateRequest defines the structure of a request to
// create an advisor check.
// At most one of `Category` or `Inspectors` may be set. If neither is set,
// every inspector is re-checked.
type AdvisorCheckCreateRequest struct {
	// Category to scope the check to. Mutually exclusive with `Inspectors`.
	Category string `json:"category,omitempty"`
	// List of inspectors to re-check. Mutually exclusive with `Category`.
	Inspectors []*AdvisorCheckInspector `json:"inspectors,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *AdvisorCheckCreateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Category != "" && len(r.Inspectors) > 0 {
		return errors.New("category and inspectors are mutually exclusive")
	}

	for _, inspector := range r.Inspectors {
		if inspector == nil || inspector.Key == "" {
			return errors.New("inspector key is required")
		}
		if inspector.Options == nil || inspector.Options.Mode == "" {
			continue
		}

		switch inspector.Options.Mode {
		case AdvisorAIModeAuto, AdvisorAIModeAll: // valid
		default:
			return fmt.Errorf("invalid mode: %q, must be one of %s, %s",
				inspector.Options.Mode, AdvisorAIModeAuto, AdvisorAIModeAll)
		}
	}

	return nil
}

type (
	// AdvisorInsight represents an advisor insight.
	AdvisorInsight struct {
		ID           int    `json:"id"`
		InspectorKey string `json:"inspectorKey"`
		// Category of the inspector. Nil if the inspector is no longer registered.
		Category    *string `json:"category"`
		IsDismissed bool    `json:"isDismissed"`
		// Enum: pending, checking, outdated, done.
		Status string `json:"status"`
		// Outcome of the latest check. Nil if not yet checked.
		Outcome *AdvisorInsightOutcome `json:"outcome"`
		// Enum: high, medium, low. Nil if the inspector is no longer registered.
		Severity *string `json:"severity"`
		// Enum: hourly, daily. Nil if the inspector is no longer registered.
		RefreshPolicy   *string                         `json:"refreshPolicy"`
		Metrics         []*AdvisorInsightMetric         `json:"metrics"`
		Recommendations []*AdvisorInsightRecommendation `json:"recommendations"`
		// Date when the last check was performed.
		CheckedAt *string `json:"checkedAt"`
		// The most recent AI validation run for the insight.
		LastAIRun *AdvisorAIRun `json:"lastAiRun"`
		// Data attached to the insight by the originating `advisor-inspector` app.
		Payload map[string]any `json:"payload"`
	}

	// AdvisorInsightMetric represents a structured metric of an advisor insight.
	AdvisorInsightMetric struct {
		Key   string  `json:"key"`
		Value float64 `json:"value"`
		// Enum: percent, count.
		Unit      string   `json:"unit,omitempty"`
		Threshold *float64 `json:"threshold,omitempty"`
		// Enum: default, success, danger.
		Tone string `json:"tone,omitempty"`
		// What produced the metric. Enum: deterministic, ai, app.
		Source string `json:"source,omitempty"`
		// Date when the metric was measured.
		CheckedAt *string `json:"checkedAt,omitempty"`
	}

	// AdvisorInsightRecommendation represents an actionable recommendation
	// of an advisor insight.
	AdvisorInsightRecommendation struct {
		ID      string `json:"id"`
		Primary bool   `json:"primary"`
		// Parameters provided alongside the recommendation (e.g. `projectId`).
		Params map[string]any `json:"params,omitempty"`
	}

	// AdvisorAIRun describes the most recent AI validation run for an insight.
	AdvisorAIRun struct {
		Mode     AdvisorAIMode `json:"mode"`
		PromptID int           `json:"promptId"`
		At       string        `json:"at"`
	}
)

// AdvisorInsightResponse defines the structure of a response when
// getting an advisor insight.
type AdvisorInsightResponse struct {
	Data *AdvisorInsight `json:"data"`
}

// AdvisorInsightsListResponse defines the structure of a response when
// getting a list of advisor insights.
type AdvisorInsightsListResponse struct {
	Data []*AdvisorInsightResponse `json:"data"`
}

// AdvisorInsightsListOptions specifies the optional parameters to the
// AdvisorsService.ListInsights method.
type AdvisorInsightsListOptions struct {
	// Filter insights by dismissal flag. Nil returns both dismissed
	// and active insights.
	IsDismissed *bool `json:"isDismissed,omitempty"`
	// Filter insights by status. Enum: pending, checking, outdated, done.
	Status []string `json:"status,omitempty"`
	// Filter insights by outcome. If empty, all checked insights are
	// returned (never-checked ones are excluded).
	Outcome []AdvisorInsightOutcome `json:"outcome,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of AdvisorInsightsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *AdvisorInsightsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.IsDismissed != nil {
		v.Add("isDismissed", fmt.Sprintf("%t", *o.IsDismissed))
	}
	if len(o.Status) > 0 {
		v.Add("status", strings.Join(o.Status, ","))
	}
	if len(o.Outcome) > 0 {
		v.Add("outcome", JoinSlice(o.Outcome))
	}

	return v, len(v) > 0
}

// ApplicationAdvisorInsightRequest defines the structure of a request to
// create or update an application advisor insight.
type ApplicationAdvisorInsightRequest struct {
	// Result of the app's check. Enum: flagged, clear, not_applicable.
	Outcome AdvisorInsightOutcome `json:"outcome"`
	// Date when the app performed the check (ISO 8601).
	// Defaults to the time the request is received.
	CheckedAt string `json:"checkedAt,omitempty"`
	// Structured metrics produced by the app (max 100).
	Metrics []*AdvisorInsightMetric `json:"metrics,omitempty"`
	// Actionable recommendations to surface for the insight (max 20).
	Recommendations []*AdvisorInsightRecommendation `json:"recommendations,omitempty"`
	// Free-form data to keep attached to the insight.
	Payload map[string]any `json:"payload,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ApplicationAdvisorInsightRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}

	switch r.Outcome {
	case AdvisorInsightOutcomeFlagged, AdvisorInsightOutcomeClear, AdvisorInsightOutcomeNotApplicable: // valid
	default:
		return fmt.Errorf("invalid outcome: %q, must be one of %s, %s, %s", r.Outcome,
			AdvisorInsightOutcomeFlagged, AdvisorInsightOutcomeClear, AdvisorInsightOutcomeNotApplicable)
	}

	if len(r.Metrics) > advisorInsightMaxMetrics {
		return fmt.Errorf("metrics cannot contain more than %d items", advisorInsightMaxMetrics)
	}
	if len(r.Recommendations) > advisorInsightMaxRecommendations {
		return fmt.Errorf("recommendations cannot contain more than %d items", advisorInsightMaxRecommendations)
	}

	return nil
}
