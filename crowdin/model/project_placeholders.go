package model

import (
	"errors"
	"fmt"
)

// ProjectPlaceholderType defines the priority a project placeholder is matched with.
type ProjectPlaceholderType string

const (
	ProjectPlaceholderTypeHigh ProjectPlaceholderType = "high"
	ProjectPlaceholderTypeLow  ProjectPlaceholderType = "low"
)

// ProjectPlaceholder represents a custom placeholder of the organization,
// with the settings it has in the project.
type ProjectPlaceholder struct {
	// Identifier of the placeholder in the project.
	ID int `json:"id"`
	// Custom placeholder of the organization this row assigns.
	CustomPlaceholderID int `json:"customPlaceholderId"`
	// Priority the placeholder is matched with. Enum: high, low.
	Type *string `json:"type,omitempty"`
	// Position of the placeholder among the ones assigned to the project.
	Index *int `json:"index,omitempty"`
	// If true, a mismatch of the placeholder blocks the translation from being saved.
	IsBlocking *bool `json:"isBlocking,omitempty"`
	// File formats the placeholder is applied to; an empty array means every format.
	Formats []string `json:"formats"`
}

// ProjectPlaceholderResponse defines the structure of a response when
// getting a project placeholder.
type ProjectPlaceholderResponse struct {
	Data *ProjectPlaceholder `json:"data"`
}

// ProjectPlaceholdersListResponse defines the structure of a response when
// getting a list of project placeholders.
type ProjectPlaceholdersListResponse struct {
	Data []*ProjectPlaceholderResponse `json:"data"`
}

// ProjectPlaceholderAddRequest defines the structure of a request when
// assigning a custom placeholder to the project.
type ProjectPlaceholderAddRequest struct {
	// ID of the custom placeholder the organization defined.
	CustomPlaceholderID int `json:"customPlaceholderId"`
	// Priority the placeholder is matched with. Enum: high, low. Default: low.
	Type ProjectPlaceholderType `json:"type,omitempty"`
	// Position of the placeholder among the ones assigned to the project. Default: 0.
	Index int `json:"index,omitempty"`
	// If true, a mismatch of the placeholder blocks the translation from being saved.
	// Default: false.
	IsBlocking *bool `json:"isBlocking,omitempty"`
	// File formats the placeholder is applied to; an empty array means every format.
	Formats []string `json:"formats,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *ProjectPlaceholderAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.CustomPlaceholderID == 0 {
		return errors.New("customPlaceholderId is required")
	}
	if r.Type != "" {
		switch r.Type {
		case ProjectPlaceholderTypeHigh, ProjectPlaceholderTypeLow: // valid
		default:
			return fmt.Errorf("invalid type: %q, must be one of %s, %s",
				r.Type, ProjectPlaceholderTypeHigh, ProjectPlaceholderTypeLow)
		}
	}

	return nil
}

// ProjectSystemPlaceholder represents a placeholder Crowdin ships,
// turned on or off per project.
type ProjectSystemPlaceholder struct {
	// Key of the system placeholder (e.g. "printfSpecifier"). Name it in the
	// `path` of a batch operation to turn the placeholder on or off.
	ID string `json:"id"`
	// If false, the project stops recognizing the placeholder in the QA checks
	// and in the Editor.
	IsEnabled bool `json:"isEnabled"`
	// Name of the placeholder, in the language of the request.
	Label string `json:"label"`
	// One line saying where the construct comes from, in the language of the request.
	Description string `json:"description"`
	// Constructs the placeholder matches, written out.
	Examples []string `json:"examples"`
}

// ProjectSystemPlaceholderResponse defines the structure of a response when
// getting a project system placeholder.
type ProjectSystemPlaceholderResponse struct {
	Data *ProjectSystemPlaceholder `json:"data"`
}

// ProjectSystemPlaceholdersListResponse defines the structure of a response when
// getting a list of project system placeholders.
type ProjectSystemPlaceholdersListResponse struct {
	Data []*ProjectSystemPlaceholderResponse `json:"data"`
}
