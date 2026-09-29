package model

import (
	"errors"
	"fmt"
	"net/url"
)

// Group represents a Crowdin group.
type Group struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ParentID       int    `json:"parentId"`
	OrganizationID int    `json:"organizationId"`
	UserID         int    `json:"userId"`
	SubgroupsCount int    `json:"subgroupsCount"`
	ProjectsCount  int    `json:"projectsCount"`
	WebURL         string `json:"webUrl"`
	// Report settings template identifier. Available only for group
	// managers and organization admins.
	SavingsReportSettingsTemplateID *int   `json:"savingsReportSettingsTemplateId,omitempty"`
	CreatedAt                       string `json:"createdAt"`
	UpdatedAt                       string `json:"updatedAt"`
}

// GroupsListOptions specifies the optional parameters to the GroupsService.List method.
type GroupsListOptions struct {
	ListOptions

	// Parent group identifier.
	ParentID int `json:"parentId,omitempty"`
	// Sort a list of groups by a specified field.
	// Example: orderBy=createdAt desc,name.
	OrderBy string `json:"orderBy,omitempty"`
	// Filter groups by `name`.
	Filter string `json:"filter,omitempty"`
}

// Values returns the url.Values representation of the GroupsListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *GroupsListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.ParentID > 0 {
		v.Add("parentId", fmt.Sprintf("%d", o.ParentID))
	}
	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}
	if o.Filter != "" {
		v.Add("filter", o.Filter)
	}

	return v, len(v) > 0
}

// GroupsGetResponse defines the structure of a response when retrieving a group.
type GroupsGetResponse struct {
	Data *Group `json:"data"`
}

// GroupsListResponse defines the structure of a response when getting a list of groups.
type GroupsListResponse struct {
	Data       []*GroupsGetResponse `json:"data"`
	Pagination *Pagination          `json:"pagination"`
}

// GroupsAddRequest defines the structure of a request to add a group.
type GroupsAddRequest struct {
	// Group Name (required).
	Name string `json:"name"`
	// Parent Group Identifier.
	ParentID int `json:"parentId,omitempty"`
	// Group description.
	Description string `json:"description,omitempty"`
	// Report settings template identifier.
	SavingsReportSettingsTemplateID int `json:"savingsReportSettingsTemplateId,omitempty"`
}

// Validate checks if the add request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *GroupsAddRequest) Validate() error {
	if r == nil {
		return ErrNilRequest
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
