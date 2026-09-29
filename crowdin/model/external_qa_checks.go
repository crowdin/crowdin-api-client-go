package model

// ExternalQACheck represents an external QA check of the organization.
type ExternalQACheck struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description *string        `json:"description,omitempty"`
	Config      map[string]any `json:"config,omitempty"`
	CreatedAt   string         `json:"createdAt"`
	UpdatedAt   *string        `json:"updatedAt,omitempty"`
}

// ExternalQACheckResponse defines the structure of a response when
// getting an external QA check.
type ExternalQACheckResponse struct {
	Data *ExternalQACheck `json:"data"`
}

// ExternalQAChecksListResponse defines the structure of a response when
// getting a list of external QA checks.
type ExternalQAChecksListResponse struct {
	Data []*ExternalQACheckResponse `json:"data"`
}
