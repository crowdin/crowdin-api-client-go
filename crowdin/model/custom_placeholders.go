package model

import "errors"

// CustomPlaceholder represents a custom placeholder of the organization.
type CustomPlaceholder struct {
	ID int `json:"id"`
	// Placeholder definition.
	Definition string `json:"definition,omitempty"`
	// Custom placeholder description. `nil` when none is set.
	Description *string `json:"description,omitempty"`
	// Argument delimiter.
	ArgumentDelimiter string `json:"argumentDelimiter,omitempty"`
}

// CustomPlaceholderResponse defines the structure of a response when
// getting a custom placeholder.
type CustomPlaceholderResponse struct {
	Data *CustomPlaceholder `json:"data"`
}

// CustomPlaceholdersListResponse defines the structure of a response when
// getting a list of custom placeholders.
type CustomPlaceholdersListResponse struct {
	Data []*CustomPlaceholderResponse `json:"data"`
}

// CustomPlaceholderAddRequest defines the structure of a request when
// adding a custom placeholder.
type CustomPlaceholderAddRequest struct {
	// Placeholder definition.
	Definition string `json:"definition"`
	// Custom placeholder description.
	Description string `json:"description,omitempty"`
	// Argument delimiter. Default: `"`.
	ArgumentDelimiter string `json:"argumentDelimiter,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *CustomPlaceholderAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Definition == "" {
		return errors.New("definition is required")
	}

	return nil
}
