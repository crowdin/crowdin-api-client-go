package model

import "errors"

// Group represents a Crowdin group.
type Group struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ParentID       int64  `json:"parentId"`
	OrganizationID int64  `json:"organizationId"`
	UserID         int64  `json:"userId"`
	SubgroupsCount int64  `json:"subgroupsCount"`
	ProjectsCount  int64  `json:"projectsCount"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// GroupsListOptions specifies the optional parameters to the GroupsService.List method.
type GroupsListOptions struct {
	ListOptions

	ParentID int64 `json:"parentId,omitempty"`
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
	ParentID int64 `json:"parentId,omitempty"`
	// Group description.
	Description string `json:"description,omitempty"`
}

// Validate checks if the add request is valid.
// It implements the crowdin.RequestValidator interface.
func (r *GroupsAddRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
