package model

import (
	"fmt"
	"net/url"
)

type Manager struct {
	ID    int    `json:"id"`
	User  User   `json:"user"`
	Teams []Team `json:"teams"`
}

type ManagerGetResponse struct {
	Data *Manager `json:"data"`
}

type ManagerListResponse struct {
	Data       []*ManagerGetResponse `json:"data"`
	Pagination *Pagination           `json:"pagination"`
}

type ManagerEditResponse struct {
	Data []*ManagerGetResponse `json:"data"`
}

type ManagerListOptions struct {
	ListOptions

	TeamIds int `json:"teamIds,omitempty"`

	OrderBy string `json:"orderBy,omitempty"`
}

// Values returns the url.Values representation of the ManagerListOptions.
// It implements the crowdin.ListOptionsProvider interface.
func (o *ManagerListOptions) Values() (url.Values, bool) {
	if o == nil {
		return nil, false
	}

	v, _ := o.ListOptions.Values()
	if o.TeamIds > 0 {
		v.Add("parentId", fmt.Sprintf("%d", o.TeamIds))
	}

	if o.OrderBy != "" {
		v.Add("orderBy", o.OrderBy)
	}

	return v, len(v) > 0
}
