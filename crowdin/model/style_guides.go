package model

import (
	"errors"
	"fmt"
	"net/url"
)

// StyleGuide represents a Crowdin style guide.
type StyleGuide struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	AIInstructions *string  `json:"aiInstructions"`
	UserID         int      `json:"userId"`
	GroupID        int      `json:"groupId,omitempty"` // Enterprise only.
	LanguageIDs    []string `json:"languageIds"`
	ProjectIDs     []int    `json:"projectIds"`
	IsShared       bool     `json:"isShared"`
	WebURL         string   `json:"webUrl"`
	DownloadLink   string   `json:"downloadLink"`
	CreatedAt      string   `json:"createdAt"`
	UpdatedAt      string   `json:"updatedAt"`
}

// StyleGuideResponse defines the structure of a response when
// getting a style guide.
type StyleGuideResponse struct {
	Data *StyleGuide `json:"data"`
}

// StyleGuidesListResponse defines the structure of a response when
// getting a list of style guides.
type StyleGuidesListResponse struct {
	Data []*StyleGuideResponse `json:"data"`
}

// StyleGuidesListOptions specifies the optional parameters to the
// StyleGuidesService.List method.
type StyleGuidesListOptions struct {
	// Sort style guides by the specified field.
	// Example: orderBy=createdAt desc,name.
	OrderBy string `json:"orderBy,omitempty"`
	// List style guides of a specific user.
	UserID int `json:"userId,omitempty"`

	ListOptions
}

// Values returns the url.Values representation of StyleGuidesListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *StyleGuidesListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.UserID > 0 {
		v.Add("userId", fmt.Sprintf("%d", o.UserID))
	}

	return v, len(v) > 0
}

// StyleGuideCreateRequest defines the structure of a request to create a style guide.
type StyleGuideCreateRequest struct {
	// Style guide name.
	Name string `json:"name"`
	// Storage identifier of the style guide file.
	// The field is always sent: a nil value is transmitted as `null`.
	StorageID *int `json:"storageId"`
	// Group Identifier - defines group to which the style guide is added.
	// If omitted, root group is used.
	// Note: Enterprise only.
	GroupID *int `json:"groupId,omitempty"`
	// Rules to be used by AI models.
	// Note: maxLength (10240) counts bytes, not characters.
	AIInstructions string `json:"aiInstructions,omitempty"`
	// Style guide language identifiers.
	LanguageIDs []string `json:"languageIds,omitempty"`
	// Project identifiers to assign the style guide to.
	ProjectIDs []int `json:"projectIds,omitempty"`
	// Whether the style guide should be shared across all projects
	// within the account (Crowdin) or the group (Enterprise).
	IsShared *bool `json:"isShared,omitempty"`
}

// Validate checks if the request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *StyleGuideCreateRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
